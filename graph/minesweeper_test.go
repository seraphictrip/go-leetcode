package graph_test

import (
	"dsa/graph"
	"fmt"
	"strconv"
	"testing"
)

var UpdateBoardTests = []struct {
	board [][]byte
	click []int
}{
	{
		[][]byte{{'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'M', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'}},
		[]int{3, 0},
	},
}

func TestUpdateBoard(t *testing.T) {
	for i, e := range UpdateBoardTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := graph.UpdateBoard(e.board, e.click)
			for _, row := range actual {
				fmt.Println(string(row))
			}
		})
	}
}
