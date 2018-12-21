package ast

func newContextModel() *contextModel {
	return &contextModel{
		register: newRegister(),
	}
}

// contextModel represents the context of a model parser
type contextModel struct {
	ModelName str
	register  *register
}

// TrySetModelName tries to set the model name and returns true if the given
// name was either set or is equal to the already set one. Returns false if the
// already set name differs from the given one
func (ctx *contextModel) TrySetModelName(name str) bool {
	if ctx.ModelName == nil {
		ctx.ModelName = name
		return true
	} else if equal(ctx.ModelName, name) {
		return true
	}
	return false
}
