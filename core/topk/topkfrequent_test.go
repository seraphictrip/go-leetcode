package topk_test

import (
	"dsa/core/topk"
	"slices"
	"strconv"
	"testing"
)

var TopKFrequentTests = []struct {
	nums     []int
	k        int
	expected []int
}{
	{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	{[]int{1}, 1, []int{1}},
}

func TestTopKFrequent(t *testing.T) {
	for i, e := range TopKFrequentTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := topk.TopKFrequent(e.nums, e.k)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("TopKFrequent(%v, %v) = %v, want %v", e.nums, e.k, actual, e.expected)
			}
		})
	}
}
