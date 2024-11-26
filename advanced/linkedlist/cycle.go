package linkedlist

func CycleStart(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// cycle detected
			break
		}
	}

	if fast == nil || fast.Next == nil {
		// we completed without cycle, so advertise no cycle
		return nil
	}
	// we detected a cycle, so start another slow pointer to meet existing slow at cycle head
	// P = distance from head to cycle (maybe zero)
	// C = cycle distance
	// C-X = distance from start of cycle to fast/slow intersect
	// X = distance from fast/slow intersect to head of cycle
	// 2*slow = fast
	// slow = P + C - X // slow pointer travel all of P, and C-X of C
	// fast = 2*slow = 2(P+C-X) = 2P + 2C -2x
	// fast = P + C + C -X // fast traveled all of P, all of C, and C-X to intersect
	// fast = 2P + 2C -2x = P + C + C -X
	// P-X = 0
	// P = X therefore the distance from P to start of cycle is same distance initial slow must travel to reach
	slow2 := head
	for slow != slow2 {
		slow = slow.Next
		slow2 = slow2.Next
	}
	return slow
}

func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]bool{}
	for head != nil {
		if seen[head] {
			return head
		}
		seen[head] = true
		head = head.Next
	}
	return nil
}
