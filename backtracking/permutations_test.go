package backtracking_test

import (
	bt "dsa/backtracking"
	"fmt"
	"slices"
	"strconv"
	"testing"
)

var PermuteTests = []struct {
	nums     []int
	expected [][]int
}{
	{[]int{1, 2, 3}, [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{2, 1, 3},
	}},
}

func TestPermute(t *testing.T) {
	for i, e := range PermuteTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := bt.Permute(e.nums)
			AssertPermutations(t, e.nums, actual, e.expected)
		})
	}
}

func AssertPermutations(t *testing.T, nums []int, actual, expected [][]int) {
	t.Helper()
	sortPerms(actual)
	sortPerms(expected)
	fmt.Println(actual)
	if len(actual) != len(expected) {
		t.Fatalf("Permute(%v) = %v, want %v", nums, actual, expected)
	}
}

func sortPerms(perms [][]int) {
	slices.SortFunc(perms, func(a, b []int) int {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return a[i] - b[i]
			}
		}
		return 0
	})
}
