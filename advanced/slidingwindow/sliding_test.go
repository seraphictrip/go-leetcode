package slidingwindow_test

import (
	"dsa/advanced/slidingwindow"
	"strconv"
	"testing"
)

var CloseDuplicatesBruteForceTests = []struct {
	arr      []int
	k        int
	expected bool
}{
	{[]int{1, 2, 3, 2, 3, 3}, 2, true},
	{[]int{1, 2, 3, 2, 3, 3}, 3, true},
	{[]int{1, 2, 3, 2, 3, 3}, 1, false},
	{[]int{1, 2, 3, 2, 3, 3}, 12, true},
	{[]int{1, 2, 3, 4, 5, 6}, 12, false},
}

func TestCloseDuplicatesBruteForce(t *testing.T) {
	for i, e := range CloseDuplicatesBruteForceTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := slidingwindow.CloseDuplicatesWithHash(e.arr, e.k)
			if actual != e.expected {
				t.Fatalf("CloseDuplicatesBruteForce(%v, %v) = %v, want %v", e.arr, e.k, actual, e.expected)
			}
		})
	}
}

var LongestSameSubArrayTests = []struct {
	arr      []int
	expected int
}{
	{[]int{1, 2, 2, 3, 2, 3, 3, 3}, 3},
}

func TestLongestSameSubArray(t *testing.T) {
	for i, e := range LongestSameSubArrayTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := slidingwindow.LongestSameSubArray(e.arr)
			if actual != e.expected {
				t.Fatalf("LongestSameSubArray(%v) = %v, want %v", e.arr, actual, e.expected)
			}
		})
	}
}

var ShortestSubarrayTests = []struct {
	arr      []int
	target   int
	expected int
}{
	{[]int{2, 3, 1, 2, 4, 3}, 6, 2},
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 60, 0},
}

func TestShortestSubarray(t *testing.T) {
	for i, e := range ShortestSubarrayTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := slidingwindow.ShortestSubarray(e.arr, e.target)
			if actual != e.expected {
				t.Fatalf("ShortestSubarray(%v, %v) = %v, want %v", e.arr, e.target, actual, e.expected)
			}
		})
	}
}
