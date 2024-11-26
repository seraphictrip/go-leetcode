package graph

import "fmt"

func NearestExit(maze [][]byte, entrance []int) int {
	ROWS, COLS := len(maze), len(maze[0])
	queue := make([][]int, 0, len(maze))
	// enqueue
	queue = append(queue, entrance)
	maze[entrance[0]][entrance[1]] = '+'
	// isEmpty
	level := 0
	for len(queue) != 0 {
		n := len(queue)
		fmt.Println(n)
		// process all nodes for level
		for n > 0 {
			// dequeue
			node := queue[0]
			queue = queue[1:]
			row, col := node[0], node[1]
			// process
			if isExit(entrance, row, col, ROWS, COLS) {
				return level
			}
			// queue children
			for _, coor := range neighborss(row, col) {
				if inboundss(coor, ROWS, COLS) && maze[coor[0]][coor[1]] == '.' {
					maze[coor[0]][coor[1]] = '+'
					queue = append(queue, coor)
				}
			}

			n--
		}
		level++
	}

	return -1
}

type coordinate []int

func neighborss(row, col int) [][]int {
	return [][]int{{row + 1, col}, {row - 1, col}, {row, col + 1}, {row, col - 1}}
}

func inboundss(coor []int, ROWS, COLS int) bool {
	row, col := coor[0], coor[1]
	return row >= 0 && row < ROWS && col >= 0 && col < COLS
}

func isExit(entrance []int, row, col, ROWS, COLS int) bool {
	// not entrance
	if row == entrance[0] && col == entrance[1] {
		return false
	}
	if row-1 == -1 || row+1 == ROWS || col-1 == -1 || col+1 == COLS {
		return true
	}
	return false
}
