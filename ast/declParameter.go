package ast

// DeclParameter represents an attribute parameter field expression
type DeclParameter struct {
	BaseExpression
	BaseDocumented
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclParameter) ExpressionType() ExpressionType {
	return ExprDeclParam
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclParameter) DeclarationType() DeclarationType {
	return DeclarationParameter
}
