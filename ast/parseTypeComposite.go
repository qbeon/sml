package ast

import "fmt"

// parseTypeComposite parses a struct type declaration statement
func parseTypeComposite(
	ctx *contextFile,
	start Cursor,
	compositeType GraphNodeTypeComposite,
) (AbstractTypeDeclaration, error) {
	src := ctx.Source()

	// Skip the keyword
	var indexAfterKeyword uint
	switch compositeType {
	case GNTCStruct:
		indexAfterKeyword = start.Index + uint(len(keywordStruct))
	case GNTCEntity:
		indexAfterKeyword = start.Index + uint(len(keywordEntity))
	case GNTCUser:
		indexAfterKeyword = start.Index + uint(len(keywordUser))
	default:
		panic("unexpected branch")
	}
	end := Cursor{
		Index:  indexAfterKeyword,
		Column: indexAfterKeyword,
		Line:   start.Line,
	}

	// Read name
	nameFragment := readLatinWord(ctx.Input, end)
	if nameFragment.IsEOF() {
		return nil, &ErrSyntax{
			Err: fmt.Errorf(
				"missing %s name",
				compositeType.String(),
			),
		}
	}

	// Validate type name
	if err := validateTypeName(
		nameFragment,
		GraphNodeType(compositeType),
	); err != nil {
		return nil, err
	}

	// Expect block opener
	blockOpenerFragment, err := expectFragment(
		ctx.Input,
		nameFragment.End,
		FragmentTypeOpenerBlock,
	)
	if err != nil {
		return nil, err
	}

	// Parse properties
	properties := make(map[string]*DeclProperty)
	end = blockOpenerFragment.End
	for {
		next := readFragment(ctx.Input, end)
		if next.IsEOF() {
			return nil, &ErrSyntax{
				Fragment: next.Fragment,
				Err: fmt.Errorf(
					"unexpected end of file, expected closure of %s block",
					compositeType.String(),
				),
			}
		}
		if next.FragmentType == FragmentTypeCloserBlock {
			// End of the properties block
			end = next.End
			break
		}

		// Parse property
		property, err := parseProperty(ctx, end, compositeType)
		if err != nil {
			return nil, err
		}

		// Shift the cursor
		end = property.Fragment.End

		properties[string(property.Name)] = property
	}

	// Determine basic expression information
	var result AbstractTypeDeclaration
	baseExpression := BaseExpression{
		Fragment: Fragment{
			Raw:   src[start.Index:end.Index],
			File:  ctx.Input,
			Start: start,
			End:   end,
		},
	}
	baseDocumented := BaseDocumented{}
	baseType := BaseDeclaration{VName: nameFragment.Raw}
	baseComposite := BaseComposite{Properties: properties}

	switch compositeType {
	case GNTCStruct:
		result = &DeclStruct{
			BaseExpression:  baseExpression,
			BaseDocumented:  baseDocumented,
			BaseDeclaration: baseType,
			BaseComposite:   baseComposite,
		}
	case GNTCEntity:
		result = &DeclEntity{
			BaseExpression:  baseExpression,
			BaseDocumented:  baseDocumented,
			BaseDeclaration: baseType,
			BaseComposite:   baseComposite,
		}
	case GNTCUser:
		result = &DeclUser{
			DeclEntity: DeclEntity{
				BaseExpression:  baseExpression,
				BaseDocumented:  baseDocumented,
				BaseDeclaration: baseType,
				BaseComposite:   baseComposite,
			},
		}
	}

	// Try register the parsed composite type
	if err := ctx.ModelCtx.register.tryRegisterType(result); err != nil {
		return nil, err
	}

	return result, nil
}
