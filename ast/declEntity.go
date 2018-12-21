package ast

// DeclEntity represents an entity type declaration
type DeclEntity struct {
	BaseExpression
	BaseDocumented
	BaseDeclaration
	BaseComposite

	// DeclRelations represents a map of relations to other entities
	DeclRelations map[*DeclEntity]*DeclRelation
}

// GraphNodeType implements the AbstractTypeDeclaration interface
func (tp *DeclEntity) GraphNodeType() GraphNodeType {
	return GNTEntity
}

// ExpressionType implements the AbstractExpression interface
func (tp *DeclEntity) ExpressionType() ExpressionType {
	return ExprDeclEntity
}

// DeclarationType implements the AbstractDeclaration interface
func (tp *DeclEntity) DeclarationType() DeclarationType {
	return DeclarationTypeEntity
}

// GlobalDeclarationType implements the AbstractGlobalDeclaration interface
func (tp *DeclEntity) GlobalDeclarationType() GDeclarationType {
	return GDeclarationTypeEntity
}
