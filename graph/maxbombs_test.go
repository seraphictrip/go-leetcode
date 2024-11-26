package graph_test

import (
	"dsa/graph"
	"strconv"
	"testing"
)

var MaximumDetonationTests = []struct {
	bombs    [][]int
	expected int
}{
	{
		[][]int{{2, 1, 3}, {6, 1, 4}}, 2,
	},
	{
		[][]int{{1, 1, 5}, {10, 10, 5}}, 1,
	},
	{
		[][]int{
			{1, 2, 3},
			{2, 3, 1},
			{3, 4, 2},
			{4, 5, 3},
			{5, 6, 4},
		}, 5,
	},
}

func TestMaximumDetonation(t *testing.T) {
	for i, e := range MaximumDetonationTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := graph.MaximumDetonation(e.bombs)
			if actual != e.expected {
				t.Fatalf("MaximumDetonation(%v) = %v, want %v", e.bombs, actual, e.expected)
			}
		})
	}
}
