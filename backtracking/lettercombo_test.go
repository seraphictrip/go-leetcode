package backtracking_test

import (
	"dsa/backtracking"
	"slices"
	"strconv"
	"testing"
)

var LetterCombinationsTests = []struct {
	digits   string
	expected []string
}{
	{"23", nil},
}

func TestLetterCombinations(t *testing.T) {
	for i, e := range LetterCombinationsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := backtracking.LetterCombinations(e.digits)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("LetterCombinations(%v) = %v, want %v", e.digits, actual, e.expected)
			}
		})
	}
}
