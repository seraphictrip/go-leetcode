package heap

import (
	"cmp"
	"container/heap"
	"fmt"
	"slices"
)

type MinHeap[T cmp.Ordered] struct {
	heap []T
}

func NewMinHeap[T cmp.Ordered]() *MinHeap[T] {
	return new(MinHeap[T])
}

func (h *MinHeap[T]) ToArray() []T {
	return slices.Clone(h.heap)
}

func (h *MinHeap[T]) Push(item T) {
	h.heap = append(h.heap, item)
	i := len(h.heap) - 1

	// percolate up
	for h.heap[i] < h.heap[(i-1)/2] {
		swap(h.heap, i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func (h *MinHeap[T]) Min() T {
	return h.heap[0]
}

func (h *MinHeap[T]) Pop() T {
	tmp := h.heap[0]
	n := len(h.heap)

	swap(h.heap, 0, len(h.heap)-1)

	// percolate down
	i := 0
	for {
		j1 := LeftChildIndex(i)
		// j1 < 0 on overflow
		if j1 >= n {
			// we have nor more down to go
			break
		}
		// set j to leftchild, we will check if we need to change it
		j := j1
		j2 := RightChildIndex(i)
		if j2 < n {
			// we have to consider right child, check for min value
			rightVal := h.heap[j2]
			if min(h.heap[j1], h.heap[j2]) == rightVal {
				j = j2
			}
		}

		if h.heap[i] > h.heap[j] {
			swap(h.heap, i, j)
		}
		i = j
	}
	h.heap = h.heap[:len(h.heap)-1]

	return tmp
}

func (h *MinHeap[T]) Len() int {
	return len(h.heap)
}

// in 0-based parent is at (i-1)//2
//
//	// to indicate int division
//
// for 1-based we could just do //2
func ParentIndex(i int) int {
	return (i - 1) / 2
}

// in 0-based leftchild is at 2i+1
// for 1-based we would just use 2i
func LeftChildIndex(i int) int {
	return 2*i + 1
}

// in 0-based leftchild is at 2i+2
// for 1-based we would just use 2i + 1
func RightChildIndex(i int) int {
	return 2*i + 2
}

// O(n^2), there is more optimal
func MinHeapifyNaive[T cmp.Ordered](inputs []T) *MinHeap[T] {
	h := NewMinHeap[T]()
	for i := range inputs {
		h.Push(inputs[i])
	}
	return h
}

// a more optimal version can just percolate down on left half?
// will need to explore and write out trees
func MinHeapify[T cmp.Ordered](inputs []T) *MinHeap[T] {
	h := &MinHeap[T]{heap: inputs}
	n := len(inputs)
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
	return h
}

func down[T cmp.Ordered](heap *MinHeap[T], i, n int) {

}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func min[T cmp.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// IntHeap is a wrapper around []int which fulfills the "container/heap" Interface
// mostly taken from example
//
//	type Interface interface {
//			sort.Interface (Len(), Less(i, j), Swap(i, j))
//			Push(x any) // add x as element Len() (think stackpush/append)
//			Pop() any	// remove and return element Len() -1 (think stack pop)
//		}
type IntHeap []int

// Len is the number of elements in the collection.
func (h IntHeap) Len() int {
	return len(h)
}

func NewIntHeap(initial []int) *IntHeap {
	// cast []initial to an IntHeap
	var h IntHeap = IntHeap(initial)
	// heapify
	heap.Init(&h)
	return &h
}

func (h *IntHeap) String() string {
	if h == nil {
		return "nil"
	}
	return fmt.Sprintf("*%v", *h)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
func (h IntHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements with indexes i and j.
func (h IntHeap) Swap(i int, j int) {
	swap(h, i, j)
}

// This is underlaying push, think stack push, not a a heap push
// for heap push use heap.Push(h, val)
// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop is a stack pop, not a heap pop
// we use heap.Pop(h) for a a heap pop
// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	min := old[n-1]
	*h = old[0 : n-1]
	return min
}

func (h IntHeap) Min() int {
	return h[0]
}
