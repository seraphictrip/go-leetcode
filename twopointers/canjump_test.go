package twopointers_test

import (
	"dsa/twopointers"
	"strconv"
	"testing"
)

var CanJumpTests = []struct {
	nums     []int
	expected bool
}{
	{[]int{2, 5, 0, 0}, true},
}

func TestCanJump(t *testing.T) {
	for i, e := range CanJumpTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := twopointers.CanJump(e.nums)
			if actual != e.expected {
				t.Fatalf("CanJump(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}

var JumpTests = []struct {
	nums     []int
	expected int
}{
	{[]int{2, 1, 1, 1, 1}, 3},
	{[]int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}, 5},
}

func TestJump(t *testing.T) {
	for i, e := range JumpTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := twopointers.Jump(e.nums)
			if actual != e.expected {
				t.Fatalf("Jump(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}
