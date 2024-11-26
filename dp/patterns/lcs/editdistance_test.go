package lcs_test

import (
	"dsa/dp/patterns/lcs"
	"strconv"
	"testing"
)

var MinDistanceTests = []struct {
	word1, word2 string
	expected     int
}{
	{"horse", "ros", 3},
	{"intention", "execution", 5},
}

func TestMinDistance(t *testing.T) {
	for i, e := range MinDistanceTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := lcs.MinDistanceBottomUp(e.word1, e.word2)
			if actual != e.expected {
				t.Fatalf("MinDistance(%v, %v) = %v, want %v", e.word1, e.word2, actual, e.expected)
			}
		})
	}
}
