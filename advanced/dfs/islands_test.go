package dfs_test

import (
	"dsa/advanced/dfs"
	"strconv"
	"testing"
)

var islandsTests = []struct {
	matrix   [][]byte
	expected int
}{
	{[][]byte{{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}, 1},
}

func TestIslands(t *testing.T) {
	for i, e := range islandsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := dfs.NumOfIslands(e.matrix)
			if actual != e.expected {
				t.Fatalf("NumOfIslands(%v) = %v, want %v", e.matrix, actual, e.expected)
			}
		})
	}
}
