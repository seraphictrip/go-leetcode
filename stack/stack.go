package stack

import (
	"cmp"
	"errors"
	"fmt"
)

var (
	ErrUnderflow  = errors.New("underflow")
	ErrEmptyStack = errors.New("empty stack")
)

type Node[T any] struct {
	data T
	next *Node[T]
}

// Time Complexity: O(n)
func (n *Node[T]) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("(%v)->%v", n.data, n.next)
}

// Time Complexity: O(1)
func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		data: data,
	}
}

type Stack[T any] struct {
	head *Node[T]
}

// Time Complexity: O(n) taken from node
func (s *Stack[T]) String() string {
	return s.head.String()
}

// Time Complexity: O(1)
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Time Complexity: O(1)
func (s *Stack[T]) IsEmpty() bool {
	return s.head == nil
}

// Time Complexity: O(n)
func (s *Stack[T]) Len() int {
	l := 0
	head := s.head
	for head != nil {
		l++
		head = head.next
	}
	return l
}

// Time Complexity: O(1)
func (s *Stack[T]) Push(item T) {
	node := NewNode(item)
	node.next = s.head
	s.head = node
}

// Time Complexity: O(1)
func (s *Stack[T]) Top() (result T, err error) {
	if s.IsEmpty() {
		return result, ErrUnderflow
	}
	return s.head.data, nil
}

// Time Complexity: O(1)
func (s *Stack[T]) Pop() error {
	if s.IsEmpty() {
		return ErrUnderflow
	}
	s.head = s.head.next
	return nil
}

type MonotonicIncreasingStack[T cmp.Ordered] struct {
	*Stack[T]
}

func NewMonotonicIncreasingStack[T cmp.Ordered]() *MonotonicIncreasingStack[T] {
	return &MonotonicIncreasingStack[T]{
		New[T](),
	}
}

// Time Complexity: O(n)
func (s *MonotonicIncreasingStack[T]) Push(item T) {
	if s.IsEmpty() {
		s.Stack.Push(item)
		return
	}

	top, err := s.Top()
	for !s.IsEmpty() && err == nil && top > item {
		if top > item {
			s.Pop()
			top, err = s.Top()
		}
	}
	s.Stack.Push(item)
}

type MonotonicDecreasingStack[T cmp.Ordered] struct {
	*Stack[T]
}

func NewMonotonicDecreasingStack[T cmp.Ordered]() *MonotonicDecreasingStack[T] {
	return &MonotonicDecreasingStack[T]{
		New[T](),
	}
}

func (s *MonotonicDecreasingStack[T]) Push(item T) {
	// Invariant: stack is monotonically descreasing
	// which is to item must be less than anything else on the stack
	if s.IsEmpty() {
		s.Stack.Push(item)
	}

	top, _ := s.Top()
	for !s.IsEmpty() && item > top {
		s.Pop()
	}
	s.Stack.Push(item)
}

type StackSlice[T any] struct {
	stack []T
}

func NewStackSlice[T any]() *StackSlice[T] {
	return &StackSlice[T]{
		stack: []T{},
	}
}

func (s *StackSlice[T]) Push(item T) {
	s.stack = append(s.stack, item)
}

func (s *StackSlice[T]) Pop() (val T, err error) {
	if s.IsEmpty() {
		return val, ErrEmptyStack
	}
	val, _ = s.Top()
	s.stack = s.stack[0 : len(s.stack)-1]
	return val, nil
}

func (s *StackSlice[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *StackSlice[T]) Top() (val T, err error) {
	if s.IsEmpty() {
		return val, ErrEmptyStack
	}
	return s.stack[len(s.stack)-1], nil
}

/*
Design a stack class that supports the push, pop, top, and getMin operations.

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
Each function should run in
O
(
1
)
O(1) time.

Example 1:

Input: ["MinStack", "push", 1, "push", 2, "push", 0, "getMin", "pop", "top", "getMin"]

Output: [null,null,null,null,0,null,2,1]

Explanation:
MinStack minStack = new MinStack();
minStack.push(1);
minStack.push(2);
minStack.push(0);
minStack.getMin(); // return 0
minStack.pop();
minStack.top();    // return 2
minStack.getMin(); // return 1
Constraints:

-2^31 <= val <= 2^31 - 1.
pop, top and getMin will always be called on non-empty stacks.

*/

type MinStack struct {
	mins  []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		// Stack that only tracks mins
		// pop if poping same value on main stack
		mins: make([]int, 0),
		// main stack
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	// if mins is empty ,r o
	if len(this.mins) == 0 || this.mins[len(this.mins)-1] >= val {
		this.mins = append(this.mins, val)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	n := len(this.stack)
	if n == 0 {
		return
	}
	top := this.Top()
	minsTopIndex := len(this.mins) - 1
	if this.mins[minsTopIndex] == top {
		// pop
		this.mins = this.mins[:minsTopIndex]
	}
	// pop
	this.stack = this.stack[:n-1]

}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.mins) == 0 {
		return 0
	}
	return this.mins[len(this.mins)-1]
}
