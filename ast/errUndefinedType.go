package ast

import "fmt"

// ErrUndefinedType represents an error indicating undefined type references
type ErrUndefinedType struct {
	References []*Fragment
	TypeName   string
}

// Error implements the error interface
func (err *ErrUndefinedType) Error() string {
	references := make([]string, len(err.References))
	refIdx := 0
	for _, ref := range err.References {
		references[refIdx] = ref.Trace()
		refIdx++
	}

	return fmt.Sprintf(
		"undefined type %s referenced in %v",
		err.TypeName,
		err.References,
	)
}
