package ast_test

import (
	"reflect"
	"smlp/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

// RegisterExclusive expects the given declarations exclusively
func RegisterExclusive(
	t *testing.T,
	register ast.Register,
	declarations ...ast.AbstractDeclaration,
) {
	for _, declaration := range declarations {
		found := register.FindGlobalDeclaration(string(declaration.Name()))
		require.NotNil(t, found)

		switch declaration := declaration.(type) {
		case *ast.DeclError:
			//TODO: implement
			t.Fatal("not yet implemented")
		case *ast.DeclTransaction:
			//TODO: implement
			t.Fatal("not yet implemented")
		case *ast.DeclScalar:
			require.IsType(t, &ast.DeclScalar{}, found)
			require.Equal(t, declaration, found.(*ast.DeclScalar))
		case *ast.DeclEnum:
			require.IsType(t, &ast.DeclEnum{}, found)
			require.Equal(t, declaration, found.(*ast.DeclEnum))
		case *ast.DeclStruct:
			require.IsType(t, &ast.DeclStruct{}, found)
			require.Equal(t, declaration, found.(*ast.DeclStruct))
		case *ast.DeclEntity:
			require.IsType(t, &ast.DeclEntity{}, found)
			require.Equal(t, declaration, found.(*ast.DeclEntity))
		case *ast.DeclUser:
			require.IsType(t, &ast.DeclUser{}, found)
			require.Equal(t, declaration, found.(*ast.DeclUser))
		default:
			t.Fatalf("unexpected type: %s", reflect.TypeOf(declaration))
		}
	}
}
