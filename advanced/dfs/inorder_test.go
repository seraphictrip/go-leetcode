package dfs_test

import (
	"dsa/advanced/dfs"
	"slices"
	"strconv"
	"testing"
)

var CreateInorderTests = []struct {
	nums []int
}{
	{[]int{1, 2, 3, 4, 5}},
}

func TestCreateInorder(t *testing.T) {
	for i, e := range CreateInorderTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			tree := dfs.CreateInorder(e.nums)
			actual := dfs.InorderIterative(tree)
			if !slices.Equal(actual, e.nums) {
				t.Fatalf("%v != %v", actual, e.nums)
			}
		})
	}
}
