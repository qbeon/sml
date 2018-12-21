package ast

// DeclEnum represents a scalar type declaration
type DeclEnum struct {
	BaseExpression
	BaseDocumented
	BaseDeclaration
}

// GraphNodeType implements the AbstractTypeDeclaration interface
func (tp *DeclEnum) GraphNodeType() GraphNodeType {
	return GNTEnum
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclEnum) ExpressionType() ExpressionType {
	return ExprDeclEnum
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclEnum) DeclarationType() DeclarationType {
	return DeclarationTypeEnum
}

// GlobalDeclarationType implements the AbstractGlobalDeclaration interface
func (tp *DeclEnum) GlobalDeclarationType() GDeclarationType {
	return GDeclarationTypeEnum
}
