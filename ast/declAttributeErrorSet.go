package ast

// DeclAttributeErrorSet represents an error-set transaction attribute
type DeclAttributeErrorSet struct {
	BaseExpression
	BaseDocumented
	BaseAttribute
	TargetTransactionName str
	TargetTransaction     *DeclTransaction
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclAttributeErrorSet) ExpressionType() ExpressionType {
	return ExprDeclErrorSet
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclAttributeErrorSet) DeclarationType() DeclarationType {
	return DeclarationErrorSet
}

// GlobalDeclarationType implements the AbstractGlobalDeclaration interface
func (tp *DeclAttributeErrorSet) GlobalDeclarationType() GDeclarationType {
	return GDeclarationErrorSet
}
