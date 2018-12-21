package ast_test

import (
	"smlp/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParseStructUndefinedPropType tests parsing a model with a single struct
// type declaration containing a property of an undefined type "U" expecting the
// parser to fail
func TestParseStructUndefinedPropType(t *testing.T) {
	modelSml := `model A struct S {prop Undefined}`
	input := []ast.InputFile{
		ast.InputFile{
			Name:     str("model"),
			Path:     str("/test/"),
			FullPath: str("/test/model.sml"),
			Contents: str(modelSml),
		},
	}
	syntaxTree, err := ast.Parse(input)
	require.Error(t, err)
	require.Nil(t, syntaxTree)

	require.IsType(t, &ast.ErrUndefinedType{}, err)
	errt := err.(*ast.ErrUndefinedType)
	require.Equal(t, "Undefined", errt.TypeName)
	require.Len(t, errt.References, 1)
	require.Equal(t, int(0), int(errt.References[0].Start.Line))
	require.Equal(t, int(23), int(errt.References[0].Start.Column))
	require.Equal(t, int(0), int(errt.References[0].Start.Line))
	require.Equal(t, int(32), int(errt.References[0].End.Column))
}
