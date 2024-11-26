package adjacencylist_test

import (
	al "dsa/adjacencylist"
	"fmt"
	"slices"
	"strconv"
	"testing"
)

var MakeAdjacencyListTests = []struct {
	edges [][]string
}{
	{[][]string{{"A", "B"}, {"B", "C"}, {"B", "E"}, {"C", "E"}, {"E", "D"}, {"Z", "D"}}},
}

func TestMakeAdjacencyList(t *testing.T) {
	for i, e := range MakeAdjacencyListTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := al.MakeAdjacencyList(e.edges)
			fmt.Println(actual)
			// fmt.Println(al.CountPaths("A", "E", actual, map[string]bool{}))
			fmt.Println(al.CountPathsBFS(actual, "A", "E"))
			// fmt.Println(al.GetPaths(actual, "A", "Z"))
			// fmt.Println(al.GetPaths(actual, "A", "E"))
		})
	}
}

var GetPathsTests = []struct {
	edges       [][]int
	src, target int
	expected    [][]int
}{
	{
		[][]int{{0, 1}, {1, 2}, {1, 4}, {2, 4}, {4, 3}}, 0, 3,
		[][]int{{0, 1, 2, 4, 3}, {0, 1, 4, 3}},
	},
}

func TestGetPaths(t *testing.T) {
	for i, e := range GetPathsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			graph := al.MakeAdjacencyList(e.edges)
			actual := al.GetPaths(graph, e.src, e.target)
			if len(actual) != len(e.expected) {
				t.Fatalf("GetPaths(%v, %v, %v) = %v, want %v", graph, e.src, e.target, actual, e.expected)
			}
			for i := range actual {
				if !slices.Equal(actual[i], e.expected[i]) {
					t.Fatalf("GetPaths(%v, %v, %v) = %v, want %v", graph, e.src, e.target, actual, e.expected)
				}
			}
		})
	}
}

var ShortestPathTests = []struct {
	edges       [][]string
	src, target string
	expected    int
}{
	{
		[][]string{{"A", "B"}, {"B", "C"}, {"B", "E"}, {"C", "E"}, {"E", "D"}, {"Z", "D"}},
		"A", "E", 2,
	},
	{
		[][]string{{"A", "B"}, {"B", "C"}, {"B", "E"}, {"C", "E"}, {"E", "D"}, {"Z", "D"}},
		"A", "Z", -1,
	},
}

func TestShortestPath(t *testing.T) {
	for i, e := range ShortestPathTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			graph := al.MakeAdjacencyList(e.edges)
			actual := al.ShortestPath(graph, e.src, e.target)
			if actual != e.expected {
				t.Fatalf("ShortestPath(%v, %v, %v) = %v, want %v", graph, e.src, e.target, actual, e.expected)
			}
		})
	}
}
