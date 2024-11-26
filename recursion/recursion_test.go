package recursion_test

import (
	"dsa/recursion"
	"strconv"
	"testing"
)

var FactorialTests = []struct {
	n        int
	expected uint64
}{
	{-1, 1},
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 6},
	{5, 120},
	{6, 720},
	{7, 7 * 720},
}

func TestFactorial(t *testing.T) {
	for i, e := range FactorialTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := recursion.FactorialIterative(e.n)
			if actual != e.expected {
				t.Fatalf("Factorial(%d) = %d, want %d", e.n, actual, e.expected)
			}
		})
	}
}
