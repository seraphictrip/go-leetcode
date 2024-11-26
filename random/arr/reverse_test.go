package arr_test

import (
	"dsa/random/arr"
	"strconv"
	"testing"
)

var ReverseVowelsTests = []struct {
	s, expected string
}{
	{
		"IceCreAm", "AceCreIm",
	},
}

func TestReverseVowels(t *testing.T) {
	for i, e := range ReverseVowelsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := arr.ReverseVowels(e.s)
			if actual != e.expected {
				t.Fatalf("ReverseVowels(%v) = %v, want %v", e.s, actual, e.expected)
			}
		})
	}
}

var ReverseWordsTests = []struct {
	s, expected string
}{
	{
		"a good   example",
		"example good a",
	},
}

func TestReverseWords(t *testing.T) {
	for i, e := range ReverseWordsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := arr.ReverseWords(e.s)
			if actual != e.expected {
				t.Fatalf("ReverseWords(%v) = %v, want %v", e.s, actual, e.expected)
			}
		})
	}
}
