package cycle

import "fmt"

type ListNode struct {
	Val     int
	Next    *ListNode
	IsCycle bool
}

func FromArray(items []int) *ListNode {
	head := &ListNode{}
	if len(items) == 0 {
		return head
	}
	head.Val = items[0]
	node := head
	for i := 1; i < len(items); i++ {
		node.Next = &ListNode{Val: items[i]}
		node = node.Next
	}
	return head
}

// Mostly used for side effects, loop over list and "visit" each node
// break if process function returns false
func Visit(node *ListNode, fn func(n *ListNode) bool) {
	if node == nil {
		return
	}
	head := node
	for head != nil {
		if fn(head) {
			break
		}
		head = head.Next
	}
}

func FromArrayCycleAt(items []int, at int) *ListNode {
	list := FromArray(items)
	i := 0
	var cycleNode *ListNode
	Visit(list, func(node *ListNode) bool {
		if i == at {
			cycleNode = node
		}
		i++

		if node.Next == nil {
			node.Next = cycleNode
			return true
		}

		return false
	})

	return list
}

func (node *ListNode) String() string {
	if node == nil {
		return "nil"
	}

	if node.IsCycle {
		return fmt.Sprintf("CYCLE(%v)->...", node.Val)
	}
	return fmt.Sprintf("{%v}->", node.Val)
}

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			fast.IsCycle = true
			return true
		}
		if fast.IsCycle || slow.IsCycle {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

func Print(node *ListNode) {
	i := 0
	cycleCount := 0
	for node != nil {
		fmt.Print(node)
		if node.IsCycle {
			cycleCount++
			if cycleCount > 2 || i > 100 {
				break
			}

		}
		i++
		node = node.Next
	}
	fmt.Println()

}
