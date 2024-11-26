package recurrence_test

import (
	"dsa/advanced/recurrence"
	"strconv"
	"testing"
)

var SumTests = []struct {
	n, expected int
}{
	{4, 10},
	{1, 1},
	{2, 3},
	{3, 6},
	{5, 15},
}

func TestSum(t *testing.T) {
	for i, e := range SumTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := recurrence.Sum(e.n)
			if actual != e.expected {
				t.Fatalf("Sum(%v) = %v, want %v", e.n, actual, e.expected)
			}
		})
	}
}
