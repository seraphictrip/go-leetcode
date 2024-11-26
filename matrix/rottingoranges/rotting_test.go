package rottingoranges_test

import (
	ro "dsa/matrix/rottingoranges"
	"fmt"
	"strconv"
	"testing"
)

var NewNodeTests = []struct {
	val any
}{
	{1},
	{"a"},
}

func TestNewNode(t *testing.T) {
	for i, e := range NewNodeTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			node := ro.NewNode(e.val)
			// sets right value, and is generic via different test cases
			if node.Val != e.val {
				t.Fatalf("NewNode(%v).Val != %v", e.val, e.val)
			}
			// does not set neighors
			if node.Up() != nil || node.Right() != nil || node.Down() != nil || node.Left() != nil {
				t.Fatalf("unexpected children %+v", node)
			}
		})
	}
}

var SetNeighborTests = []struct {
	node, neighbor *ro.Node[int]
	direction      ro.Direction
}{
	{ro.NewNode(0), ro.NewNode(1), ro.UP},
	{ro.NewNode(0), ro.NewNode(1), ro.RIGHT},
	{ro.NewNode(0), ro.NewNode(1), ro.DOWN},
	{ro.NewNode(0), ro.NewNode(1), ro.LEFT},
}

func TestSetNeighbor(t *testing.T) {
	for i, e := range SetNeighborTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			e.node.SetNeighbor(e.direction, e.neighbor)
			if e.node.Look(e.direction) != e.neighbor {
				t.Fatalf("SetNeighbor(%v, %v) does not work", e.direction, e.neighbor)
			}
		})
	}
}

var SetNeighborsTests = []struct {
	node, up, right, down, left *ro.Node[int]
}{
	{ro.NewNode(0), ro.NewNode(1), ro.NewNode(2), ro.NewNode(3), ro.NewNode(4)},
}

func TestSetNeighbors(t *testing.T) {
	for i, e := range SetNeighborsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			neighbors := [4]*ro.Node[int]{e.up, e.right, e.down, e.left}
			e.node.SetNeighbors(neighbors)

			if e.node.Neighbors() != neighbors {
				t.Fatalf("Neighbors() = %v, want %v", e.node.Neighbors(), neighbors)
			}

		})
	}
}

var BuildGraphFromMatrixTests = []struct {
	matrix [][]int
}{
	{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}},
}

func TestBuildGraphFromMatrix(t *testing.T) {
	for i, e := range BuildGraphFromMatrixTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			graph := ro.BuildGraphFromMatrix(e.matrix)
			fmt.Println(graph)
		})
	}
}

var OranageRottingTests = []struct {
	grid     [][]int
	expected int
}{
	{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
	{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
	{[][]int{{0, 2}}, 0},
}

func TestOranageRotting(t *testing.T) {
	for i, e := range OranageRottingTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := ro.OrangesRotting(e.grid)
			if actual != e.expected {
				t.Fatalf("OrangesRotting(%v) = %v, want %v", e.grid, actual, e.expected)
			}
		})
	}
}
