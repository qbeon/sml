package ast

func newContextFile(
	modelCtx *contextModel,
	input *InputFile,
) *contextFile {
	return &contextFile{
		ModelCtx: modelCtx,
		Input:    input,
	}
}

// contextFile represents the context of a file parser
type contextFile struct {
	ModelCtx *contextModel
	Input    *InputFile
}

// RootModelContext implements the context interface
func (ctx *contextFile) RootModelContext() *contextModel {
	return ctx.ModelCtx
}

// Source implements the context interface
func (ctx *contextFile) Source() str {
	return ctx.Input.Contents
}
