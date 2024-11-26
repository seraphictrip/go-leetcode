package bfs_test

import (
	"dsa/bfs"
	"slices"
	"strconv"
	"testing"
)

var BoustrophedonWalkTests = []struct {
	board    [][]int
	expected []int
}{
	{
		[][]int{
			{36, 35, 34, 33, 32, 31},
			{25, 26, 27, 28, 29, 30},
			{24, 23, 22, 21, 20, 19},
			{13, 14, 15, 16, 17, 18},
			{12, 11, 10, 9, 8, 7},
			{1, 2, 3, 4, 5, 6},
		},
		Sequence(1, 36),
	},
	{
		[][]int{
			{21, 22, 23, 24, 25},
			{20, 19, 18, 17, 16},
			{11, 12, 13, 14, 15},
			{10, 9, 8, 7, 6},
			{1, 2, 3, 4, 5},
		},
		Sequence(1, 25),
	},
}

func Sequence(start, end int) []int {
	seq := make([]int, 0)
	for i := start; i <= end; i++ {
		seq = append(seq, i)
	}
	return seq
}

func TestBoustrophedonWalk(t *testing.T) {
	for i, e := range BoustrophedonWalkTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := bfs.BoustrophedonWalk(e.board)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("BoustrophedonWalk(%v) = %v, want %v", e.board, actual, e.expected)
			}
		})
	}
}
