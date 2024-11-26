package linkedlist

/*
Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.



Example 1:


Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.
Example 2:


Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.


Constraints:

The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100
*/

func MiddleNode(head *ListNode) *ListNode {
	n := length(head)
	mid := n / 2

	for i := 0; i < mid; i++ {
		head = head.Next
	}
	return head
}

func MiddleNode2Ptr(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func length(head *ListNode) int {
	if head == nil {
		return 0
	}
	return 1 + length(head.Next)
}
