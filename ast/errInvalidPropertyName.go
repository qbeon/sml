package ast

import "fmt"

// ErrInvalidPropertyName represents an error indicating an invalid property
// name
type ErrInvalidPropertyName struct {
	Fragment  *Fragment
	Violation string
}

// Error implements the error interface
func (err *ErrInvalidPropertyName) Error() string {
	if len(err.Violation) < 1 {
		return fmt.Sprintf(
			"%s: invalid property name: '%s'",
			err.Fragment.Trace(),
			string(err.Fragment.Raw),
		)
	}
	return fmt.Sprintf(
		"%s: invalid property name: '%s', %s",
		err.Fragment.Trace(),
		string(err.Fragment.Raw),
		err.Violation,
	)
}
