package rmdupes

import (
	"strconv"
	"strings"
)

/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;

	for (int i = 0; i < k; i++) {
	    assert nums[i] == expectedNums[i];
	}

If all assertions pass, then your solution will be accepted.

Example 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Constraints:

1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order
*/
func RemoveDuplicates(nums []int) int {
	l := 1
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}

/*
You are given the head of a linked list.

Remove every node which has a node with a greater value anywhere to the right side of it.

Return the head of the modified linked list.



Example 1:


Input: head = [5,2,13,3,8]
Output: [13,8]
Explanation: The nodes that should be removed are 5, 2 and 3.
- Node 13 is to the right of node 5.
- Node 13 is to the right of node 2.
- Node 8 is to the right of node 3.
Example 2:

Input: head = [1,1,1,1]
Output: [1,1,1,1]
Explanation: Every node has value 1, so no nodes are removed.


Constraints:

The number of the nodes in the given list is in the range [1, 105].
1 <= Node.val <= 105

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return "nil"
	}
	return strings.Join([]string{strconv.Itoa(n.Val), "=>", n.Next.String()}, "")
}

type Stack struct {
	data []*ListNode
}

func NewStack() *Stack {
	return &Stack{
		data: make([]*ListNode, 0),
	}
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(node *ListNode) {
	s.data = append(s.data, node)
}

func (s *Stack) Pop() *ListNode {
	item := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return item
}

func (s *Stack) Top() *ListNode {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

// [5, 2, 13, 3, 8]
// S: [] cur = 5
// S: [5], cur = 2
// S: [5,2], curr = 13
//		S: [13]. curr = 3
// S: [13, 3], curr =8
//		S: [13, 8]
// S: [13], cur(8).Next = nil. next p cur(8)

func RemoveNodes(head *ListNode) *ListNode {
	cur := head
	stack := NewStack()

	for cur != nil {
		for !stack.Empty() && stack.Top().Val < cur.Val {
			stack.Pop()
		}
		stack.Push(cur)
		cur = cur.Next
	}

	var next *ListNode = nil
	for !stack.Empty() {
		cur = stack.Pop()
		cur.Next = next
		next = cur
	}

	return cur

}

func New(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func FromArray(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := New(arr[0])
	head.Next = FromArray(arr[1:])
	return head
}
