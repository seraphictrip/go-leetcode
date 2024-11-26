package lcs_test

import (
	"dsa/dp/patterns/lcs"
	"strconv"
	"testing"
)

var NumDistinctTests = []struct {
	s, t     string
	expected int
}{
	{
		"rabbbit", "rabbit", 3,
	},
	{
		"bbaba", "ba", 5,
	},
}

func TestNumDistinct(t *testing.T) {
	for i, e := range NumDistinctTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := lcs.NumDistinctBottomUp(e.s, e.t)
			if actual != e.expected {
				t.Fatalf("NumDistinct(%v, %v) = %v, want %v", e.s, e.t, actual, e.expected)
			}
		})
	}
}
