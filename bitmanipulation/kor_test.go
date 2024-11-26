package bitmanipulation_test

import (
	"dsa/bitmanipulation"
	"strconv"
	"testing"
)

var FindKOrTests = []struct {
	nums     []int
	k        int
	expected int
}{
	{
		[]int{7, 12, 9, 8, 9, 15},
		4,
		9,
	},
}

func TestFindKOr(t *testing.T) {
	for i, e := range FindKOrTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := bitmanipulation.FindKOr(e.nums, e.k)
			if actual != e.expected {
				t.Fatalf("FindKOr(%v, %v) = %v, want %v", e.nums, e.k, actual, e.expected)
			}
		})
	}
}
