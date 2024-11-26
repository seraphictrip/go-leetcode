package topk

import "container/heap"

/*
347. Top K Frequent Elements
* Top K so immediately thinking heap, most frequent so max-heap
* frequency so thinking frequnecy map

https://leetcode.com/problems/top-k-frequent-elements/description/


Given an integer array nums and an integer k, return the k most frequent elements.
 You may return the answer in any order.

 I could also see accumulating into map and then extracting

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
freq = {1:3, 2: 2, 3: 1}
take k max
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
freq = {1:1}
take k max
Output: [1]


Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.


Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/

func TopKFrequent(nums []int, k int) []int {
	f := freq(nums)
	return takeKMax(f, k)
}

func takeKMax(freqs map[int]int, k int) []int {
	maxheap := MaxHeap(make([]CountOf[int], 0, len(freqs)))
	for of, count := range freqs {
		c := CountOf[int]{count, of}
		maxheap = append(maxheap, c)
	}
	heap.Init(&maxheap)
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(&maxheap).(CountOf[int]).Of
	}
	return result
}

func freq(nums []int) map[int]int {
	freqs := make(map[int]int, 0)
	for _, num := range nums {
		freqs[num]++
	}
	return freqs
}

type CountOf[T comparable] struct {
	Count int
	Of    T
}

type MaxHeap []CountOf[int]

// Len is the number of elements in the collection.
func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i int, j int) bool {
	// max heap, so >
	return h[i].Count > h[j].Count
}

// Swap swaps the elements with indexes i and j.
func (h MaxHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

// This is just stack push, call heap.Push(h, val) to push on stack
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(CountOf[int]))
}

func (h *MaxHeap) Pop() any {
	old := *h
	val := old[h.Len()-1]
	*h = old[:h.Len()-1]
	return val
}
