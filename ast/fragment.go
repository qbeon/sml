package ast

// FragmentType represents the type of a fragment
type FragmentType uint8

const (
	_ FragmentType = iota

	// FragmentTypeOpenerBlock represents "{"
	FragmentTypeOpenerBlock

	// FragmentTypeCloserBlock represents "}"
	FragmentTypeCloserBlock

	// FragmentTypeOpenerArgumentList represents "("
	FragmentTypeOpenerArgumentList

	// FragmentTypeCloserArgumentList represents ")"
	FragmentTypeCloserArgumentList

	// FragmentTypeDelimiterList represents ","
	FragmentTypeDelimiterList

	// FragmentTypeSpecifierOptional represents "?"
	FragmentTypeSpecifierOptional

	// FragmentTypeSpecifierList represents "[]"
	FragmentTypeSpecifierList

	// FragmentTypeCommentLine represents "//"
	FragmentTypeCommentLine

	// FragmentTypeDocumentationLine represents "#"
	FragmentTypeDocumentationLine
)

// String returns the name of the fragment type
func (tf FragmentType) String() string {
	switch tf {
	case FragmentType(0):
		return "expression"
	case FragmentTypeOpenerBlock:
		return "block-opener"
	case FragmentTypeCloserBlock:
		return "block-closer"
	case FragmentTypeOpenerArgumentList:
		return "argument-list-opener"
	case FragmentTypeCloserArgumentList:
		return "argument-list-closer"
	case FragmentTypeDelimiterList:
		return "list-delimiter"
	case FragmentTypeSpecifierOptional:
		return "optional-type-specifier"
	case FragmentTypeSpecifierList:
		return "list-type-specifier"
	case FragmentTypeCommentLine:
		return "code-comment-line-initializer"
	case FragmentTypeDocumentationLine:
		return "documentation-line-initializer"
	}
	return ""
}

const fragmentOpenerBlockRep = rune('{')
const fragmentCloserBlockRep = rune('}')
const fragmentOpenerArgumentListRep = rune('(')
const fragmentCloserArgumentListRep = rune(')')
const fragmentDelimiterListRep = rune(',')
const fragmentTypeSpecifierOptionalRep = rune('?')
const fragmentDocumentationLineRep = rune('#')

var fragmentTypeSpecifierListRep = str("[]")
var fragmentCommentLineRep = str("//")

// Representation returns the textual representation of the fragment type
func (tf FragmentType) Representation() string {
	switch tf {
	case FragmentType(0):
		return "[expression]"
	case FragmentTypeOpenerBlock:
		return string(fragmentOpenerBlockRep)
	case FragmentTypeCloserBlock:
		return string(fragmentCloserBlockRep)
	case FragmentTypeOpenerArgumentList:
		return string(fragmentOpenerArgumentListRep)
	case FragmentTypeCloserArgumentList:
		return string(fragmentCloserArgumentListRep)
	case FragmentTypeDelimiterList:
		return string(fragmentDelimiterListRep)
	case FragmentTypeSpecifierOptional:
		return string(fragmentTypeSpecifierOptionalRep)
	case FragmentTypeSpecifierList:
		return string(fragmentTypeSpecifierListRep)
	case FragmentTypeCommentLine:
		return string(fragmentCommentLineRep)
	case FragmentTypeDocumentationLine:
		return string(fragmentDocumentationLineRep)
	}
	return ""
}

// TypedFragment represents a typed fragment
type TypedFragment struct {
	*Fragment
	FragmentType FragmentType
}

// IsEOF returns true if the fragment points to EOF, otherwise returns false
func (tf TypedFragment) IsEOF() bool {
	return tf.Fragment == nil
}
