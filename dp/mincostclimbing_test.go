package dp_test

import (
	"dsa/dp"
	"strconv"
	"testing"
)

var MinCostClimbingStairsTests = []struct {
	costs    []int
	expected int
}{
	// min(10)
	{[]int{10, 15, 20}, 15},
	{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
}

func TestMinCostClimbingStairs(t *testing.T) {
	for i, e := range MinCostClimbingStairsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := dp.MinCostClimbingStairsTopDown(e.costs)
			if actual != e.expected {
				t.Fatalf("MinCostClimbingStairs(%v) = %v, want %v", e.costs, actual, e.expected)
			}
		})
	}
}
