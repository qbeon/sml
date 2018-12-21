package ast

// context represents an abstract context
type context interface {
	// RootModelContext returns the root model context
	RootModelContext() *contextModel
}
