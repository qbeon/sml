package ast

import "fmt"

type str = []rune

// Fragment represents a code fragment
type Fragment struct {
	Raw   str
	File  *InputFile
	Start Cursor
	End   Cursor
}

// IsEOF returns true of the fragment points to EOF, otherwise returns false
func (frag *Fragment) IsEOF() bool {
	return frag.Raw == nil
}

// IsZero returns true of the fragment points to EOF, otherwise returns false
func (frag *Fragment) IsZero() bool {
	return frag.File == nil
}

// Trace returns the stringified fragment trace
func (frag *Fragment) Trace() string {
	return fmt.Sprintf(
		"%s:%d:%d",
		string(frag.File.FullPath),
		frag.Start.Line+1,
		frag.Start.Column+1,
	)
}

// AST represents the SML abstract syntax tree
type AST struct {
	SourceFiles []SourceFile
	Register    Register
}
