package ast

// SourceFile represents a parsed source file
type SourceFile struct {
	register *register
	Original *InputFile
	Register Register
}

// NewSourceFile creates a new source file instance
func NewSourceFile(in *InputFile) SourceFile {
	reg := newRegister()
	return SourceFile{
		register: reg,
		Original: in,
		Register: reg,
	}
}
