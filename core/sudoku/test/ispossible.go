package main

import (
	"dsa/core/sudoku"
	"fmt"
)

var (
	board1 = [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
)

func main() {
	printTestCases(board1, "board1")

}

func printTestCases(board [9][9]int, name string) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for n := 1; n <= 9; n++ {
				fmt.Printf("{%v, %d, %d, %d, %v},\n", name, i, j, n, sudoku.IsPossible(board, i, j, n))
			}
		}
	}

}
