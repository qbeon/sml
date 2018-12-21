package ast

// DeclTransaction represents a transaction declaration
type DeclTransaction struct {
	BaseExpression
	BaseDocumented
	BaseDeclaration
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclTransaction) ExpressionType() ExpressionType {
	return ExprDeclTransact
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclTransaction) DeclarationType() DeclarationType {
	return DeclarationTransaction
}

// GlobalDeclarationType implements the AbstractGlobalDeclaration interface
func (tp *DeclTransaction) GlobalDeclarationType() GDeclarationType {
	return GDeclarationTransaction
}
