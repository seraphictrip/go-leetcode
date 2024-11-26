package linkedlist_test

import (
	"dsa/advanced/linkedlist"
	"slices"
	"strconv"
	"testing"
)

var MiddleTests = []struct {
	head          *linkedlist.ListNode
	expectedValue int
}{
	{linkedlist.FromArray([]int{1, 2, 3, 4, 5}), 3},

	// [1,2,3,4,5,6]
	// s/f
	//   s  f
	//      s   f
	//		  s	    f(nil)
	{linkedlist.FromArray([]int{1, 2, 3, 4, 5, 6}), 4},
}

func TestMiddle(t *testing.T) {
	for i, e := range MiddleTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := linkedlist.MiddleNode2Ptr(e.head)
			if actual.Val != e.expectedValue {
				t.Fatalf("MiddleNode(%v) = %v, want {%v}", e.head, actual, e.expectedValue)
			}
		})
	}
}

var FromArrayTests = []struct {
	nums []int
}{
	{[]int{}},
	{[]int{1, 2, 3, 4, 5}},
}

func TestNewListNode(t *testing.T) {
	for i, e := range FromArrayTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			head := linkedlist.FromArray(e.nums)
			collected := linkedlist.Collect(head)
			if !slices.Equal(e.nums, collected) {
				t.Fatalf("got %v, want %v", collected, e.nums)
			}
		})
	}
}
