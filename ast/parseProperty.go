package ast

import "fmt"

// parseProperty parses a composite type property declaration statement
func parseProperty(
	ctx *contextFile,
	start Cursor,
	compositeType GraphNodeTypeComposite,
) (*DeclProperty, error) {
	src := ctx.Source()

	// Read name
	nameFragment := readLatinWord(ctx.Input, start)
	if nameFragment.IsEOF() {
		return nil, &ErrSyntax{
			Fragment: nameFragment,
			Err: fmt.Errorf(
				"missing %s name",
				compositeType.String(),
			),
		}
	}

	// Validate property name
	if err := validatePropertyName(
		nameFragment,
		GraphNodeType(compositeType),
	); err != nil {
		return nil, err
	}

	// Read type name
	typeNameFragment := readLatinWord(ctx.Input, nameFragment.End)
	if typeNameFragment.IsEOF() {
		return nil, &ErrSyntax{
			Fragment: typeNameFragment,
			Err: fmt.Errorf(
				"missing %s property type name",
				compositeType.String(),
			),
		}
	}

	// Register a reference to the type of the property
	ctx.ModelCtx.register.registerTypeReference(typeNameFragment)

	return &DeclProperty{
		BaseExpression: BaseExpression{
			Fragment: Fragment{
				Raw:   src[start.Index:typeNameFragment.End.Index],
				File:  ctx.Input,
				Start: start,
				End:   typeNameFragment.End,
			},
		},
		Name:     nameFragment.Raw,
		TypeName: typeNameFragment.Raw,
		//TODO: add TypeRef
	}, nil
}
