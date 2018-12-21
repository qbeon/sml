package ast

// DeclModel represents a parsed model
type DeclModel struct {
	BaseExpression
	BaseDocumented
}

// ExpressionType implements the AbstractExpression interface
func (*DeclModel) ExpressionType() ExpressionType {
	return ExprDeclModel
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclModel) DeclarationType() DeclarationType {
	return DeclarationModel
}

// GlobalDeclarationType implements the AbstractGlobalDeclaration interface
func (tp *DeclModel) GlobalDeclarationType() GDeclarationType {
	return GDeclarationModel
}
