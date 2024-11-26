package priorityqueue

import "container/heap"

type Item[T any] struct {
	// value of the item
	value T
	// priority of item in queue
	priority int
	// index is needed by update and is maintained by the heap.Interface
	index int
}

// A PriorityQueue implements heap.Interface and holds Items
type PriorityQueue[T any] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// we want pop to give higest priority, essentially max heap
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

// Pop is not heap pop, it is removing last element
// think stack pop
func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // allow GC to reclaims
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Update(item *Item[T], val T, priority int) {
	item.value = val
	item.priority = priority
	heap.Fix(pq, item.index)
}
