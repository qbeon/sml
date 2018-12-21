package ast

// equal returns true if the given rune slices are deeply equal
func equal(a, b []rune) bool {
	// XNOR
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// hasPrefix returns true if s has the given prefix
func hasPrefix(s str, prefix str) bool {
	if len(prefix) > len(s) {
		return false
	}
	for i := range prefix {
		if prefix[i] != s[i] {
			return false
		}
	}
	return true
}

// isLineBreak returns true if the rune at the given position of the source
// represents a line-break character, otherwise returns false
func isLineBreak(at uint, src str) bool {
	char := src[at]
	if char == '\n' || (char == '\r' &&
		at+1 < uint(len(src)) && src[at+1] == '\n') {
		// line-break (unix | windows)
		return true
	}
	return false
}

// isSpace returns true if the rune at the given position of the source
// represents a space, tab or break-line character
func isSpace(at uint, src str) bool {
	char := src[at]
	if char == ' ' || char == '\t' || isLineBreak(at, src) {
		// line-break (unix | windows)
		return true
	}
	// non-empty character
	return false
}

// isLatinWordChar returns true if the given rune represents a latin letter
// character, otherwise returns false
func isLatinWordChar(char rune) bool {
	if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) {
		return true
	}
	return false
}
