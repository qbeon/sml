package ast

// BaseDeclaration represents basic declared type information
type BaseDeclaration struct {
	VName str
}

// Name implements the AbstractTypeDeclaration interface
func (tp *BaseDeclaration) Name() str {
	return tp.VName
}
