package ast

import "fmt"

// ErrInvalidTypeName represents an error indicating an invalid type name
type ErrInvalidTypeName struct {
	Fragment  *Fragment
	Violation string
}

// Error implements the error interface
func (err *ErrInvalidTypeName) Error() string {
	if len(err.Violation) < 1 {
		return fmt.Sprintf(
			"%s: invalid type name: '%s'",
			err.Fragment.Trace(),
			string(err.Fragment.Raw),
		)
	}
	return fmt.Sprintf(
		"%s: invalid type name: '%s', %s",
		err.Fragment.Trace(),
		string(err.Fragment.Raw),
		err.Violation,
	)
}
