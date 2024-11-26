package kadanes

import (
	"strconv"
	"testing"
)

var kadanesTests = []struct {
	nums     []int
	expected int
}{

	{[]int{1}, 1},
	// [-1, 1]
	{[]int{-1, 1}, 1},
	// [2, 1]
	{[]int{2, -1}, 2},
	// [-2, 1, -2, 4, 3, 5, 6, 1, 5]
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	// [5, 9, 8, 15, 23]
	{[]int{5, 4, -1, 7, 8}, 23},
	// [4, 3, 5, -2, 3, 7]
	{[]int{4, -1, 2, -7, 3, 4}, 7},
}

func TestKadanes(t *testing.T) {
	for i, e := range kadanesTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := Kadanes(e.nums)
			if actual != e.expected {
				t.Fatalf("Kadanes(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}
