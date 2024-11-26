package links

import "fmt"

type link[T any] struct {
	Val  T
	next *link[T]
}

type Queue[T any] struct {
	head *link[T]
	tail *link[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue Appends item to the end of the list
func (q *Queue[T]) Enqueue(item T) {
	node := &link[T]{Val: item}
	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.next = node
	q.tail = node

}

// Dequeue returns value from head of the list or 0 value
// with err would probably be better implementation in real world
func (q *Queue[T]) Dequeue() (out T) {
	head := q.head
	if head == nil {
		// TODO: should probaby have default value as part of queue definition
		return out
	}
	if head == q.tail {
		q.tail = nil
		q.head = nil
		return head.Val
	}
	q.head = head.next
	return head.Val
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}

func Print[T any](q *Queue[T]) {
	for cur := q.head; cur != nil; cur = cur.next {
		fmt.Printf("(%v)->", cur.Val)
	}
	fmt.Println()
}
