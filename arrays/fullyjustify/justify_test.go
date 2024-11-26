package fullyjustify_test

import (
	"dsa/arrays/fullyjustify"
	"slices"
	"strconv"
	"testing"
)

var FullJustifyTests = []struct {
	words    []string
	maxWidth int
	expected []string
}{
	{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16, []string{"This    is    an",
		"example  of text",
		"justification.  "}},
	{
		[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
		16,
		[]string{"What   must   be",
			"acknowledgment  ",
			"shall be        "},
	},
	{
		[]string{"a", "b", "c", "d", "e"}, 3, []string{"a b", "c d", "e  "},
	},
}

func TestFullJustify(t *testing.T) {
	for i, e := range FullJustifyTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := fullyjustify.FullJustify(e.words, e.maxWidth)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("FullJustify(%v, %v) = %v, want %v", e.words, e.maxWidth, actual, e.expected)
			}
		})
	}
}

var justifyBetweenTests = []struct {
	input    []string
	maxWidth int
	expected string
}{
	{
		[]string{"This", "is", "an"}, 16, "This    is    an",
	},
}

func TestJustifyBetween(t *testing.T) {
	for i, e := range justifyBetweenTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := fullyjustify.JustifyBetween(e.input, e.maxWidth)
			if actual != e.expected {
				t.Fatalf("JustifyBetween(%v, %v) = %v, want %v", e.input, e.maxWidth, actual, e.expected)
			}
		})
	}
}
