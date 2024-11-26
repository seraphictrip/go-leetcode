package bfs

import (
	bt "dsa/binarytree"
)

// Level order traversal
// we put nodes in a queue for processing
// upon reaching the front of the line
// we visit the node and have children enter the queue
func BFS[T any](root *bt.TreeNode[T], visit func(*bt.TreeNode[T])) {
	if root == nil {
		return
	}
	q := NewQueue[*bt.TreeNode[T]](0)
	q.Enqueue(root)
	for !q.IsEmpty() {
		cur := q.Dequeue()
		// it doesn't really matter if we load our children first
		// or visit first, as I will be visited before next "customer" is processed
		visit(cur)
		if cur.Left != nil {
			q.Enqueue(cur.Left)
		}
		if cur.Right != nil {
			q.Enqueue(cur.Right)
		}
	}

}

func BFSWithLevel[T any](root *bt.TreeNode[T], visit func(node *bt.TreeNode[T], level int)) int {
	if root == nil {
		return 0
	}
	q := NewQueue[*bt.TreeNode[T]](0)
	q.Enqueue(root)

	level := 0
	for !q.IsEmpty() {
		// to keep track of level we have to process
		// all of the currently queued items, and put all direct children in queue for next pass
		n := q.Size()
		for n > 0 {
			cur := q.Dequeue()
			visit(cur, level+1)
			if cur.Left != nil {
				q.Enqueue(cur.Left)
			}
			if cur.Right != nil {
				q.Enqueue(cur.Right)
			}
			n--
		}
		level++

	}

	return level
}

// Wrapper around a simple "dynamic array" (via slice) queue
type Queue[T any] struct {
	queue []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.queue = append(q.queue, item)
}

func (q *Queue[T]) Dequeue() T {
	// panic if not there, not our problem
	item := q.queue[0]
	if q.Size() == 1 {
		q.queue = nil
	} else {
		q.queue = q.queue[1:]
	}
	return item
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.queue)
}

func NewQueue[T any](initialCap int) *Queue[T] {
	return &Queue[T]{
		queue: make([]T, 0, initialCap),
	}
}
