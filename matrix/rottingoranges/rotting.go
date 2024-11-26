package rottingoranges

import (
	"container/list"
	"errors"
	"fmt"
)

/*
You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange
 becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.
If this is impossible, return -1.
			(2)
		(1)		(1)
	(0)	 (1)
Input: grid = [[2,1,1],
			   [1,1,0],
			   [0,1,1]]
Output: 4
Example 2:

Input: grid = [[2,1,1],
			   [0,1,1],
			   [1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
Example 3:

Input: grid = [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.


Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 10
grid[i][j] is 0, 1, or 2.
*/

type state int

const (
	EMPTY state = iota
	FRESH
	ROTTEN
)

type coordinate [2]int

func OrangesRotting(grid [][]int) int {
	ROWS, COLS := len(grid), len(grid[0])
	q := NewQueue[coordinate]()
	fresh := 0
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if grid[i][j] == 1 {
				fresh++
			}
			if grid[i][j] == 2 {
				// rotten, queue for spread
				q.Enqueue([2]int{i, j})
			}
		}
	}

	if fresh == 0 {
		return 0
	}
	mins := -1
	for !q.IsEmpty() {
		n := q.Len()
		for n > 0 {
			rotten := q.Dequeue()
			r, c := rotten[0], rotten[1]
			// up
			if inbounds(grid, r-1, c) && grid[r-1][c] == 1 {
				grid[r-1][c] = 2
				q.Enqueue([2]int{r - 1, c})
				fresh--
			}
			// right
			if inbounds(grid, r, c+1) && grid[r][c+1] == 1 {
				grid[r][c+1] = 2
				q.Enqueue([2]int{r, c + 1})
				fresh--
			}
			// down
			if inbounds(grid, r+1, c) && grid[r+1][c] == 1 {
				grid[r+1][c] = 2
				q.Enqueue([2]int{r + 1, c})
				fresh--
			}
			// left
			if inbounds(grid, r, c-1) && grid[r][c-1] == 1 {
				grid[r][c-1] = 2
				q.Enqueue([2]int{r, c - 1})
				fresh--
			}

			n--
		}
		mins++
	}

	if fresh == 0 {
		return mins
	}

	return -1
}

type QueueInterface[T any] interface {
	Enqueue(T)
	Dequeue() T
	Len() int
	IsEmpty() bool
}

type Queue[T any] struct {
	l *list.List
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		l: list.New(),
	}
}

func (q Queue[T]) Enqueue(item T) {
	q.l.PushBack(item)
}

func (q Queue[T]) Dequeue() T {
	front := q.l.Front()
	result := q.l.Remove(front).(T)
	return result
}

func (q Queue[T]) Len() int {
	return q.l.Len()
}

func (q Queue[T]) IsEmpty() bool {
	return q.l.Len() == 0
}

/*
Input: grid = [[2,1,1],
			   [1,1,0],
			   [0,1,1]]

*/

/*

	root := NewNode(2)
	coor0_1 := NewNode(1)
	coor1_0 := NewNode(1)
	root.SetNeighbor(RIGHT, coor0_1)
	coor0_1.SetNeighbor(LEFT, root)
	root.SetNeighbor(DOWN, coor1_0)
	coor1_0.SetNeighor(UP, root)

	coor0_2 := NewNode(1)

	r1.SetNeighbor(RIGHT, r1_r1)
	r1.SetNeighbor(DOWN, r1_d1)



*/

type Direction uint8

const (
	UP Direction = iota + 1
	// LEFT_UP
	LEFT
	// LEFT_DOWN
	DOWN
	// RIGHT_DOWN
	RIGHT
	// RIGHT_UP
)

var (
	ErrUnknownDirection = errors.New("unknown direction")
	ErrNoNeighbor       = errors.New("no neighbor")
)

// A node represents a square in a matrix
// neighors are up, down, left, right
// nil neighbors represent out of bounds
// nil values should be stored in node if valid nodes
type Node[T any] struct {
	Val   T
	up    *Node[T]
	down  *Node[T]
	left  *Node[T]
	right *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{
		Val: val,
	}
}

// SetNeighbor simply sets MY neighbor
// They are responsible for setting me as a neighor if required
// see BuildGraph  for convienence methods
func (n *Node[T]) SetNeighbor(d Direction, node *Node[T]) {
	switch d {
	case UP:
		n.up = node
	case RIGHT:
		n.right = node
	case DOWN:
		n.down = node
	case LEFT:
		n.left = node
	default:
		fmt.Println("currently unsupported")
	}
}

// Set neighbors in [UP, RIGHT, DOWN, LEFT] (clockwise/NESW)
// use nil to indicate matrix boundaries
// ex: (0,0) on 2x2: root.SetNeighbors([nil, n0_1, n1_0, nil])
// ex: (1,1) on 3x3+: root.SetNeighbors([n0_1, n0_2, n2_1, n1_0])
// ex: middle: root.SetNeighbors([up, right, down, left])
func (n *Node[T]) SetNeighbors(neighbors [4]*Node[T]) {
	up := neighbors[0]
	right := neighbors[1]
	down := neighbors[2]
	left := neighbors[3]
	n.SetNeighbor(UP, up)
	n.SetNeighbor(RIGHT, right)
	n.SetNeighbor(DOWN, down)
	n.SetNeighbor(LEFT, left)
}

// Get Neighbors in [UP, RIGHT, DOWN, LEFT]
// nil indicates matrix boundary
func (n *Node[T]) Neighbors() [4]*Node[T] {
	return [4]*Node[T]{n.up, n.right, n.down, n.left}
}

// Look in a direction
func (n *Node[T]) Look(d Direction) *Node[T] {
	switch d {
	case UP:
		return n.up
	case RIGHT:
		return n.right
	case DOWN:
		return n.down
	case LEFT:
		return n.left
	default:
		fmt.Println("currently unsupported")
		return nil
	}
}

func (n *Node[T]) LookValue(d Direction) (val T, err error) {
	var node *Node[T]
	switch d {
	case UP:
		node = n.up
	case RIGHT:
		node = n.right
	case DOWN:
		node = n.down
	case LEFT:
		node = n.left
	default:
		return val, errors.ErrUnsupported
	}
	if node == nil {
		return val, ErrNoNeighbor
	}
	return node.Val, nil
}

// Get node in UP position
func (n *Node[T]) Up() *Node[T] {
	return n.up
}

func (n *Node[T]) Right() *Node[T] {
	return n.right
}

func (n *Node[T]) Down() *Node[T] {
	return n.down
}

func (n *Node[T]) Left() *Node[T] {
	return n.left
}

func BuildGraphFromMatrix[T any](matrix Matrix[T]) *Node[T] {
	ROWS := len(matrix)
	if ROWS == 0 {
		return nil
	}
	COLUMNS := len(matrix[0])
	// create empty graph
	graph := make([][]*Node[T], ROWS)
	for i := range graph {
		graph[i] = make([]*Node[T], COLUMNS)
	}
	// fill graph with nodes
	// O(n*m)
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			graph[i][j] = NewNode(matrix[i][j])
		}
	}

	// connect nodes
	// O(n*m)
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLUMNS; col++ {
			node := graph[row][col]

			if inbounds(graph, row-1, col) {
				up := graph[row-1][col]
				node.SetNeighbor(UP, up)
			}
			if inbounds(graph, row, col+1) {
				right := graph[row][col+1]
				node.SetNeighbor(RIGHT, right)
			}
			if inbounds(graph, row+1, col) {
				down := graph[row+1][col]
				node.SetNeighbor(DOWN, down)
			}
			if inbounds(graph, row, col-1) {
				left := graph[row][col-1]
				node.SetNeighbor(LEFT, left)
			}
		}
	}

	return graph[0][0]
}

// check if coordinates are inbound for a specific grid
// an index i,j is in bounds if it is a node within grid size n x m
func inbounds[T any](grid Matrix[T], i, j int) bool {
	if len(grid) == 0 {
		return false
	}
	rows, cols := len(grid), len(grid[0])
	if i < 0 || j < 0 {
		return false
	}
	if i >= rows || j >= cols {
		return false
	}

	return true
}

// Get the dimensions of a matrix nxm
func Dimensions[T any](matrix Matrix[T]) (n int, m int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

type Matrix[T any] [][]T

func BuildTree[T any]([][]T) *Node[T] {
	panic("not implemented")
}
