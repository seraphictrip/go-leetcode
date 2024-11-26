package dp_test

import (
	"dsa/dp"
	"strconv"
	"testing"
)

var ClimbingStairsTests = []struct {
	n, expected int
}{
	{1, 1},
	{2, 2},
	{3, 3},
	{45, 1836311903},
}

func TestClimbingStairs(t *testing.T) {
	for i, e := range ClimbingStairsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := dp.ClimbStairs(e.n)
			if actual != e.expected {
				t.Fatalf("ClimbStairs(%v) = %v, want %v", e.n, actual, e.expected)
			}
		})
	}
}
