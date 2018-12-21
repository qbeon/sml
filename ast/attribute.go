package ast

// Attribute represents a basic attribute structure
type Attribute struct {
	Params []*DeclParameter
}

// Parameters implements the AbstractAttributeDeclaration interface
func (attr *Attribute) Parameters() []*DeclParameter {
	return attr.Params
}
