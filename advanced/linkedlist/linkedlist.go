package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func FromArray(arr []int) *ListNode {
	n := len(arr)
	if n == 0 {
		return nil
	}
	head := NewListNode(arr[0])
	cur := head
	for i := 1; i < n; i++ {
		cur.Next = NewListNode(arr[i])
		cur = cur.Next
	}
	return head
}

// [1,3,4,2,2]
// given an array, create a cycle anytime see same value
// this is a helper method for testing cycle code but also useful for gaining
// intution for https://leetcode.com/problems/find-the-duplicate-number/description/
func FromArrayCycle(arr []int) *ListNode {
	seen := map[int]*ListNode{}

	n := len(arr)
	if n == 0 {
		return nil
	}
	head := NewListNode(arr[0])
	cur := head
	for i := 1; i < n; i++ {
		if node, ok := seen[arr[i]]; ok {
			// we create a cycle, no need to go further
			cur.Next = node
			break
		}
		cur.Next = NewListNode(arr[i])
		cur = cur.Next
		seen[arr[i]] = cur
	}
	return head
}

func Collect(head *ListNode) []int {
	result := make([]int, 0)

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func CollectCycle(head *ListNode) []int {
	result := make([]int, 0)
	seen := make(map[*ListNode]bool)
	for head != nil {
		result = append(result, head.Val)
		if seen[head] {
			break
		}
		seen[head] = true
		head = head.Next
	}

	return result
}
