package matrix

import "container/list"

/*
Count unique paths from the top left to the bottom right.  A single path may only move along
0's and can't visit the same cell more than once.
*/

func CountPaths(grid [][]int) int {
	return dfs(grid, 0, 0, map[[2]int]bool{})
}

// Count paths (backtracking)
func dfs(grid [][]int, r, c int, visit map[[2]int]bool) int {
	ROWS, COLS := len(grid), len(grid[0])
	// ran off edge, previously visited, blocked cell
	if r < 0 || c < 0 {
		// ran off top or left
		return 0
	}
	if r == ROWS || c == COLS {
		// ran off bottom, or right
		return 0
	}
	if _, seen := visit[[2]int{r, c}]; seen {
		// already visited this node
		return 0
	}
	if grid[r][c] == 1 {
		// impassable terrain
		return 0
	}

	// reach destination, lower right corner
	if r == ROWS-1 && c == COLS-1 {
		return 1
	}

	// add to visited list
	visit[[2]int{r, c}] = true
	// search all paths
	count := 0
	count += dfs(grid, r+1, c, visit)
	count += dfs(grid, r-1, c, visit)
	count += dfs(grid, r, c+1, visit)
	count += dfs(grid, r, c-1, visit)

	// backtrack
	delete(visit, [2]int{r, c})
	return count
}

type coord [2]int

func Coord(x, y int) coord {
	return [2]int{x, y}
}

func bfs(grid [][]int, visitor func(coor coord)) {
	// get coordinates
	ROWS, COLS := len(grid), len(grid[0])
	visited := map[coord]bool{}
	// queue "root"
	q := NewQueue[coord]()
	q.Enqueue([2]int{0, 0})

	for !q.IsEmpty() {
		n := q.Len()
		for n > 0 {
			coor := q.Dequeue()
			visitor(coor)
			visited[coor] = true
			r, c := coor[0], coor[1]
			isValid := func(cs coord) bool {
				r, c := cs[0], cs[1]
				if r < 0 || c < 0 {
					return false
				}
				if r >= ROWS || c >= COLS {
					return false
				}
				if grid[r][c] == 1 || visited[coor] {
					return false
				}
				return true
			}
			right := Coord(r+1, c)
			down := Coord(0, c+1)
			left := Coord(r-1, c)
			up := Coord(r, c-1)
			if isValid(right) {
				q.Enqueue(right)
			}
			if isValid(left) {
				q.Enqueue(left)
			}
			if isValid(up) {
				q.Enqueue(up)
			}
			if isValid(down) {
				q.Enqueue(down)
			}

		}
	}
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
	val := q.l.Remove(front).(T)
	return val
}

func (q Queue[T]) IsEmpty() bool {
	return q.l.Len() == 0
}

func (q Queue[T]) Len() int {
	return q.l.Len()
}
