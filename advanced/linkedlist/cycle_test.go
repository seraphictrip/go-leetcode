package linkedlist_test

import (
	"dsa/advanced/linkedlist"
	"fmt"
	"slices"
	"strconv"
	"testing"
)

var CycleStartTests = []struct {
}{}

func TestCycleStart(t *testing.T) {
	for i, e := range CycleStartTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			fmt.Println(e)
		})
	}
}

var FromArrayCycleTests = []struct {
	nums     []int
	expected []int
}{
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{1, 2, 3, 4, 5, 2, 1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 2}},
}

func TestFromArrayCycle(t *testing.T) {
	for i, e := range FromArrayCycleTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			head := linkedlist.FromArrayCycle(e.nums)
			actual := linkedlist.CollectCycle(head)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("FromArrayCycle(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}
