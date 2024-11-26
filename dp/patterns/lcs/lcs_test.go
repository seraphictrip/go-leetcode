package lcs_test

import (
	"dsa/dp/patterns/lcs"
	"strconv"
	"testing"
)

var LCSBruteForceTests = []struct {
	s1, s2   string
	expected int
}{
	{"ADCB", "ABC", 2},
}

func TestLCSBruteForce(t *testing.T) {
	for i, e := range LCSBruteForceTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := lcs.LCSBottomUp(e.s1, e.s2)
			if actual != e.expected {
				t.Fatalf("LCSBruteForce(%v, %v) = %v, want %v", e.s1, e.s2, actual, e.expected)
			}
		})
	}
}
