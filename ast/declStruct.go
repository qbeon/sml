package ast

// DeclStruct represents a struct type declaration
type DeclStruct struct {
	BaseExpression
	BaseDocumented
	BaseDeclaration
	BaseComposite
}

// GraphNodeType implements the AbstractTypeDeclaration interface
func (tp *DeclStruct) GraphNodeType() GraphNodeType {
	return GNTStruct
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclStruct) ExpressionType() ExpressionType {
	return ExprDeclStruct
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclStruct) DeclarationType() DeclarationType {
	return DeclarationTypeStruct
}

// GlobalDeclarationType implements the AbstractGlobalDeclaration interface
func (tp *DeclStruct) GlobalDeclarationType() GDeclarationType {
	return GDeclarationTypeStruct
}
