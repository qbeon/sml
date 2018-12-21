package ast

// BaseComposite represents the basic information of a composite type (struct,
// entity, user entity)
type BaseComposite struct {
	Properties map[string]*DeclProperty
}
