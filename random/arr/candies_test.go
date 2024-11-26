package arr_test

import (
	"dsa/random/arr"
	"slices"
	"strconv"
	"testing"
)

var KidsWithCandiesTests = []struct {
	nums     []int
	extra    int
	expected []bool
}{
	{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
	{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
	{[]int{12, 1, 2}, 10, []bool{true, false, true}},
}

func TestKidsWithCandies(t *testing.T) {
	for i, e := range KidsWithCandiesTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := arr.KidsWithCandies(e.nums, e.extra)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("KidsWithCandies(%v, %v) = %v, want %v", e.nums, e.extra, actual, e.expected)
			}
		})
	}
}
