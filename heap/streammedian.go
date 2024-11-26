package myheap

import "container/heap"

type MedianFinder struct {
	// small should maintain invariant all items in small are less than all items in
	// large
	small IntHeap
	large IntHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		small: MaxHeap(),
		large: MinHeap(),
	}
}

func (m *MedianFinder) AddNum(num int) {
	// push
	heap.Push(&m.small, num)
	// maintain invariant
	if m.large.Len() != 0 && m.small.Peek() >= m.large.Peek() {
		val := heap.Pop(&m.small).(int)
		heap.Push(&m.large, val)
	}

	// maintain balance
	if m.small.Len() > m.large.Len()+1 {
		val := heap.Pop(&m.small).(int)
		heap.Push(&m.large, val)
	}
	if m.large.Len() > m.small.Len()+1 {
		val := heap.Pop(&m.large).(int)
		heap.Push(&m.small, val)
	}

}

func (m *MedianFinder) FindMedian() float64 {
	// if odd, take from longest
	if m.small.Len() > m.large.Len() {
		return float64(m.small.Peek())
	}
	if m.large.Len() > m.small.Len() {
		return float64(m.large.Peek())
	}
	// else take avg
	return (float64(m.small.Peek()) + float64(m.large.Peek())) / 2
}

type IntHeap struct {
	heap  []int
	isMax bool
}

func MaxHeap() IntHeap {
	return IntHeap{
		heap:  make([]int, 0, 10),
		isMax: true,
	}
}

func MinHeap() IntHeap {
	return IntHeap{
		heap: make([]int, 0, 10),
	}
}

// Len is the number of elements in the collection.
func (h IntHeap) Len() int {
	return len(h.heap)
}

func (h IntHeap) Less(i int, j int) bool {
	if h.isMax {
		return h.heap[i] > h.heap[j]
	}
	return h.heap[i] < h.heap[j]
}

// Swap swaps the elements with indexes i and j.
func (h IntHeap) Swap(i int, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *IntHeap) Push(x any) {
	// this is stack push, not heap push
	h.heap = append(h.heap, x.(int))
}

func (h *IntHeap) Pop() any {
	// this is stack pop, not heap pop
	top := h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	return top
}

func (h IntHeap) Peek() int {
	return h.heap[0]
}
