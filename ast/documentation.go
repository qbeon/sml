package ast

// Documentation represents documentation of an expression
type Documentation struct {
	BaseExpression
	Contents str
}

// ExpressionType implements the AbstractExpression interface
func (tp *Documentation) ExpressionType() ExpressionType {
	return ExprDoc
}
