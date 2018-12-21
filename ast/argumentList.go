package ast

// ArgumentList represents an argument list declaration
type ArgumentList struct {
	BaseExpression
}

// ExpressionType implements the AbstractExpression interface
func (tp *ArgumentList) ExpressionType() ExpressionType {
	return ExprDeclArgumentList
}
