package ast

// Cursor represents a source-code cursor
type Cursor struct {
	// Index represents the linear char index
	Index uint

	// Line represents the file line (starts with 0)
	Line uint

	// Column represents the file column (starts with 0)
	Column uint
}

// IsEOF returns true if the cursor is pointing at the end of the given source,
// otherwise returns false
func (crs Cursor) IsEOF(src str) bool {
	return crs.Index >= uint(len(src))
}
