package ast

import "fmt"

func parseFile(ctx *contextFile) (SourceFile, error) {
	parsedSourceFile := NewSourceFile(ctx.Input)

	src := ctx.Source()
	if len(src) < 7 {
		// file contents too short to fit the minimal package identifier
		return SourceFile{}, fmt.Errorf(
			"missing model identifier in %s",
			string(ctx.Input.FullPath),
		)
	}

	crs := Cursor{}

	// Parse package identifier
	modelName, fragModelident, err := parseModelIdent(ctx, crs)
	if err != nil {
		return SourceFile{}, err
	}

	// Try to define model name
	if !ctx.ModelCtx.TrySetModelName(modelName) {
		//TODO: use specialized error type (mismatching model identifiers)
		return SourceFile{}, &ErrSyntax{
			Fragment: fragModelident,
			Err: fmt.Errorf(
				"unexpected model name: %s",
				string(modelName),
			),
		}
	}

	end := fragModelident.End
	for {
		// Parse expressions until the end of the file
		fragHead := readSequence(ctx.Input, end)

		var expr AbstractExpression

		if fragHead.IsEOF() {
			// End of file
			break
		} else if hasPrefix(fragHead.Raw, str("//")) {
			// Comment expression
			comment, err := parseComment(ctx, fragHead.Start)
			if err != nil {
				return SourceFile{}, err
			}
			expr = comment
		} else if hasPrefix(fragHead.Raw, str("#")) {
			// Documentation expression
			documented, err := parseDocumentedExpression(ctx, fragHead.Start)
			if err != nil {
				return SourceFile{}, err
			}
			expr = documented

		} else if equal(fragHead.Raw, keywordScalar) {
			// Scalar type declaration
			scalarType, err := parseTypeScalar(ctx, fragHead.Start)
			if err != nil {
				return SourceFile{}, err
			}
			expr = scalarType

		} else if equal(fragHead.Raw, keywordEnum) {
			// Enum type declaration
			enumType, err := parseTypeEnum(ctx, fragHead.Start)
			if err != nil {
				return SourceFile{}, err
			}
			expr = enumType

		} else if equal(fragHead.Raw, keywordStruct) {
			// Struct type declaration
			structT, err := parseTypeComposite(ctx, fragHead.Start, GNTCStruct)
			if err != nil {
				return SourceFile{}, err
			}
			expr = structT.(*DeclStruct)

		} else if equal(fragHead.Raw, keywordEntity) {
			// Entity type declaration
			entityT, err := parseTypeComposite(ctx, fragHead.Start, GNTCEntity)
			if err != nil {
				return SourceFile{}, err
			}
			expr = entityT.(*DeclEntity)

		} else if equal(fragHead.Raw, keywordUser) {
			// User entity type declaration
			userT, err := parseTypeComposite(ctx, fragHead.Start, GNTCUser)
			if err != nil {
				return SourceFile{}, err
			}
			expr = userT.(*DeclUser)

		} else if equal(fragHead.Raw, keywordUser) {
			// Error type declaration
			errorT, err := parseTypeError(ctx, fragHead.Start)
			if err != nil {
				return SourceFile{}, err
			}
			expr = errorT

		} else if equal(fragHead.Raw, keywordCache) {
			// Cache control declaration
			attrCache, err := parseAttrCache(ctx, fragHead.Start)
			if err != nil {
				return SourceFile{}, err
			}
			expr = attrCache

		} else if equal(fragHead.Raw, keywordModel) {
			// DeclModel declaration
			modelDecl, err := parseModelDecl(ctx, fragHead.Start)
			if err != nil {
				return SourceFile{}, err
			}
			expr = modelDecl

		} else if equal(fragHead.Raw, keywordAccess) {
			// Access permissions attribute declaration
			attrAccess, err := parseAttrAccess(ctx, fragHead.Start)
			if err != nil {
				return SourceFile{}, err
			}
			expr = attrAccess
		} else {
			return SourceFile{}, &ErrSyntax{
				Fragment: fragHead,
				Err: fmt.Errorf(
					"unexpected expression: '%s'",
					string(fragHead.Raw),
				),
			}
		}

		// Shift the cursor
		end = expr.SourceFragment().End
	}

	return parsedSourceFile, nil
}
