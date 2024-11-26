package random_test

import (
	"dsa/random"
	"slices"
	"strconv"
	"testing"
)

var AddTwoNumbersTests = []struct {
	l1, l2   *random.ListNode
	expected []int
}{
	{toLinkedList([]int{7}), toLinkedList([]int{2}), []int{9}},
	{
		toLinkedList([]int{2, 4, 3}),
		toLinkedList([]int{5, 6, 4}),
		[]int{7, 0, 8},
	},
	{
		toLinkedList([]int{9, 9, 9, 9, 9, 9, 9}),
		toLinkedList([]int{9, 9, 9, 9}),
		[]int{8, 9, 9, 9, 0, 0, 0, 1},
	},
}

func TestAddTwoNumbers(t *testing.T) {
	for i, e := range AddTwoNumbersTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := random.AddTwoNumbers(e.l1, e.l2)
			if !slices.Equal(Collect(actual), e.expected) {
				t.Fatalf("AddTwoNumbers(%v, %v) = %v, want %v", Collect(e.l1), Collect(e.l2), Collect(actual), e.expected)
			}
		})
	}
}

func Collect(node *random.ListNode) []int {
	result := make([]int, 0)

	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}

	return result
}

func toLinkedList(nums []int) *random.ListNode {
	head := &random.ListNode{}
	cur := head
	for _, num := range nums {
		node := &random.ListNode{Val: num}
		cur.Next = node
		cur = node
	}

	return head.Next
}
