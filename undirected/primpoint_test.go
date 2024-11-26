package undirected_test

import (
	"dsa/undirected"
	"fmt"
	"strconv"
	"testing"
)

var MakeAdjTests = []struct {
	points   [][]int
	expected map[int][][2]int
}{
	{[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, nil},
}

func TestMakeAdj(t *testing.T) {
	for i, e := range MakeAdjTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := undirected.MakeAdjList(e.points)
			fmt.Println(actual)
		})
	}
}

var MinCostConnectPointsTests = []struct {
	points   [][]int
	expected int
}{
	{[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, 20},
}

func TestMinCostConnectPoints(t *testing.T) {
	for i, e := range MinCostConnectPointsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := undirected.MinCostConnectPoints(e.points)
			if actual != e.expected {
				t.Fatalf("MinCostConnectPoints(%v) = %v, want %v", e.points, actual, e.expected)
			}
		})
	}
}
