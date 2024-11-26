package sudoku

import (
	"fmt"
	"strings"
)

const (
	GRID_SIZE = 9
)

func SolveSudoku(board [][]byte) {
	nums := []byte("123456789")
	isPossible := func(row, col int, num byte) bool {
		if board[row][col] != '.' {
			return false
		}
		// check row
		for _, val := range board[row] {
			if val == num {
				return false
			}
		}
		// check column
		for i := 0; i < 9; i++ {
			if board[i][col] == num {
				return false
			}
		}
		// check square
		row0, col0 := (row/3)*3, (col/3)*3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[row0+i][col0+j] == num {
					return false
				}
			}
		}
		return true
	}
	var solve func() bool
	solve = func() bool {
		for row := 0; row < GRID_SIZE; row++ {
			for col := 0; col < GRID_SIZE; col++ {
				if board[row][col] == '.' {
					// for each spot in grid not yet filled try all possibilities
					for _, num := range nums {
						if isPossible(row, col, num) {
							board[row][col] = num
							if solve() {
								// if it results in a full solution return
								return true
							}
							// otherwise backtack nad keep trying
							// backtrack
							board[row][col] = '.'
						}
					}
					return false
				}
			}
		}
		PrintGrid(board)
		return true
		// if we filled every square we are done?
	}
	solve()

}

func PrintGrid(board [][]byte) {
	for _, row := range board {
		srow := string(row)
		numrow := strings.Split(srow, "")
		fmt.Println(numrow)
	}
}
