package ast

// ExpressionType represents the type of an expression
type ExpressionType = uint8

const (
	_ ExpressionType = iota

	// ExprComment represents a source-code comment
	ExprComment

	// ExprDeclScalar represents a scalar type declaration expression
	ExprDeclScalar

	// ExprDeclEnum represents an enumeration type declaration expression
	ExprDeclEnum

	// ExprDeclStruct represents a structure type declaration expression
	ExprDeclStruct

	// ExprDeclEntity represents an entity declaration expression
	ExprDeclEntity

	// ExprDeclUser represents a user entity declaration expression
	ExprDeclUser

	// ExprDeclError represents an error type declaration expression
	ExprDeclError

	// ExprDeclTransact represents a transaction declaration expression
	ExprDeclTransact

	// ExprDeclErrorSet represents an error-set declaration expression
	ExprDeclErrorSet

	// ExprDeclAttrCache represents a cache attribute declaration expression
	ExprDeclAttrCache

	// ExprDeclAttrAccess represents an access control attribute declaration
	// expression
	ExprDeclAttrAccess

	// ExprDeclProperty represents a property declaration expression
	ExprDeclProperty

	// ExprDoc represents a documentation expression
	ExprDoc

	// ExprPropType represents a property type definition expression
	ExprPropType

	// ExprPropListType represents a list-type property type expression
	ExprPropListType

	// ExprPropOptionalType represents an optional property declaration
	// expression
	ExprPropOptionalType

	// ExprDeclModel represents a model declaration expression
	ExprDeclModel

	// ExprDeclParam represents an attribute parameter declaration
	ExprDeclParam

	// ExprDeclArgumentList represents an argument list declaration
	ExprDeclArgumentList
)

// AbstractExpression represents an abstract expression
type AbstractExpression interface {
	ExpressionType() ExpressionType
	SourceFragment() Fragment
}

// AbstractAttributeDeclaration represents an abstract attribute declaration
type AbstractAttributeDeclaration interface {
	AbstractDeclaration
	Parameters() []*DeclParameter
}

// AbstractTypeDeclaration represents an abstract type declaration
type AbstractTypeDeclaration interface {
	AbstractDeclaration
	GraphNodeType() GraphNodeType
}
