package dp_test

import (
	"dsa/dp"
	"slices"
	"strconv"
	"testing"
)

var ProductExceptSelfTests = []struct {
	nums, expected []int
}{
	{},
	{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
}

func TestProductExceptSelf(t *testing.T) {
	for i, e := range ProductExceptSelfTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := dp.ProductExceptSelfNaive(e.nums)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("%v != %v", actual, e.expected)
			}
		})
	}
}

var MaxSubArrayTests = []struct {
	nums     []int
	expected int
}{
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	{[]int{1}, 1},
	{[]int{-1}, -1},
	{[]int{-1, 1}, 1},
	{[]int{1, -1}, 1},
	{[]int{2, 3, -8, 7, -1}, 7},
	{[]int{2, 3, -8, 7, -1, 2, 3}, 11},
}

func TestMaxSubArray(t *testing.T) {
	for i, e := range MaxSubArrayTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := dp.MaxSubArray(e.nums)
			if actual != e.expected {
				t.Fatalf("MaxSubArray(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}
