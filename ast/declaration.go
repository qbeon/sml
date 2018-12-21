package ast

// DeclarationType represents the type of a declaration
type DeclarationType uint8

const (
	_ DeclarationType = 0

	/* GLOBAL */

	// DeclarationTypeScalar represents the declaration of a scalar type
	DeclarationTypeScalar = DeclarationType(ExprDeclScalar)

	// DeclarationTypeEnum represents the declaration of an enumeration type
	DeclarationTypeEnum = DeclarationType(ExprDeclEnum)

	// DeclarationTypeStruct represents the declaration of a struct type
	DeclarationTypeStruct = DeclarationType(ExprDeclStruct)

	// DeclarationTypeEntity represents the declaration of an entity
	DeclarationTypeEntity = DeclarationType(ExprDeclEntity)

	// DeclarationTypeUser represents the declaration of a user entity
	DeclarationTypeUser = DeclarationType(ExprDeclUser)

	// DeclarationErrorSet represents the declaration of an error set
	DeclarationErrorSet = DeclarationType(ExprDeclErrorSet)

	// DeclarationTypeError represents the declaration of an error type
	DeclarationTypeError = DeclarationType(ExprDeclError)

	// DeclarationTransaction represents the declaration of a transaction
	DeclarationTransaction = DeclarationType(ExprDeclTransact)

	// DeclarationAttrCache represents the declaration of a cache attribute
	DeclarationAttrCache = DeclarationType(ExprDeclAttrCache)

	// DeclarationAttrAccess represents the declaration of an access attribute
	DeclarationAttrAccess = DeclarationType(ExprDeclAttrAccess)

	// DeclarationModel represents the declaration of a model
	DeclarationModel = DeclarationType(ExprDeclModel)

	/* LOCAL */

	// DeclarationProperty represents the declaration of a property
	DeclarationProperty = DeclarationType(ExprDeclProperty)

	// DeclarationParameter represents the declaration of a parameter
	DeclarationParameter = DeclarationType(ExprDeclParam)
)

// GDeclarationType represents the type of a  global declaration
type GDeclarationType uint8

const (
	_ GDeclarationType = 0

	// GDeclarationTypeScalar represents the declaration of a scalar type
	GDeclarationTypeScalar = GDeclarationType(DeclarationTypeScalar)

	// GDeclarationTypeEnum represents the declaration of an enumeration type
	GDeclarationTypeEnum = GDeclarationType(DeclarationTypeEnum)

	// GDeclarationTypeStruct represents the declaration of a struct type
	GDeclarationTypeStruct = GDeclarationType(DeclarationTypeStruct)

	// GDeclarationTypeEntity represents the declaration of an entity
	GDeclarationTypeEntity = GDeclarationType(DeclarationTypeEntity)

	// GDeclarationTypeUser represents the declaration of a user entity
	GDeclarationTypeUser = GDeclarationType(DeclarationTypeUser)

	// GDeclarationErrorSet represents the declaration of an error set
	GDeclarationErrorSet = GDeclarationType(DeclarationErrorSet)

	// GDeclarationTypeError represents the declaration of an error type
	GDeclarationTypeError = GDeclarationType(DeclarationTypeError)

	// GDeclarationTransaction represents the declaration of a transaction
	GDeclarationTransaction = GDeclarationType(DeclarationTransaction)

	// GDeclarationAttrCache represents the declaration of a cache attribute
	GDeclarationAttrCache = GDeclarationType(DeclarationAttrCache)

	// GDeclarationAttrAccess represents the declaration of an access attribute
	GDeclarationAttrAccess = GDeclarationType(DeclarationAttrAccess)

	// GDeclarationModel represents the declaration of a model
	GDeclarationModel = GDeclarationType(DeclarationModel)
)

// AbstractDeclaration represents an abstract declaration
type AbstractDeclaration interface {
	AbstractExpression
	Name() str
	DeclarationType() DeclarationType
}

// AbstractGlobalDeclaration represents an abstract global declaration
type AbstractGlobalDeclaration interface {
	AbstractDeclaration
	GlobalDeclarationType() GDeclarationType
}
