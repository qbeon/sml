package ast_test

import (
	"smlp/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

type str = []rune

// TestParseMinimal tests parsing a model with no declarations
func TestParseMinimal(t *testing.T) {
	modelSml := `model A`
	input := []ast.InputFile{
		ast.InputFile{
			Name:     str("model"),
			Path:     str("/test/"),
			FullPath: str("/test/model.sml"),
			Contents: str(modelSml),
		},
	}
	syntaxTree, parseErr := ast.Parse(input)
	require.Error(t, parseErr)
	require.IsType(t, ast.ErrSyntax{}, parseErr)
	require.Nil(t, syntaxTree)
}

// TestParseStruct tests parsing a model with a single struct type declaration
// containing a single property declaration of a primitive type
func TestParseStruct(t *testing.T) {
	modelSml := `model A struct S {prop Int}`
	inputFile := &ast.InputFile{
		Name:     str("model"),
		Path:     str("/test/"),
		FullPath: str("/test/model.sml"),
		Contents: str(modelSml),
	}
	input := []ast.InputFile{
		inputFile.Clone(),
	}
	syntaxTree, parseErr := ast.Parse(input)
	require.NoError(t, parseErr)
	require.NotNil(t, syntaxTree)

	// Construct expected declaration
	expectedStructDecl := &ast.DeclStruct{
		BaseExpression: ast.BaseExpression{
			Fragment: ast.Fragment{
				Raw:  str("struct S {prop Int}"),
				File: inputFile,
				Start: ast.Cursor{
					Index:  uint(8),
					Column: uint(8),
					Line:   uint(0),
				},
				End: ast.Cursor{
					Index:  uint(27),
					Column: uint(27),
					Line:   uint(0),
				},
			},
		},
		BaseComposite: ast.BaseComposite{
			Properties: map[string]*ast.DeclProperty{
				"prop": &ast.DeclProperty{
					BaseExpression: ast.BaseExpression{
						Fragment: ast.Fragment{
							Raw:  str("prop Int"),
							File: inputFile,
							Start: ast.Cursor{
								Index:  18,
								Column: 18,
								Line:   0,
							},
							End: ast.Cursor{
								Index:  26,
								Column: 26,
								Line:   0,
							},
						},
					},
					Name:     str("prop"),
					TypeName: str("Int"),
					TypeRef:  nil,
				},
			},
		},
		BaseDeclaration: ast.BaseDeclaration{
			VName: str("S"),
		},
		BaseDocumented: ast.BaseDocumented{},
	}

	require.Len(t, syntaxTree.SourceFiles, 1)
	//RegisterExclusive(t, syntaxTree.SourceFiles[0], expectedStructDecl)
	RegisterExclusive(t, syntaxTree.Register, expectedStructDecl)
}
