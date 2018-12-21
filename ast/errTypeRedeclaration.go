package ast

import "fmt"

// ErrTypeRedeclaration represents a traced parser error indicating a
// redeclaration of a type name
type ErrTypeRedeclaration struct {
	Redeclaration           Fragment
	PreviousDeclaration     Fragment
	TypeName                string
	PreviousDeclarationType GraphNodeType
	RedeclarationType       GraphNodeType
}

// Error implements the error interface
func (err *ErrTypeRedeclaration) Error() string {
	return fmt.Sprintf(
		"%s: type %s (%s) redeclared as %s (previous declaration: %s)",
		err.Redeclaration.Trace(),
		err.TypeName,
		err.PreviousDeclarationType.String(),
		err.RedeclarationType.String(),
		err.PreviousDeclaration.Trace(),
	)
}
