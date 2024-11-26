package prefixsum_test

import (
	"dsa/advanced/prefixsum"
	"strconv"
	"testing"
)

var PivotIndexTests = []struct {
	nums     []int
	expected int
}{
	{[]int{1, 7, 3, 6, 5, 6}, 3},
	{[]int{2, -1, 1}, 0},
}

func TestPivotIndex(t *testing.T) {
	for i, e := range PivotIndexTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := prefixsum.PivotIndexBruteForce1(e.nums)
			if actual != e.expected {
				t.Fatalf("PivotIndex(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}
