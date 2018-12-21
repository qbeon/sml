package ast

// BaseExpression represents a basic naked expression
type BaseExpression struct {
	Fragment
}

// ExpressionType implements the AbstractExpression interface
func (expr *BaseExpression) ExpressionType() ExpressionType {
	return ExpressionType(0)
}

// SourceFragment implements the AbstractExpression interface
func (expr *BaseExpression) SourceFragment() Fragment {
	return expr.Fragment
}
