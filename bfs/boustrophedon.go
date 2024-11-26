package bfs

// walk a grid as if game board in boustrophedon style
func BoustrophedonWalk(board [][]int) []int {
	ROWS, COLS := len(board), len(board[0])
	result := make([]int, 0, ROWS*COLS)
	row, col := ROWS-1, 0
	direction := 1
	for row >= 0 {
		for i := 0; i < COLS; i++ {
			result = append(result, board[row][col])
			col += direction
		}
		// decrement row
		row--
		// flip column position
		direction = -direction
		// bring back into bounds
		col += direction
	}
	return result
}

/*

[]
[1,2,3,4,5]
*/
