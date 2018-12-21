package ast

import "fmt"

// parseModelIdent parses a model identifier statement
func parseModelIdent(
	ctx *contextFile,
	crs Cursor,
) (str, *Fragment, error) { // model name, expression, error
	src := ctx.Source()
	fragKeyword := readSequence(ctx.Input, crs)
	if !equal(fragKeyword.Raw, keywordModel) {
		// Unexpected token
		return nil, nil, &ErrSyntax{
			Fragment: &Fragment{
				Raw:   nil,
				File:  ctx.Input,
				Start: fragKeyword.Start,
				End:   fragKeyword.End,
			},
			Err: fmt.Errorf(
				"unexpected token '%s', expected keyword: '%s'",
				string(fragKeyword.Raw),
				string(keywordModel),
			),
		}
	}

	fragModelName := readSequence(ctx.Input, fragKeyword.End)

	// Validate the model name
	if err := validateModelName(fragModelName); err != nil {
		return nil, nil, err
	}

	return fragModelName.Raw, &Fragment{
		Raw:   src[fragKeyword.Start.Index:fragModelName.End.Index],
		File:  ctx.Input,
		Start: fragKeyword.Start,
		End:   fragModelName.End,
	}, nil
}
