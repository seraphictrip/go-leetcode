package arrandhash_test

import (
	"dsa/blind/arrandhash"
	"slices"
	"strconv"
	"testing"
)

var TwoSumTests = []struct {
	nums     []int
	target   int
	expected []int
}{
	{
		[]int{2, 7, 11, 15},
		9,
		[]int{0, 1},
	},
	{
		[]int{3, 2, 4}, 6, []int{1, 2},
	},
	{
		[]int{3, 3}, 6, []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for i, e := range TwoSumTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := arrandhash.TwoSumBruteForce(e.nums, e.target)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("TwoSum(%v, %v) = %v, want %v", e.nums, e.target, actual, e.expected)
			}
		})
	}
}
