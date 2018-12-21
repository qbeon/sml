package ast

// DeclAttributeAccess represents an access control attribute
type DeclAttributeAccess struct {
	BaseExpression
	BaseDocumented
	BaseAttribute
	//TODO: implement access specifier target
	Targets []interface{}
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclAttributeAccess) ExpressionType() ExpressionType {
	return ExprDeclAttrAccess
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclAttributeAccess) DeclarationType() DeclarationType {
	return DeclarationAttrAccess
}

// GlobalDeclarationType implements the AbstractGlobalDeclaration interface
func (tp *DeclAttributeAccess) GlobalDeclarationType() GDeclarationType {
	return GDeclarationAttrAccess
}
