package shortestpath_test

import (
	"dsa/matrix/shortestpath"
	"strconv"
	"testing"
)

var ShortestPathTests = []struct {
	grid     [][]int
	expected int
}{
	{[][]int{{0, 0, 0, 0}}, 3},
	{[][]int{{0, 0}, {0, 0}}, 2},
	{[][]int{{0, 0}, {1, 0}}, 2},
	{[][]int{{0, 1}, {0, 0}}, 2},
	{[][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}, 6},
}

func TestShortestPath(t *testing.T) {
	for i, e := range ShortestPathTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := shortestpath.ShortestPath(e.grid)

			if actual != e.expected {
				t.Fatalf("ShortestPath(%v) = %v, want %v", e.grid, actual, e.expected)
			}
		})
	}
}

var ShortestPathBinaryMatrixTests = []struct {
	grid     [][]int
	expected int
}{
	// {nil, -1},
	// {[][]int{{0}}, 1},
	{[][]int{{0, 1}, {1, 0}}, 2},
	{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 4},
	{[][]int{
		{0, 1, 0, 1},
		{1, 0, 0, 1},
		{0, 1, 1, 1},
		{1, 0, 0, 0},
	}, 6},
	{[][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}, -1},
}

func TestShortestPathBinaryMatrix(t *testing.T) {
	for i, e := range ShortestPathBinaryMatrixTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := shortestpath.ShortestPathBinaryMatrix(e.grid)
			if actual != e.expected {
				t.Fatalf("ShortestPathBinaryMatrix(%v) = %v, want %v", e.grid, actual, e.expected)
			}
		})
	}
}
