package linkedlist

import (
	"fmt"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("{%v}->%v", n.Val, n.Next)
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	stack := make([]*ListNode, 0)
	for head != nil {
		// push head onto stack
		stack = append(stack, head)
		head = head.Next
	}
	// stack ready to be popped in right order
	// head = nil
	newHead := stack[len(stack)-1]
	head = newHead
	stack = stack[:len(stack)-1]
	for topIndex := len(stack) - 1; topIndex >= 0; topIndex-- {
		top := stack[topIndex]
		head.Next = top
		head = top
		head.Next = nil
		// pop, though not entirely necessary
		stack = stack[:topIndex]
	}
	return newHead
}

func ReverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev, curr *ListNode = nil, head
	for curr != nil {
		tmp := curr.Next
		// reverse pointer
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}

func ReverseList3(head *ListNode) *ListNode {
	return rreverse(head, nil)
}

func rreverse(head, prev *ListNode) *ListNode {
	if head == nil {
		return prev
	}
	tmp := head.Next
	head.Next = prev
	return rreverse(tmp, head)
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func (l LinkedList) String() string {
	if l.head == nil {
		return "Empty()"
	}
	if l.head == l.tail {
		return fmt.Sprintf("Just(%v)", l.head.Val)
	}
	return fmt.Sprintf("(Head({%v}), Tail(%v))", l.head.Val, l.head.Next)
}

func (l *LinkedList) Head() *ListNode {
	return l.head
}

func (l *LinkedList) Tail() *ListNode {
	return l.tail
}

// Creational FromArray, this will overwrite any existing data
func (l *LinkedList) FromArray(items []int) {
	// special case, zero out array
	if len(items) == 0 {
		l.head = nil
		l.tail = nil
		return
	}
	// head and tail same in Just case
	head := &ListNode{items[0], nil}
	tail := head

	for i := 1; i < len(items); i++ {
		tail.Next = &ListNode{items[i], nil}
		tail = tail.Next
	}

	l.head = head
	l.tail = tail
}

func (ll *LinkedList) Append(val int) {
	node := &ListNode{val, nil}
	ll.AppendNode(node)
}

func (ll *LinkedList) AppendMany(vals ...int) {
	for _, val := range vals {
		ll.Append(val)
	}
}

func (ll *LinkedList) AppendNode(node *ListNode) {
	if ll.tail == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.Next = node
		ll.tail = node
	}
}

func (ll *LinkedList) Prepend(val int) {
	node := &ListNode{val, ll.head}
	ll.PrependNode(node)
}

func (ll *LinkedList) PrependNode(node *ListNode) {
	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.head = node
	}
}

func (ll *LinkedList) Delete(val int) bool {
	node := ll.FindNodeByVal(val)
	return ll.DeleteNode(node)
}

func (ll *LinkedList) DeleteNode(node *ListNode) bool {
	// handle nil node, handle head nil
	if node == nil || ll.head == nil {
		return false
	}
	// handle node == head
	// 1->2->3
	if node == ll.head {
		if ll.head == ll.tail {
			ll.tail = nil
		}
		ll.head = ll.head.Next
		return true
	}

	// handle middle
	curr := ll.head
	for curr != nil {
		if curr.Next == node {
			// handle tail
			if ll.tail == node {
				ll.tail = curr
			}
			curr.Next = node.Next
			return true
		}
		curr = curr.Next
	}
	return false
}

func (ll *LinkedList) FindNodeByVal(val int) *ListNode {
	if ll.head == nil {
		return nil
	}
	for curr := ll.head; curr != nil; curr = curr.Next {
		if curr.Val == val {
			return curr
		}
	}
	return nil
}

// Equal LinkedList have same ref OR have same elements in same order
func (ll *LinkedList) Equals(other *LinkedList) bool {
	if ll == other {
		// same ptr
		return true
	}
	if ll == nil || other == nil {
		return false
	}

	return slices.Equal(ll.ToArray(), other.ToArray())
}

// Convert to an array
func (l *LinkedList) ToArray() []int {
	result := make([]int, 0)

	for next := l.head; next != nil; next = next.Next {
		result = append(result, next.Val)
	}
	return result
}

// Static Methods
//
// Creational, create linked list from an arry
func FromArray(vals []int) *LinkedList {
	ll := new(LinkedList)
	ll.FromArray(vals)

	return ll
}

func Just(val int) *LinkedList {
	return FromArray([]int{val})
}

func Empty() *LinkedList {
	return new(LinkedList)
}

func Zero() *LinkedList {
	return Empty()
}

func NewLinkedList(vals ...int) *LinkedList {
	return FromArray(vals)
}

// ToArray
func ToArray(ll *LinkedList) []int {
	return ll.ToArray()
}

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
*/

type DLLNode struct {
	Val  int
	Next *DLLNode
	Prev *DLLNode
}

func (n *DLLNode) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("{%d}->%v", n.Val, n.Next)
}

func NewDLLNode(val int, next, prev *DLLNode) *DLLNode {
	return &DLLNode{val, next, prev}
}

type MyLinkedList struct {
	// Left bound sentenial so I can treat head/tail the same
	left  *DLLNode
	right *DLLNode
}

func MyLinkedListFromArray([]int) *MyLinkedList {
	mll := Constructor()
	return mll
}

func Constructor() *MyLinkedList {
	right := NewDLLNode(0, nil, nil)
	left := NewDLLNode(0, right, nil)
	right.Prev = left

	return &MyLinkedList{left, right}
}

func (ll *MyLinkedList) String() string {
	if ll.left.Next == ll.right {
		// out of bounds
		return "nil"
	}
	return ll.left.Next.String()
}

func (ll *MyLinkedList) Len() int {
	l := 0
	head := ll.left.Next
	for head != ll.right {
		l++
		head = head.Next
	}
	return l
}

func (ll *MyLinkedList) FindNodeAtIndex(index int) *DLLNode {
	curr := ll.left.Next // head
	for curr != ll.right && index >= 0 {
		curr = curr.Next
		index--
	}
	if index != 0 || curr == ll.right {
		return nil
	}
	return curr
}

func (ll *MyLinkedList) Get(index int) int {
	return 0
}

func (ll *MyLinkedList) AddAtHead(val int) {

}

func (ll *MyLinkedList) AddAtTail(val int) {

}

// void addAtIndex(int index, int val) Add a node of value val before the indexth node in the linked list.
// If index equals the length of the linked list, the node will be appended to the end of the linked list.
// If index is greater than the length, the node will not be inserted.
func (ll *MyLinkedList) AddAtIndex(index int, val int) {

}

func (ll *MyLinkedList) DeleteAtIndex(index int) {

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
