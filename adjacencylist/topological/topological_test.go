package topological_test

import (
	"dsa/adjacencylist/topological"
	"errors"
	"slices"
	"strconv"
	"testing"
)

var TopologicalSortTests = []struct {
	graph    [][]int
	V        int
	expected []int
	err      error
}{
	{
		[][]int{{}},
		1,
		[]int{0},
		nil,
	},
	{
		[][]int{{0}},
		1,
		[]int{},
		topological.ErrCycle,
	},
	{
		// (0)->(1)
		[][]int{{1}, {}},
		2,
		[]int{0, 1},
		nil,
	},
	{
		//		(0)<->(1)
		[][]int{{1}, {0}}, 2, nil, topological.ErrCycle,
	},
	{
		//		(0)->(1)->(2)
		[][]int{{1}, {2}, {}}, 3, []int{0, 1, 2}, nil,
	},
	{
		//		(0)<->(1)->(2)
		[][]int{{1}, {0, 2}, {}}, 3, []int{}, topological.ErrCycle,
	},
	{
		//			(0)->(1)->(2)->(3)
		[][]int{{1}, {2}, {3}, {}}, 4, []int{0, 1, 2, 3}, nil,
	},
	{
		//  				(0)
		//					/
		// 			(2)-(1)
		//			|	\|
		//			(3)-(4)
		[][]int{{1}, {2, 4}, {3, 4}, {4}, {}}, 5, []int{0, 1, 2, 3, 4}, nil,
	},
	{
		//  				(0)
		//					/
		// 		   c(2)-(1)
		//			|	\|
		//			(3)-(4)
		[][]int{{1}, {2, 4}, {2, 3, 4}, {4}, {}}, 5, []int{}, topological.ErrCycle,
	},
}

func TestTopologicalSort(t *testing.T) {
	for i, e := range TopologicalSortTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual, err := topological.Sort(e.graph, e.V)
			if !errors.Is(err, e.err) {
				t.Fatalf("unexpected error: %v", err)
			}
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("Sort(%v, %v) = %v, want %v", e.graph, e.V, actual, e.expected)
			}
			m := toMap(e.graph)
			actual, err = topological.TopologicalSort(m, e.V)
			if !errors.Is(err, e.err) {
				t.Fatalf("unexpected error: %v", err)
			}
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("SortGeneric(%v, %v) = %v, want %v", e.graph, e.V, actual, e.expected)
			}
		})
	}
}

func toMap(adj [][]int) map[int][]int {
	m := make(map[int][]int, len(adj))
	for i := range adj {
		m[i] = adj[i]
	}
	return m
}
