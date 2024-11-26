package dp_test

import (
	"dsa/dp"
	"strconv"
	"testing"
)

var maxProductTests = []struct {
	nums     []int
	expected int
}{
	{[]int{2, 3, -2, 4}, 6},
	{[]int{-2, 0, -1}, 0},
	{[]int{1, 2, 3, 4, 5}, 120},
	{[]int{-1, 2, 3, 4, 5}, 120},
	{[]int{-1, -2, -3, -4, -5}, 120},
}

func TestMaxProduct(t *testing.T) {
	for i, e := range maxProductTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := dp.MaxProduct(e.nums)
			if actual != e.expected {
				t.Fatalf("MaxProduct(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}
