package myheap

import "slices"

func MedianSlidingWindow(nums []int, k int) []float64 {
	// small := NewMaxMyHeap()
	// large := NewMinMyHeap()
	return nil
}

// [0, 1, 2, 3]
func median(nums []int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		return (float64(nums[mid-1]) + float64(nums[mid])) / 2
	}
	return float64(nums[mid])
}

type MyHeap struct {
	data  []int
	isMax bool
}

func NewMinMyHeap() MyHeap {
	return MyHeap{
		data: make([]int, 0, 10),
	}
}

func NewMaxMyHeap() MyHeap {
	return MyHeap{
		data:  make([]int, 0, 10),
		isMax: true,
	}
}

// Len is the number of elements in the collection.
func (h MyHeap) Len() int {
	return len(h.data)
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
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (h MyHeap) Less(i int, j int) bool {
	if h.isMax {
		return h.data[i] > h.data[j]
	}
	return h.data[i] < h.data[j]
}

// Swap swaps the elements with indexes i and j.
func (h MyHeap) Swap(i int, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MyHeap) Push(x any) {
	h.data = append(h.data, x.(int))
}

func (h *MyHeap) Pop() any {
	top := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return top
}

func (h MyHeap) FindIndex(val int) int {
	return slices.Index(h.data, val)
}
