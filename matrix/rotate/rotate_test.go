package rotate_test

import (
	"dsa/matrix/rotate"
	"fmt"
	"slices"
	"strconv"
	"testing"
)

var RotateTests = []struct {
	matrix   [][]int
	expected [][]int
}{
	{
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
	},
	{
		[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
		[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
	},
}

func TestRotate(t *testing.T) {
	for i, e := range RotateTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cloned := CloneMatrix(e.matrix)
			rotate.Rotate90Inplace(e.matrix)
			PrintMatrix(e.matrix)
			for i := 0; i < len(e.matrix); i++ {
				if !slices.Equal(e.matrix[i], e.expected[i]) {
					t.Fatalf("Rotate(%v) = %v, want %v", cloned, e.matrix, e.expected)
				}
			}
		})
	}
}

func CloneMatrix(matrix [][]int) [][]int {
	ROWS := len(matrix)
	clone := make([][]int, ROWS)
	for i := 0; i < len(clone); i++ {
		clone[i] = slices.Clone(matrix[i])
	}
	return clone
}

func PrintMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}

var ClockwiseTests = []struct {
	row, col, n int
}{
	{0, 0, 3},
}

func TestClockwise(t *testing.T) {
	for i, e := range ClockwiseTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			row, col := e.row, e.col
			for i := 0; i < 4; i++ {
				row, col = rotate.Clockwise90(row, col, e.n)
				fmt.Printf("(%v, %v)\n", row, col)
			}
		})
	}
}

var CounterClockwise90Tests = []struct {
	row, col, n int
}{
	{0, 0, 3},
}

func TestCounterClockwise90(t *testing.T) {
	for i, e := range CounterClockwise90Tests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			row, col := e.row, e.col
			for i := 0; i < 4; i++ {
				row, col = rotate.CounterClockwise90(row, col, e.n)
				fmt.Printf("(%v, %v)\n", row, col)
			}
		})
	}
}
