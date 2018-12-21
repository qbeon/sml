package ast

import "fmt"

// ErrInvalidModelName represents an error indicating an invalid model name
type ErrInvalidModelName struct {
	Fragment  *Fragment
	Violation string
}

// Error implements the error interface
func (err *ErrInvalidModelName) Error() string {
	if len(err.Violation) < 1 {
		return fmt.Sprintf(
			"%s: invalid model name: '%s'",
			err.Fragment.Trace(),
			string(err.Fragment.Raw),
		)
	}
	return fmt.Sprintf(
		"%s: invalid model name: '%s', %s",
		err.Fragment.Trace(),
		string(err.Fragment.Raw),
		err.Violation,
	)
}
