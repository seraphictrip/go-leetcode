package slidingwindow_test

import (
	"dsa/random/slidingwindow"
	"strconv"
	"testing"
)

var MaxVowelsTests = []struct {
	s        string
	k        int
	expected int
}{
	{"abciiidef", 3, 3},
	{"aeiou", 2, 2},
	{"leetcode", 3, 2},
}

func TestMaxVowels(t *testing.T) {
	for i, e := range MaxVowelsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := slidingwindow.MaxVowels(e.s, e.k)
			if actual != e.expected {
				t.Fatalf("MaxVowels(%v, %v) = %v, want %v", e.s, e.k, actual, e.expected)
			}
		})
	}
}
