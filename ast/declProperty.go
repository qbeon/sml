package ast

// DeclProperty represents a property declaration
type DeclProperty struct {
	BaseExpression
	BaseDocumented
	Name     str
	TypeName str
	TypeRef  AbstractTypeDeclaration
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclProperty) ExpressionType() ExpressionType {
	return ExprDeclProperty
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclProperty) DeclarationType() DeclarationType {
	return DeclarationProperty
}
