package laststone_test

import (
	"dsa/binarytree/heap/laststone"
	"strconv"
	"testing"
)

var LastStoneWeightTests = []struct {
	stones   []int
	expected int
}{
	// {},
	{[]int{2, 7, 4, 1, 8, 1}, 1},
	{[]int{1}, 1},
	{[]int{1, 1}, 0},
	{[]int{6, 1, 1, 1, 1}, 2},
}

func TestLastStoneWeight(t *testing.T) {
	for i, e := range LastStoneWeightTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := laststone.LastStoneWeightRecursive(e.stones)
			if actual != e.expected {
				t.Fatalf("LastStoneWeight(%v) = %v, want %v", e.stones, actual, e.expected)
			}
		})
	}
}
