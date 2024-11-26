package shortestpath

import "container/list"

/*
Find the length of the shortest path from the top left of the grid to the bottom right
*/

type coordinate [2]int

func ShortestPath(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	// BFS to find paths and take the shortest
	ROWS, COLS := len(grid), len(grid[0])
	goal := [2]int{ROWS - 1, COLS - 1}
	visited := map[coordinate]bool{}
	q := NewQueue[coordinate]()
	coor := [2]int{0, 0}
	visited[coor] = true
	q.Enqueue(coor)

	length := 0
	for !q.IsEmpty() {
		// process everyhing at this level
		n := q.Len()
		for n > 0 {
			// dequeue and process
			coor := q.Dequeue()
			r, c := coor[0], coor[1]
			if coor == goal {
				return length
			}
			enqueueIfValid(grid, q, visited, r+1, c)
			enqueueIfValid(grid, q, visited, r-1, c)
			enqueueIfValid(grid, q, visited, r, c+1)
			enqueueIfValid(grid, q, visited, r, c-1)
			n--
		}
		length += 1
	}
	return -1
}

func enqueueIfValid(grid [][]int, q Queue[coordinate], visited map[coordinate]bool, r, c int) {
	// check bonds
	ROWS, COLS := len(grid), len(grid[0])
	if r < 0 || c < 0 || r >= ROWS || c >= COLS {
		return
	}
	coor := [2]int{r, c}
	// check passability
	if grid[r][c] == 1 || visited[coor] {
		return
	}

	q.Enqueue(coor)
	// setting visited here seems odd to me, I'll run tests
	visited[coor] = true

}

type QueueInterface[T any] interface {
	Enqueue(T)
	Dequeue() T
	Len() int
	IsEmpty() bool
}

// Queue wrapper around
type Queue[T any] struct {
	l *list.List
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		l: list.New(),
	}
}

// FIFO, items enter the rear of the queue
// and exit the front
func (q Queue[T]) Enqueue(item T) {
	q.l.PushBack(item)
}

// FIFO, items enter the rear of the queue
// and exit the front
func (q Queue[T]) Dequeue() T {
	front := q.l.Front()
	val := q.l.Remove(front).(T)
	return val
}

func (q Queue[T]) Len() int {
	return q.l.Len()
}

func (q Queue[T]) IsEmpty() bool {
	return q.l.Len() == 0
}

/*
Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix. If there is no clear path, return -1.

A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right cell (i.e., (n - 1, n - 1)) such that:

All the visited cells of the path are 0.
All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they share an edge or a corner).
The length of a clear path is the number of visited cells of this path.

*/

func ShortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return -1
	}
	length := 0
	// Given an n x n binary matrix grid, return the length
	ROWS, COLS := n, n // nxn so didn't need cols, but might as well practice
	q := NewQueue[coordinate]()
	root := [2]int{0, 0}
	goal := [2]int{ROWS - 1, COLS - 1}
	q.Enqueue(root)
	visited := map[coordinate]bool{}
	visited[root] = true
	for !q.IsEmpty() {
		// check all on level
		n := q.Len()
		for n > 0 {
			cur := q.Dequeue()
			r, c := cur[0], cur[1]
			if cur == goal {
				return length + 1
			}
			enqueueIfValid(grid, q, visited, r+1, c)
			enqueueIfValid(grid, q, visited, r+1, c+1)
			enqueueIfValid(grid, q, visited, r+1, c-1)
			enqueueIfValid(grid, q, visited, r, c+1)
			enqueueIfValid(grid, q, visited, r, c-1)
			enqueueIfValid(grid, q, visited, r-1, c)
			enqueueIfValid(grid, q, visited, r-1, c+1)
			enqueueIfValid(grid, q, visited, r-1, c-1)

			n--
		}
		length++
	}

	// of the shortest clear path in the matrix. If there is no clear path, return -1.
	return -1
}
