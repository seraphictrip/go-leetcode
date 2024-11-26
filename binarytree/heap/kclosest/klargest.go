package kclosest

import (
	"container/heap"
	"fmt"
	"slices"
)

/*
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?



Example 1:

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Example 2:

Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4


Constraints:

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/

func FindKthLargestNatural(nums []int, k int) int {
	slices.Sort(nums)
	// 0-based
	return nums[len(nums)-k]
}

func FindKthLargestRecursive(nums []int, k int) int {
	// find max index
	index := slices.Index(nums, slices.Max(nums))
	// if k == 1 return nums[maxIndex] else call with k - 1
	if k == 1 {
		return nums[index]
	}
	result := append(nums[0:index], nums[index+1:]...)
	return FindKthLargestRecursive(result, k-1)
}

func FindKthLargest(nums []int, k int) int {
	h := MaxIntHeap(nums)
	// heapify O(n)
	heap.Init(&h)

	var res int
	for k > 0 {
		res = heap.Pop(&h).(int)
		k--
	}
	return res
}

type MaxIntHeap []int

// Len is the number of elements in the collection.
func (h MaxIntHeap) Len() int {
	return len(h)
}

// Less reports whether the element with index i
// must sort before the element with index j.
func (h MaxIntHeap) Less(i int, j int) bool {
	// want a max heap, so greater than
	return h[i] > h[j]
}

// Swap swaps the elements with indexes i and j.
func (h MaxIntHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxIntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() any {
	old := *h
	n := len(old)
	top := old[n-1]
	*h = old[0 : n-1]
	return top

}

func FindKthLargestQuickSelect(nums []int, k int) int {
	fmt.Printf("FindKthLargestQuickSelect(%v, %v)\n", nums, k)
	// partition as if we doing quicksort partition
	n := len(nums)
	// ki is the expected index of k as kth largest
	// assuming asc order, that is k back from end of list...
	ki := n - k
	return quickselect(nums, 0, n-1, ki)

}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// quickselect the kth element
func quickselect(nums []int, l, r, k int) int {
	pivotValue := nums[r]
	pi := l

	for i := l; i < r; i++ {
		if nums[i] < pivotValue {
			swap(nums, pi, i)
			pi++
		}
	}
	// swap partition value into place
	swap(nums, pi, r)

	if k == pi {
		return nums[pi]
	}
	if k < pi {
		return quickselect(nums, l, pi-1, k)
	} else {
		// k > pi
		return quickselect(nums, pi+1, r, k)
	}
}
