package bfs_test

import (
	"dsa/bfs"
	"fmt"
	"strconv"
	"testing"
)

var GetCoordinateTests = []struct {
	pos, n   int
	expected [2]int
}{
	{1, 6, [2]int{5, 0}},
	{12, 6, [2]int{4, 0}},
	{13, 6, [2]int{3, 0}},
	{24, 6, [2]int{2, 0}},
	{25, 6, [2]int{1, 0}},
	{36, 6, [2]int{0, 0}},
	{31, 6, [2]int{0, 5}},
	{30, 6, [2]int{1, 5}},
	{19, 6, [2]int{2, 5}},
	{18, 6, [2]int{3, 5}},
	{7, 6, [2]int{4, 5}},
	{6, 6, [2]int{5, 5}},
	{1, 3, [2]int{2, 0}},
	{2, 3, [2]int{2, 1}},
	{3, 3, [2]int{2, 2}},
	{4, 3, [2]int{1, 2}},
	{5, 3, [2]int{1, 1}},
	{6, 3, [2]int{1, 0}},
	{7, 3, [2]int{0, 0}},
	{8, 3, [2]int{0, 1}},
	{9, 3, [2]int{0, 2}},
}

func TestGetCoordinate(t *testing.T) {
	for i, e := range GetCoordinateTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			row, col := bfs.GetCoordinate(e.pos, e.n)
			actual := [2]int{row, col}
			if actual != e.expected {
				t.Fatalf("GetCoordinate(%v, %v) = %v, want %v", e.pos, e.n, actual, e.expected)
			}
		})
	}
}

var divTests = []struct {
}{}

func TestDiv(t *testing.T) {
	for i := 1; i <= 36; i++ {
		fmt.Println((i - 1) / 6)
	}
}
