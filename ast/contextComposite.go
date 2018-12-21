package ast

import "fmt"

func newContextComposite(
	fileCtx *contextFile,
	nodeType GraphNodeTypeComposite,
	entityTypeName str,
) *contextComposite {
	return &contextComposite{
		FileCtx:    fileCtx,
		NodeType:   nodeType,
		TypeName:   entityTypeName,
		Properties: make(map[*DeclProperty]struct{}),
	}
}

// contextComposite represents the context of a composite type parser
type contextComposite struct {
	FileCtx    *contextFile
	NodeType   GraphNodeTypeComposite
	TypeName   str
	Properties map[*DeclProperty]struct{}
}

// RootModelContext implements the context interface
func (ctx *contextComposite) RootModelContext() *contextModel {
	return ctx.FileCtx.RootModelContext()
}

// Source implements the context interface
func (ctx *contextComposite) Source() str {
	return ctx.FileCtx.Source()
}

// HasProperty returns true if a property with the given name is already
// defined, otherwise returns false
func (ctx *contextComposite) HasProperty(name str) bool {
	for prop := range ctx.Properties {
		if equal(prop.Name, name) {
			return true
		}
	}
	return false
}

// DeclareProperty declares a new property
func (ctx *contextComposite) DeclareProperty(prop *DeclProperty) error {
	if ctx.HasProperty(prop.Name) {
		return fmt.Errorf("property %s al", string(prop.Name))
	}
	ctx.Properties[prop] = struct{}{}
	return nil
}
