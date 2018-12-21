package ast

func newContextProperty(
	typeCtx *contextComposite,
	propName str,
) *contextProperty {
	return &contextProperty{
		TypeCtx:  typeCtx,
		PropName: propName,
	}
}

// contextProperty represents the context of property parser
type contextProperty struct {
	TypeCtx  *contextComposite
	PropName str
}

// RootModelContext implements the context interface
func (ctx *contextProperty) RootModelContext() *contextModel {
	return ctx.TypeCtx.RootModelContext()
}

// Source implements the context interface
func (ctx *contextProperty) Source() str {
	return ctx.TypeCtx.Source()
}
