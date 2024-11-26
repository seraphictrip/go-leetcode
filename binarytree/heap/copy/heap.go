package copy

import (
	"sort"
)

/**
This code will start out as a line by line copy of container/heap, so I can better understand it
where it goes from there, nowhere? somewhere? is yet to be known.
*/

// The Interface type describes the requirements
// for a type using the routines in this package.
// Any type that implements it may be used as a
// min-heap with the following invariants (established after
// [Init] has been called or if the data is empty or sorted):
//
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//
// Note that [Push] and [Pop] in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use [heap.Push] and [heap.Pop].
type Heap[T any] interface {
	sort.Interface
	Push(x T)
	Pop() T
}

func Init[T any](h Heap[T]) {
	// heapify

}
