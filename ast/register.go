package ast

import (
	"fmt"
	"reflect"
)

// Register represents an structure information register
type Register interface {
	FindGlobalDeclaration(name string) AbstractGlobalDeclaration
	FindType(name string) AbstractTypeDeclaration

	Declarations() map[string]AbstractDeclaration
	Types() map[string]AbstractTypeDeclaration
	TypeReferences() map[string][]*Fragment
	TypesScalar() map[*DeclScalar]struct{}
	TypesEnum() map[*DeclEnum]struct{}
	TypesStruct() map[*DeclStruct]struct{}
	TypesEntity() map[*DeclEntity]struct{}
	TypesUser() map[*DeclUser]struct{}
	AttrsAccess() map[*DeclAttributeAccess]struct{}
	AttrsCache() map[*DeclAttributeCache]struct{}
	Transactions() map[*DeclTransaction]struct{}
}

func newRegister() *register {
	return &register{
		typeReferences: make(map[string][]*Fragment),
		declarations:   make(map[string]AbstractDeclaration),
		types:          make(map[string]AbstractTypeDeclaration),
		undefinedTypes: make(map[string]struct{}),
		typesScalar:    make(map[*DeclScalar]struct{}),
		typesEnum:      make(map[*DeclEnum]struct{}),
		typesStruct:    make(map[*DeclStruct]struct{}),
		typesEntity:    make(map[*DeclEntity]struct{}),
		typesUser:      make(map[*DeclUser]struct{}),
		attrsAccess:    make(map[*DeclAttributeAccess]struct{}),
		attrsCache:     make(map[*DeclAttributeCache]struct{}),
		transaction:    make(map[*DeclTransaction]struct{}),
	}
}

// register represents a declaration register
type register struct {
	typeReferences map[string][]*Fragment
	declarations   map[string]AbstractDeclaration
	types          map[string]AbstractTypeDeclaration
	undefinedTypes map[string]struct{}
	typesScalar    map[*DeclScalar]struct{}
	typesEnum      map[*DeclEnum]struct{}
	typesStruct    map[*DeclStruct]struct{}
	typesEntity    map[*DeclEntity]struct{}
	typesUser      map[*DeclUser]struct{}
	attrsAccess    map[*DeclAttributeAccess]struct{}
	attrsCache     map[*DeclAttributeCache]struct{}
	transaction    map[*DeclTransaction]struct{}
}

// FindGlobalDeclaration implements the Register interface
func (reg *register) FindGlobalDeclaration(
	name string,
) AbstractGlobalDeclaration {
	if len(name) < 1 {
		return nil
	}
	if decl, found := reg.declarations[name]; found {
		if decl, ok := decl.(AbstractGlobalDeclaration); ok {
			return decl
		}
	}
	return nil
}

// FindType looks for a type declaration given the type name. Returns nil
// if none is found
func (reg *register) FindType(name string) AbstractTypeDeclaration {
	if len(name) < 1 {
		return nil
	}
	if tp, found := reg.types[name]; found {
		return tp
	}
	return nil
}

// registerTypeReference registers a type reference
func (reg *register) registerTypeReference(frag *Fragment) {
	typeNameStr := string(frag.Raw)
	if list, exists := reg.typeReferences[typeNameStr]; exists {
		// Append to reference list
		reg.typeReferences[typeNameStr] = append(list, frag)
	} else {
		// Initialize reference list
		reg.typeReferences[typeNameStr] = []*Fragment{frag}
	}
	// Mark type as undefined if it's yet unknown and not a primitive type
	if !isPrimitiveType(typeNameStr) {
		reg.undefinedTypes[typeNameStr] = struct{}{}
	}
}

// tryRegisterType tries to register a new type
func (reg *register) tryRegisterType(
	typeDecl AbstractTypeDeclaration,
) error {
	// Check for redeclaration
	if found := reg.FindType(string(typeDecl.Name())); found != nil {
		return &ErrTypeRedeclaration{
			Redeclaration:           typeDecl.SourceFragment(),
			PreviousDeclaration:     found.SourceFragment(),
			TypeName:                string(typeDecl.Name()),
			PreviousDeclarationType: found.GraphNodeType(),
			RedeclarationType:       typeDecl.GraphNodeType(),
		}
	}

	// Register a new type
	reg.types[string(typeDecl.Name())] = typeDecl
	switch tp := typeDecl.(type) {
	case *DeclScalar:
		reg.typesScalar[tp] = struct{}{}
	case *DeclEnum:
		reg.typesEnum[tp] = struct{}{}
	case *DeclStruct:
		reg.typesStruct[tp] = struct{}{}
	case *DeclEntity:
		reg.typesEntity[tp] = struct{}{}
	case *DeclUser:
		reg.typesUser[tp] = struct{}{}
	default:
		panic(fmt.Errorf("unexpected type: %s", reflect.TypeOf(typeDecl)))
	}

	reg.declarations[string(typeDecl.Name())] = typeDecl

	return nil
}

// PostProcess analyzes and links the model structure. It checks for references
// to undefined types
func (reg *register) PostProcess() error {
	// Define undefined types if possible
	for undefinedTypeName := range reg.undefinedTypes {
		if _, inReg := reg.types[undefinedTypeName]; inReg {
			delete(reg.undefinedTypes, undefinedTypeName)
		}
	}

	// Check for and report undefined types
	for undefinedTypeName := range reg.undefinedTypes {
		references := reg.typeReferences[undefinedTypeName]

		return &ErrUndefinedType{
			TypeName:   undefinedTypeName,
			References: references,
		}
	}

	//TODO: Check for struct recursions
	//TODO:

	return nil
}

// Declarations implements the Register interface
func (reg *register) Declarations() map[string]AbstractDeclaration {
	return reg.declarations
}

// Types implements the Register interface
func (reg *register) Types() map[string]AbstractTypeDeclaration {
	return reg.types
}

// TypeReferences implements the Register interface
func (reg *register) TypeReferences() map[string][]*Fragment {
	return reg.typeReferences
}

// TypesScalar implements the Register interface
func (reg *register) TypesScalar() map[*DeclScalar]struct{} {
	return reg.typesScalar
}

// TypesEnum implements the Register interface
func (reg *register) TypesEnum() map[*DeclEnum]struct{} {
	return reg.typesEnum
}

// TypesStruct implements the Register interface
func (reg *register) TypesStruct() map[*DeclStruct]struct{} {
	return reg.typesStruct
}

// TypesEntity implements the Register interface
func (reg *register) TypesEntity() map[*DeclEntity]struct{} {
	return reg.typesEntity
}

// TypesUser implements the Register interface
func (reg *register) TypesUser() map[*DeclUser]struct{} {
	return reg.typesUser
}

// AttrsAccess implements the Register interface
func (reg *register) AttrsAccess() map[*DeclAttributeAccess]struct{} {
	return reg.attrsAccess
}

// AttrsCache implements the Register interface
func (reg *register) AttrsCache() map[*DeclAttributeCache]struct{} {
	return reg.attrsCache
}

// Transactions implements the Register interface
func (reg *register) Transactions() map[*DeclTransaction]struct{} {
	return reg.transaction
}
