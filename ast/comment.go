package ast

// Comment represents a source code comment expression
type Comment struct {
	BaseExpression
	Contents str
}

// ExpressionType implements the AbstractExpression interface
func (tp *Comment) ExpressionType() ExpressionType {
	return ExprComment
}
