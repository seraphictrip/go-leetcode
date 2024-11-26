package kclosest_test

import (
	"dsa/binarytree/heap/kclosest"
	"strconv"
	"testing"
)

var FindKthLargestTests = []struct {
	inputs []int
	k      int
	output int
}{
	{
		// quickselect
		// [3, 2, 1, 5, 6, 4] => [3, 2, 1, 4, 6, 5] ; k = 2
		// [6,5] => [5,6]; k = 2
		[]int{3, 2, 1, 5, 6, 4}, 2, 5,
	},
	{
		// quickselect
		// [3, 2, 3, 1, 2, 4, 5, 5, 6] => [3, 2, 3, 1, 2, 4, 5, 5, 6]; k = 4
		// [3, 2, 3, 1, 2, 4, 5, 5] => [3, 2, 3, 1, 2, 4, 5, 5]; k = 3
		// [3, 2, 3, 1, 2, 4, 5] = > [3, 2, 3, 1, 2, 4, 5]; k = 2
		// [3, 2, 3, 1, 2, 4] => [3, 2, 3, 1, 2, 4]; k = 1
		[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4,
	},
	{
		[]int{3, 5, 6, 1, 2, 4, 2, 5, 3}, 4, 4,
	},
}

func TestFindKthLargest(t *testing.T) {
	for i, e := range FindKthLargestTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			output := kclosest.FindKthLargestQuickSelect(e.inputs, e.k)
			if output != e.output {
				t.Fatalf("FindKthLargest(%v, %v) = %v, want %v", e.inputs, e.k, output, e.output)
			}
		})
	}
}
