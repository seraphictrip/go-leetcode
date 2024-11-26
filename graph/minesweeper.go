package graph

import "strconv"

func UpdateBoard(board [][]byte, click []int) [][]byte {
	ROWS, COLS := len(board), len(board[0])
	row, col := click[0], click[1]

	// out of bounds
	if (row < 0 || row >= ROWS) || (col < 0 || col >= COLS) {
		return board
	}

	switch board[row][col] {
	case 'M':
		board[row][col] = 'X'
		return board
	case 'E':
		close := countNeighboringBombs(board, row, col)
		if close == 0 {
			// visit all neighbors recursively
			neighs := neighbors(row, col)
			board[row][col] = 'B'
			for _, coor := range neighs {
				if inbounds(board, coor[0], coor[1]) && board[coor[0]][coor[1]] == 'E' {
					board = UpdateBoard(board, coor)
				}
			}
		} else {
			board[row][col] = []byte(strconv.Itoa(close))[0]
		}

	}

	return board
}

func inbounds(board [][]byte, row, col int) bool {
	ROWS, COLS := len(board), len(board[0])
	if (row < 0 || row >= ROWS) || (col < 0 || col >= COLS) {
		return false
	}
	return true
}

func countNeighboringBombs(board [][]byte, row, col int) int {
	n := neighbors(row, col)
	result := 0
	for _, coor := range n {
		if inbounds(board, coor[0], coor[1]) {
			if board[coor[0]][coor[1]] == 'M' {
				result++
			}
		}
	}
	return result
}

func neighbors(row, col int) [][]int {
	up := []int{row - 1, col}
	upright := []int{row - 1, col + 1}
	right := []int{row, col + 1}
	downright := []int{row + 1, col + 1}
	down := []int{row + 1, col}
	downleft := []int{row + 1, col - 1}
	left := []int{row, col - 1}
	upleft := []int{row + 1, col - 1}

	return [][]int{up, upright, right, downright, down, downleft, left, upleft}
}
