package arr_test

import (
	"dsa/random/arr"
	"strconv"
	"testing"
)

var GCDTests = []struct {
	a, b, expected int
}{
	{1, 66, 1},
	{2, 66, 2},
}

func TestGCD(t *testing.T) {
	for i, e := range GCDTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := arr.GCD(e.a, e.b)
			if actual != e.expected {
				t.Fatalf("GCD(%v, %v) = %v, want %v", e.a, e.b, actual, e.expected)
			}
		})
	}
}

var GcdOfStringsTests = []struct {
	word1, word2, expected string
}{
	{
		"ABCABC", "ABC", "ABC",
	},
	{
		"ABABAB", "ABAB", "AB",
	},
}

func TestGcdOfStrings(t *testing.T) {
	for i, e := range GcdOfStringsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := arr.GcdOfStrings(e.word1, e.word2)
			if actual != e.expected {
				t.Fatalf("GcdOfStrings(%q, %q) = %q, want %q", e.word1, e.word2, actual, e.expected)
			}
		})
	}
}
