package ast

// GraphNodeType represents a graph node type
type GraphNodeType uint8

const (
	// GNTScalar represents a scalar type
	GNTScalar = GraphNodeType(ExprDeclScalar)

	// GNTEnum represents a enumeration type
	GNTEnum = GraphNodeType(ExprDeclEnum)

	// GNTStruct represents a struct type
	GNTStruct = GraphNodeType(ExprDeclStruct)

	// GNTEntity represents an entity type
	GNTEntity = GraphNodeType(ExprDeclEntity)

	// GNTUser represents a user entity type
	GNTUser = GraphNodeType(ExprDeclUser)
)

// String returns the string value
func (gnt GraphNodeType) String() string {
	switch gnt {
	case GNTScalar:
		return "scalar"
	case GNTEnum:
		return "enum"
	case GNTStruct:
		return "struct"
	case GNTEntity:
		return "entity"
	case GNTUser:
		return "user"
	}
	return ""
}

// GraphNodeTypeComposite represents a composite graph node type
type GraphNodeTypeComposite uint8

const (
	// GNTCStruct represents a struct type
	GNTCStruct = GraphNodeTypeComposite(GNTStruct)

	// GNTCEntity represents an entity type
	GNTCEntity = GraphNodeTypeComposite(GNTEntity)

	// GNTCUser represents a user entity type
	GNTCUser = GraphNodeTypeComposite(GNTUser)
)

// String returns the string value
func (gntc GraphNodeTypeComposite) String() string {
	switch gntc {
	case GNTCStruct:
		return "struct"
	case GNTCEntity:
		return "entity"
	case GNTCUser:
		return "user"
	}
	return ""
}
