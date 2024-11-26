package mylinkedlist

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next is a pointer/reference to the next node.
If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

Implement the MyLinkedList class:

MyLinkedList() Initializes the MyLinkedList object.
int get(int index) Get the value of the indexth node in the linked list. If the index is invalid, return -1.
void addAtHead(int val) Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
void addAtTail(int val) Append a node of value val as the last element of the linked list.
void addAtIndex(int index, int val) Add a node of value val before the indexth node in the linked list. If index equals the length of the linked list, the node will be appended to the end of the linked list. If index is greater than the length, the node will not be inserted.
void deleteAtIndex(int index) Delete the indexth node in the linked list, if the index is valid.


Example 1:

Input
["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
[[], [1], [3], [1, 2], [1], [1], [1]]
Output
[null, null, null, null, 2, null, 3]

Explanation
MyLinkedList myLinkedList = new MyLinkedList();
myLinkedList.addAtHead(1);
myLinkedList.addAtTail(3);
myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
myLinkedList.get(1);              // return 2
myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
myLinkedList.get(1);              // return 3


Constraints:

0 <= index, val <= 1000
Please do not use the built-in LinkedList library.
At most 2000 calls will be made to get, addAtHead, addAtTail, addAtIndex and deleteAtIndex.
*/

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

func (node ListNode) String() string {
	str := ""
	if node.Prev != nil {
		str = "<{"
	} else {
		str = "{"
	}
	str += strconv.Itoa(node.Val) + "}"
	if node.Next != nil {
		str += ">"
	}
	return str
}

func NewListNode(val int, next, prev *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
		Prev: prev,
	}
}

type MyLinkedList struct {
	// Left Bound
	left *ListNode
	// Right Bound
	right *ListNode
}

func Constructor() MyLinkedList {
	left := NewListNode(0, nil, nil)
	right := NewListNode(0, nil, left)
	left.Next = right
	return MyLinkedList{left, right}
}

func (ll MyLinkedList) String() string {
	parts := []string{"L|"}
	ll.Visit(func(node *ListNode) bool {
		parts = append(parts, fmt.Sprintf("%v", node))
		return false
	})
	parts = append(parts, "|R")
	return strings.Join(parts, "")
}

func (ll *MyLinkedList) GetNode(index int) *ListNode {
	var node *ListNode = nil
	i := 0
	ll.Visit(func(n *ListNode) bool {
		if i == index {
			node = n
			return true
		}
		i++
		return false
	})
	return node
}

func (ll *MyLinkedList) Visit(vistor func(node *ListNode) bool) {
	head := ll.left.Next
	for head != nil && head != ll.right {
		if vistor(head) {
			break
		}
		head = head.Next
	}
}

func (ll *MyLinkedList) Get(index int) int {
	node := ll.GetNode(index)
	if node == nil {
		return -1
	}
	return node.Val
}

// L|<{val}><
func (ll *MyLinkedList) AddAtHead(val int) {
	head := ll.left.Next
	// L| <{val}> {head}>
	node := NewListNode(val, head, ll.left)
	// L|><{val}> {head}>
	ll.left.Next = node
	// L|><{val}><{head}>
	head.Prev = node

}

func (ll *MyLinkedList) AddAtTail(val int) {
	tail := ll.right.Prev
	tail.Next = NewListNode(val, ll.right, tail)
	ll.right.Prev = tail.Next
}

// Add Node at index
// (<{0}><{1}><{2}>).AddAtIndex(0, 100) => (<{100}><{0}><{1}><{2}>)
// (<{0}><{1}><{2}>).AddAtIndex(1, 100) => (<{0}><{100}><{1}><{2}>)
// // (<{0}><{1}><{2}>).AddAtIndex(3, 100) => (<{0}><{1}><{2}><{100}>)
func (ll *MyLinkedList) AddAtIndex(index int, val int) {
	// get True head
	cur := ll.left.Next
	// <{0}><{1}><{2}>
	for cur.Next != nil && index > 0 {
		cur = cur.Next // 1. <{1}> 2. <{2}> 3. <{right}>
		index--
	}

	if index == 0 {
		// <{cur.prev} <{val}> {cur}>
		node := NewListNode(val, cur, cur.Prev)
		// <{cur.prev}><{val}> {cur}>
		cur.Prev.Next = node
		// <{cur.prev}><{val}> <{cur}>
		cur.Prev = node
	}

}

// Delete Node at index
// (<{0}><{1}><{2}>).DeleteAtIndex(0) => (<{1}><{2}>)
func (ll *MyLinkedList) DeleteAtIndex(index int) {
	// grab true head <{0}><{1}><{2}>
	cur := ll.left.Next // <{0}>
	for cur.Next != nil && index > 0 {
		cur = cur.Next // <{1}>
		index--
	}
	if index == 0 && cur != nil && cur != ll.right {
		// prev=<{0}> <{1}> next=<{2}>
		// <{0}><{2}>
		prev, next := cur.Prev, cur.Next
		prev.Next, next.Prev = next, prev
	}
}

// O(n)
func (ll *MyLinkedList) Len() int {
	// true head
	cur := ll.left.Next
	i := 0
	for cur != nil && cur != ll.right {
		cur = cur.Next
		i++
	}

	return i
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
