package ast

import "fmt"

// ErrSyntax represents an error indicating a syntax error
type ErrSyntax struct {
	Fragment *Fragment
	Err      error
}

// Error implements the error interface
func (err *ErrSyntax) Error() string {
	if err.Fragment.IsZero() {
		return fmt.Sprintf("unknown parser error: %s", err.Err)
	}
	return fmt.Sprintf(
		"%s: unknown parser error: %s",
		err.Fragment.Trace(),
		err.Err,
	)
}
