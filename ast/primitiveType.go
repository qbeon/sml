package ast

import (
	"regexp"
	"time"
)

// PrimitiveType represents the type of a builtin primitive type
type PrimitiveType uint8

const (
	_ PrimitiveType = iota

	// PrimitiveBool represents a boolean value that's either true or false
	PrimitiveBool

	// PrimitiveByte represents a 1-byte value
	PrimitiveByte

	// PrimitiveInt represents a signed 32-bit integer value
	PrimitiveInt

	// PrimitiveUint represents an unsigned 32-bit integer value
	PrimitiveUint

	// PrimitiveDouble represents a double-precision 64-bit floating point
	// number value
	PrimitiveDouble

	// PrimitiveString represents a 7-bit ASCII string value
	PrimitiveString

	// PrimitiveText represents a unicode UTF-8 text string value
	PrimitiveText

	// PrimitiveTime represents a date-time value
	PrimitiveTime

	// PrimitiveBinary represents an arbitrary byte-array
	PrimitiveBinary
)

// String returns the textual representation of the value
func (pt PrimitiveType) String() string {
	switch pt {
	case PrimitiveBool:
		return "bool"
	case PrimitiveByte:
		return "byte"
	case PrimitiveInt:
		return "int"
	case PrimitiveUint:
		return "uint"
	case PrimitiveString:
		return "string"
	case PrimitiveText:
		return "text"
	case PrimitiveTime:
		return "time"
	case PrimitiveBinary:
		return "binary"
	}
	return ""
}

var primitiveTypeNames = map[string]struct{}{
	"Bool":   struct{}{},
	"Byte":   struct{}{},
	"Int":    struct{}{},
	"Uint":   struct{}{},
	"Double": struct{}{},
	"String": struct{}{},
	"Text":   struct{}{},
	"Time":   struct{}{},
	"Binary": struct{}{},
}

func isPrimitiveType(typeName string) bool {
	_, found := primitiveTypeNames[typeName]
	return found
}

// PrimitiveConstraintsInt represents the constraints of a primitive signed
// integer value
type PrimitiveConstraintsInt struct {
	Min int32
	Max int32
}

// PrimitiveConstraintsUInt represents the constraints of a primitive unsigned
// integer value
type PrimitiveConstraintsUInt struct {
	Min uint32
	Max uint32
}

// PrimitiveConstraintsDouble represents the constraints of a primitive
// double-precision floating-pointer number value
type PrimitiveConstraintsDouble struct {
	Min float64
	Max float64
}

// PrimitiveConstraintsString represents the constraints of a primitive string
// value
type PrimitiveConstraintsString struct {
	Regex regexp.Regexp
}

// PrimitiveConstraintsText represents the constraints of a primitive string
// value
type PrimitiveConstraintsText struct {
	Regex regexp.Regexp
}

// PrimitiveConstraintsTime represents the constraints of a primitive time value
type PrimitiveConstraintsTime struct {
	After  time.Time
	Before time.Time
}

// PrimitiveConstraintsBinary represents the constraints of a primitive binary
// data value
type PrimitiveConstraintsBinary struct {
	MinLen uint64
	MaxLen uint64
}
