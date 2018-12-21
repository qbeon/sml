package ast

// DeclAttributeCache represents a cache control attribute
type DeclAttributeCache struct {
	BaseExpression
	BaseDocumented
	BaseAttribute
	//TODO: implement cache specifier target
	Targets []interface{}
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclAttributeCache) ExpressionType() ExpressionType {
	return ExprDeclAttrCache
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclAttributeCache) DeclarationType() DeclarationType {
	return DeclarationAttrCache
}

// GlobalDeclarationType implements the AbstractGlobalDeclaration interface
func (tp *DeclAttributeCache) GlobalDeclarationType() GDeclarationType {
	return GDeclarationAttrCache
}
