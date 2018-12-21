package ast

// DeclUser represents a user entity type declaration
type DeclUser struct {
	DeclEntity
}

// GraphNodeType implements the AbstractTypeDeclaration interface
func (tp *DeclUser) GraphNodeType() GraphNodeType {
	return GNTUser
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclUser) ExpressionType() ExpressionType {
	return ExprDeclUser
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclUser) DeclarationType() DeclarationType {
	return DeclarationTypeUser
}

// GlobalDeclarationType implements the AbstractGlobalDeclaration interface
func (tp *DeclUser) GlobalDeclarationType() GDeclarationType {
	return GDeclarationTypeUser
}
