package rmdupes_test

import (
	"dsa/arrays/rmdupes"
	"strconv"
	"testing"
)

var RemoveNodesTests = []struct {
	head     *rmdupes.ListNode
	expected *rmdupes.ListNode
}{
	{rmdupes.FromArray([]int{1, 1, 1, 1}), rmdupes.FromArray([]int{1, 1, 1, 1})},
	{
		rmdupes.FromArray([]int{5, 2, 13, 3, 8}),
		rmdupes.FromArray([]int{13, 8}),
	},
	{
		rmdupes.FromArray([]int{13, 3, 8}),
		rmdupes.FromArray([]int{13, 8}),
	},
	{
		rmdupes.FromArray([]int{5, 2, 13, 3, 8, 9}),
		rmdupes.FromArray([]int{13, 9}),
	},
}

func TestRemoveNodes(t *testing.T) {
	for i, e := range RemoveNodesTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := rmdupes.RemoveNodes(e.head)
			if !Equal(actual, e.expected) {
				t.Fatalf("RemoveNodes(%v) = %v, want %v", e.head, actual, e.expected)
			}

		})
	}
}

var FromArrayTests = []struct {
	arr []int
}{
	{[]int{1, 2}},
}

func TestFromArray(t *testing.T) {
	for i, e := range FromArrayTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			head := rmdupes.FromArray(e.arr)

			for i := range e.arr {
				if head.Val != e.arr[i] {
					t.Fatalf("fail")
				}
				head = head.Next
			}
		})
	}
}

func Equal(a, b *rmdupes.ListNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Val == b.Val {
		return Equal(a.Next, b.Next)
	}
	return false
}
