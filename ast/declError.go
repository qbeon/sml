package ast

// DeclError represents a scalar type declaration
type DeclError struct {
	BaseExpression
	BaseDocumented
	BaseDeclaration
}

// GraphNodeType implements the AbstractTypeDeclaration interface
func (tp *DeclError) GraphNodeType() GraphNodeType {
	return GNTEnum
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclError) ExpressionType() ExpressionType {
	return ExprDeclError
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclError) DeclarationType() DeclarationType {
	return DeclarationTypeError
}

// GlobalDeclarationType implements the AbstractGlobalDeclaration interface
func (tp *DeclError) GlobalDeclarationType() GDeclarationType {
	return GDeclarationTypeError
}
