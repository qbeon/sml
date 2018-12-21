package ast

// DeclScalar represents a scalar type declaration
type DeclScalar struct {
	BaseExpression
	BaseDocumented
	BaseDeclaration
	//TODO: implement constraints
	Constraints []*DeclParameter
}

// GraphNodeType implements the AbstractTypeDeclaration interface
func (tp *DeclScalar) GraphNodeType() GraphNodeType {
	return GNTScalar
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclScalar) ExpressionType() ExpressionType {
	return ExprDeclScalar
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclScalar) DeclarationType() DeclarationType {
	return DeclarationTypeScalar
}

// GlobalDeclarationType implements the AbstractGlobalDeclaration interface
func (tp *DeclScalar) GlobalDeclarationType() GDeclarationType {
	return GDeclarationTypeScalar
}
