package arr_test

import (
	"dsa/random/arr"
	"strconv"
	"testing"
)

var MergeTests = []struct {
	word1, word2, expected string
}{
	{"abc", "prq", "apbrcq"},
}

func TestMerge(t *testing.T) {
	for i, e := range MergeTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := arr.MergeAlternately(e.word1, e.word2)
			if actual != e.expected {
				t.Fatalf("MergeAlternately(%v, %v) = %v, want %v", e.word1, e.word2, actual, e.expected)
			}
		})
	}
}
