package lcs_test

import (
	"dsa/dp/patterns/lcs"
	"strconv"
	"testing"
)

var isInterleaveTests = []struct {
	s1, s2, s3 string
	expected   bool
}{
	{"", "", "", true},
	{"aabcc", "dbbca", "aadbbcbcac", true},
	{"", "", "1", false},
	{"aabcc", "dbbca", "aadbbbaccc", false},
}

func TestIsInterleave(t *testing.T) {
	for i, e := range isInterleaveTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := lcs.IsInterleaveBottomUp(e.s1, e.s2, e.s3)
			if actual != e.expected {
				t.Fatalf("IsInterleave(%v, %v, %v) = %v, want %v", e.s1, e.s2, e.s3, actual, e.expected)
			}
		})
	}
}
