package knapsack_test

import (
	"dsa/dp/patterns/knapsack"
	"strconv"
	"testing"
)

var CanPartitionTests = []struct {
	nums     []int
	expected bool
}{
	{[]int{1, 5, 11, 5}, true},
	{[]int{1, 2, 3, 5}, false},
	{[]int{1, 2, 5}, false},
	{[]int{43, 87, 61, 26, 73, 64, 23, 9, 54, 100, 14, 47, 75, 49, 90, 50, 62, 96, 18, 86, 95, 27, 87, 67, 67, 92, 82, 19, 53, 86, 15, 37, 83, 81, 59, 84, 47, 11, 80, 6, 14, 58, 72, 13, 78, 31, 56, 72, 94, 79, 67, 86, 25, 85, 19, 54, 50, 11, 52, 95, 100, 37, 96, 88, 71, 45, 77, 58, 13, 12, 49, 83, 50, 20, 2, 54, 84, 51, 3, 25, 42, 30, 92, 35, 91, 68, 57, 19, 4, 87, 15, 17, 6, 94, 84, 85, 91, 31, 47, 33}, true},
}

func TestCanPartition(t *testing.T) {
	for i, e := range CanPartitionTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := knapsack.CanPartitionHashSet(e.nums)
			if actual != e.expected {
				t.Fatalf("CanPartition(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}
