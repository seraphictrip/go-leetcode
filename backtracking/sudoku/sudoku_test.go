package sudoku_test

import (
	"dsa/backtracking/sudoku"
	"fmt"
	"strconv"
	"testing"
)

var SolveSudokuTests = []struct {
	board   [][]byte
	expectd [][]byte
}{
	{
		[][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
		nil,
	},
}

func TestSolveSudoku(t *testing.T) {
	for i, e := range SolveSudokuTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			sudoku.SolveSudoku(e.board)
			fmt.Println(e.board)
		})
	}
}
