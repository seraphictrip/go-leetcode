package sudoku

import "fmt"

/*
36. Valid Sudoku
Determine if a 9 x 9 Sudoku board is valid.
Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.


Example 1:


Input: board =
	 c0 c1 ...
r0[	["5","3",".",".","7",".",".",".","."]
r1	,["6",".",".","1","9","5",".",".","."]
r2	,[".","9","8",".",".",".",".","6","."]
r3 ,["8",".",".",".","6",".",".",".","3"]
r4 ,["4",".",".","8",".","3",".",".","1"]
r5 ,["7",".",".",".","2",".",".",".","6"]
r6 ,[".","6",".",".",".",".","2","8","."]
r7 ,[".",".",".","4","1","9",".",".","5"]
r8 ,[".",".",".",".","8",".",".","7","9"]]

sq00 sq01 sq02
sq10 sq11 sq12
sq20 sq21 sq22
Output: true
Example 2:

Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.


Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/

// We need to group
func IsValidSudoku(board [][]byte) bool {
	// create set for seen values
	seen := make(map[string]bool)
	// iterate over all cells
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// check if already seen in row, col, sq
			// here we encode row as r0_Val, so for example
			// if we see 9 in r0 twice it would fail with key r0_9
			rowKey := fmt.Sprintf("r%v_%v", i, board[i][j])
			// same thing with column, an alt solution would have a set per row/column
			colKey := fmt.Sprintf("c%v_%v", j, board[i][j])
			// same thing wih squares, use int math to get squares
			// 00, 01, 02, 10, 11, 12 20, 21, 22
			sqKey := fmt.Sprintf("sq%v_%v_%v", i/3, j/3, board[i][j])
			if seen[rowKey] || seen[colKey] || seen[sqKey] {
				return false
			}
			seen[rowKey] = true
			seen[colKey] = true
			seen[sqKey] = true
		}
	}
	return true
}

type Board [9][9]int

// Check if it is possible (given current state of board)
// to play n at board[r][c]
func IsPossible(board Board, r, c, n int) bool {
	// already filled
	if board[r][c] != 0 {
		return false
	}
	// ensure n is not in row
	for i := 0; i < 9; i++ {
		if board[r][i] == n {
			return false
		}
	}
	// ensure n is not in col
	for i := 0; i < 9; i++ {
		if board[i][c] == n {
			return false
		}
	}

	// ensure n is not in sq
	r0 := r / 3
	c0 := c / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[r0+i][c0+j] == n {
				return false
			}
		}
	}

	return true

}

func Solve(board Board) Board {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == 0 {
				// blank, we can try
				for n := 1; n <= 9; n++ {
					if IsPossible(board, r, c, n) {
						board[r][c] = n
						board = Solve(board)
					}
					// backtrack
					board[r][c] = 0
				}
			}
		}
	}
	return board
}
