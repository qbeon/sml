package ast

import (
	"fmt"
)

// skipEmpty skips all empty space such as whitespaces, tabs and line breaks
func skipEmpty(src str, start Cursor) Cursor {
	for {
		if start.Index >= uint(len(src)) {
			// End of source file reached
			return start
		}

		char := src[start.Index]
		if char == ' ' || char == '\t' {
			// space or tab
			start.Index++
			start.Column++
		} else if isLineBreak(start.Index, src) {
			// line-break (unix | windows)
			start.Index++
			start.Line++
			start.Column = 0
		} else {
			// non-empty character
			return start
		}
	}
}

// readSequence reads a character sequence until either any space character or
// the end of the file is reached. If the given cursor starts with a space
// character then all leading spaces are skiped
func readSequence(file *InputFile, start Cursor) *Fragment {
	start = skipEmpty(file.Contents, start)

	frag := Fragment{
		File:  file,
		Start: start,
		End:   start,
	}

	for {
		if frag.End.Index >= uint(len(file.Contents)) ||
			isSpace(frag.End.Index, file.Contents) {
			// End of source file or end of character sequence reached
			if frag.Start.Index != frag.End.Index {
				frag.Raw = file.Contents[frag.Start.Index:frag.End.Index]
			}
			return &frag
		}

		frag.End.Index++
		frag.End.Column++
	}
}

// readLatinWord reads a latin letter character sequence. If the given cursor
// starts with a space character then all leading spaces are skiped
func readLatinWord(file *InputFile, start Cursor) *Fragment {
	start = skipEmpty(file.Contents, start)

	frag := Fragment{
		File:  file,
		Start: start,
		End:   start,
	}

	for {
		if frag.End.Index >= uint(len(file.Contents)) ||
			!isLatinWordChar(file.Contents[frag.End.Index]) {
			// End of source file or end of word reached
			if frag.Start.Index != frag.End.Index {
				frag.Raw = file.Contents[frag.Start.Index:frag.End.Index]
			}
			return &frag
		}

		frag.End.Index++
		frag.End.Column++
	}
}

// readFragment reads a character fragment
func readFragment(file *InputFile, start Cursor) TypedFragment {
	start = skipEmpty(file.Contents, start)

	if start.IsEOF(file.Contents) {
		return TypedFragment{}
	}

	frag := TypedFragment{
		Fragment: &Fragment{
			File:  file,
			Raw:   file.Contents[start.Index : start.Index+1],
			Start: start,
			End: Cursor{
				Index:  start.Index + 1,
				Column: start.Column + 1,
				Line:   start.Line,
			},
		},
	}
	firstChar := file.Contents[start.Index]

	if firstChar == fragmentOpenerBlockRep {
		// Block opener
		frag.FragmentType = FragmentTypeOpenerBlock
		return frag

	} else if firstChar == fragmentCloserBlockRep {
		// Block closer
		frag.FragmentType = FragmentTypeCloserBlock
		return frag

	} else if firstChar == fragmentOpenerArgumentListRep {
		// Argument list opener
		frag.FragmentType = FragmentTypeOpenerArgumentList
		return frag

	} else if firstChar == fragmentCloserArgumentListRep {
		// Argument list closer
		frag.FragmentType = FragmentTypeCloserArgumentList
		return frag

	} else if firstChar == fragmentDelimiterListRep {
		// List delimiter
		frag.FragmentType = FragmentTypeDelimiterList
		return frag

	} else if firstChar == fragmentDocumentationLineRep {
		// Documentation line initializer
		frag.FragmentType = FragmentTypeDocumentationLine
		return frag

	} else if firstChar == fragmentTypeSpecifierOptionalRep {
		// Optional type specifier
		frag.FragmentType = FragmentTypeSpecifierOptional
		return frag

	} else if start.Index+2 <= uint(len(file.Contents)) {
		frag.Raw = file.Contents[start.Index : start.Index+2]
		frag.End = Cursor{
			Index:  start.Index + 2,
			Column: start.Column + 2,
			Line:   start.Line,
		}

		// Multi-character fragment
		charPair := file.Contents[start.Index : start.Index+2]
		if equal(charPair, fragmentTypeSpecifierListRep) {
			// List type specifier
			frag.FragmentType = FragmentTypeSpecifierList
		} else if equal(charPair, fragmentCommentLineRep) {
			// Comment line initializer
			frag.FragmentType = FragmentTypeCommentLine
		}
		return frag
	}

	// Fallback to untyped fragment
	for {
		if frag.End.IsEOF(file.Contents) ||
			isSpace(frag.End.Index, file.Contents) {
			// End of source file reached
			if frag.Start.Index != frag.End.Index {
				frag.Raw = nil
			}
			return frag
		}

		frag.End.Index++
		frag.End.Column++
	}
}

// expectFragment wraps the readFragment function and returns a syntax-error if
// the type of the read fragment doesn't match any of the expected ones
func expectFragment(
	file *InputFile,
	start Cursor,
	expected ...FragmentType,
) (TypedFragment, error) {
	frag := readFragment(file, start)

	for _, expectedFragType := range expected {
		if expectedFragType == frag.FragmentType {
			// The read fragment matched one of the expected ones
			return frag, nil
		}
	}

	// Generate a list of expected fragment type names for error report
	expectedTypeNames := make([]string, len(expected))
	for i, expected := range expected {
		expectedTypeNames[i] = expected.String()
	}

	return TypedFragment{}, &ErrSyntax{
		Fragment: frag.Fragment,
		Err: fmt.Errorf(
			"unexpected fragment: %s, expected any of: %s",
			frag.FragmentType.String(),
			expectedTypeNames,
		),
	}
}

// Parse parses the given input files and returns an abstract syntax tree
func Parse(input []InputFile) (*AST, error) {
	if len(input) < 1 {
		return nil, fmt.Errorf("missing input")
	}

	sourceFiles := make([]SourceFile, len(input))

	modelCtx := newContextModel()

	for i, inFile := range input {
		parsedSourceFile, err := parseFile(newContextFile(modelCtx, &inFile))
		if err != nil {
			return nil, err
		}
		sourceFiles[i] = parsedSourceFile
	}

	if err := modelCtx.register.PostProcess(); err != nil {
		return nil, err
	}

	//TODO: update source-file registers

	return &AST{
		SourceFiles: sourceFiles,
		Register:    modelCtx.register,
	}, nil
}
