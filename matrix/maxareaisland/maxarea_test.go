package maxareaisland_test

import (
	"dsa/matrix/maxareaisland"
	"strconv"
	"testing"
)

var MaxAreaOfIslandTests = []struct {
	grid     [][]int
	expected int
}{
	{
		[][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		},
		6,
	},
}

func TestMaxAreaOfIsland(t *testing.T) {
	for i, e := range MaxAreaOfIslandTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := maxareaisland.MaxAreaOfIsland(e.grid)
			if actual != e.expected {
				t.Fatalf("MaxAreaOfIsland(..) = %v, want %v", actual, e.expected)
			}
		})
	}
}
