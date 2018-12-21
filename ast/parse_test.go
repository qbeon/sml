package ast

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadSequence(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		src := &InputFile{Contents: []rune("")}

		frag := readSequence(src, Cursor{})

		require.True(t, frag.IsEOF())
		require.Nil(t, frag.Raw)
		require.Equal(t, Cursor{Index: 0, Column: 0, Line: 0}, frag.Start)
		require.Equal(t, Cursor{Index: 0, Column: 0, Line: 0}, frag.End)
	})

	t.Run("Simple", func(t *testing.T) {
		src := &InputFile{Contents: []rune("model")}

		frag := readSequence(src, Cursor{})

		require.False(t, frag.IsEOF())
		require.Equal(t, []rune("model"), frag.Raw)
		require.Equal(t, Cursor{Index: 0, Column: 0, Line: 0}, frag.Start)
		require.Equal(t, Cursor{Index: 5, Column: 5, Line: 0}, frag.End)

		frag2 := readSequence(src, frag.End)

		require.True(t, frag2.IsEOF())
		require.Nil(t, frag2.Raw)
		require.Equal(t, Cursor{Index: 5, Column: 5, Line: 0}, frag2.Start)
		require.Equal(t, Cursor{Index: 5, Column: 5, Line: 0}, frag2.End)
	})

	t.Run("Mixed", func(t *testing.T) {
		src := &InputFile{Contents: []rune("  model \nr ")}

		frag := readSequence(src, Cursor{})

		require.False(t, frag.IsEOF())
		require.Equal(t, []rune("model"), frag.Raw)
		require.Equal(t, Cursor{Index: 2, Column: 2, Line: 0}, frag.Start)
		require.Equal(t, Cursor{Index: 7, Column: 7, Line: 0}, frag.End)

		frag2 := readSequence(src, frag.End)

		require.False(t, frag2.IsEOF())
		require.Equal(t, []rune("r"), frag2.Raw)
		require.Equal(t, Cursor{Index: 9, Column: 0, Line: 1}, frag2.Start)
		require.Equal(t, Cursor{Index: 10, Column: 1, Line: 1}, frag2.End)

		frag3 := readSequence(src, frag2.End)

		require.True(t, frag3.IsEOF())
		require.Nil(t, frag3.Raw)
		require.Equal(
			t,
			Cursor{Index: uint(len(src.Contents)), Column: 2, Line: 1},
			frag3.Start,
		)
		require.Equal(
			t,
			Cursor{Index: uint(len(src.Contents)), Column: 2, Line: 1},
			frag3.End,
		)
	})

	t.Run("MultiLine", func(t *testing.T) {
		src := &InputFile{Contents: []rune("\n\nT")}

		frag := readSequence(src, Cursor{})

		require.False(t, frag.IsEOF())
		require.Equal(t, []rune("T"), frag.Raw)
		require.Equal(t, Cursor{Index: 2, Column: 0, Line: 2}, frag.Start)
		require.Equal(t, Cursor{Index: 3, Column: 1, Line: 2}, frag.End)

		frag2 := readSequence(src, frag.End)

		require.True(t, frag2.IsEOF())
		require.Nil(t, frag2.Raw)
		require.Equal(t, Cursor{Index: 3, Column: 1, Line: 2}, frag2.Start)
		require.Equal(t, Cursor{Index: 3, Column: 1, Line: 2}, frag2.End)
	})

	t.Run("SpecialCharacters", func(t *testing.T) {
		src := &InputFile{Contents: []rune("{([//#])}")}

		frag := readSequence(src, Cursor{})

		require.False(t, frag.IsEOF())
		require.Equal(t, []rune("{([//#])}"), frag.Raw)
		require.Equal(t, Cursor{Index: 0, Column: 0, Line: 0}, frag.Start)
		require.Equal(t, Cursor{Index: 9, Column: 9, Line: 0}, frag.End)

		frag2 := readSequence(src, frag.End)

		require.True(t, frag2.IsEOF())
		require.Nil(t, frag2.Raw)
		require.Equal(t, Cursor{Index: 9, Column: 9, Line: 0}, frag2.Start)
		require.Equal(t, Cursor{Index: 9, Column: 9, Line: 0}, frag2.End)
	})
}

func TestReadFragment(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		src := &InputFile{Contents: []rune("")}
		frag := readFragment(src, Cursor{})
		require.Nil(t, frag)
	})

	testFragment := func(
		t *testing.T,
		src str,
		expectedFragType FragmentType,
	) {
		srcf := &InputFile{Contents: src}
		frag := readFragment(srcf, Cursor{})

		require.NotNil(t, frag)
		require.Equal(t, expectedFragType, frag.FragmentType)
		require.False(t, frag.IsEOF())
		require.Equal(t, src, frag.Raw)
		require.Equal(t, Cursor{Index: 0, Column: 0, Line: 0}, frag.Start)
		require.Equal(
			t,
			Cursor{Index: uint(len(src)), Column: uint(len(src)), Line: 0},
			frag.End,
		)

		require.Nil(t, readFragment(srcf, frag.End))
	}

	t.Run("BlockOpener", func(t *testing.T) {
		testFragment(t, []rune("{"), FragmentTypeOpenerBlock)
	})

	t.Run("BlockCloser", func(t *testing.T) {
		testFragment(t, []rune("}"), FragmentTypeCloserBlock)
	})

	t.Run("ArgumentListOpener", func(t *testing.T) {
		testFragment(t, []rune("("), FragmentTypeOpenerArgumentList)
	})

	t.Run("ArgumentListCloser", func(t *testing.T) {
		testFragment(t, []rune(")"), FragmentTypeCloserArgumentList)
	})

	t.Run("ListDelimiter", func(t *testing.T) {
		testFragment(t, []rune(","), FragmentTypeDelimiterList)
	})

	t.Run("TypeSpecifierOptional", func(t *testing.T) {
		testFragment(t, []rune("?"), FragmentTypeSpecifierOptional)
	})

	t.Run("TypeSpecifierList", func(t *testing.T) {
		testFragment(t, []rune("[]"), FragmentTypeSpecifierList)
	})

	t.Run("CommentLine", func(t *testing.T) {
		testFragment(t, []rune("//"), FragmentTypeCommentLine)
	})

	t.Run("Documentation", func(t *testing.T) {
		testFragment(t, []rune("#"), FragmentTypeDocumentationLine)
	})
}
