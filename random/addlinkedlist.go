package random

import (
	"slices"
)

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed tha
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sent := &ListNode{}
	cur := sent
	carry := 0
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		c := a + b + carry
		carry = c / 10
		c = c % 10
		cur.Next = &ListNode{Val: c}
		cur = cur.Next

	}
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return sent.Next
}

type intstack []int

func (s *intstack) Push(val int) {
	*s = append(*s, val)
}

func (s *intstack) Pop() int {
	old := *s
	topIndex := len(old) - 1
	top := old[topIndex]
	*s = old[:topIndex]
	return top
}

func (s intstack) IsEmpty() bool {
	return len(s) == 0
}

func removeAnagrams(words []string) []string {
	i := 1
	for i < len(words) {
		if isAnagram(words[i-1], words[i]) {
			if i == len(words)-1 {
				words = words[:len(words)-1]
			} else {
				words = append(words[:i], words[i+1:]...)
			}

			i = 1
		} else {
			i++
		}
	}

	return words
}

func isAnagram(w1, w2 string) bool {
	b1 := []byte(w1)
	b2 := []byte(w2)
	slices.Sort(b1)
	slices.Sort(b2)
	return string(b1) == string(b2)
}
