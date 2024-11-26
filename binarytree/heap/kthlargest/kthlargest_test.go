package kthlargest_test

import (
	goheap "container/heap"
	"dsa/binarytree/heap"
	"fmt"
	"slices"
	"strconv"
	"testing"
)

/*
You are part of a university admissions office and need to keep track of the kth highest test score from applicants in real-time.
This helps to determine cut-off marks for interviews and admissions dynamically as
new applicants submit their scores.

You are tasked to implement a class which, for a given integer k, maintains a stream of test
scores and continuously returns the kth highest test score after a new score has been
submitted. More specifically, we are looking for the kth highest score in the sorted list
 of all scores.

Implement the KthLargest class:

KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of
 test scores nums.
int add(int val) Adds a new test score val to the stream and returns the element
representing the kth largest element in the pool of test scores so far.
*/

// Attempt 1: just sort as add
type KthLargestRef struct {
	k      int
	scores []int
}

func Constructor(k int, nums []int) KthLargestRef {
	slices.SortFunc(nums, func(a, b int) int {
		return b - a
	})
	return KthLargestRef{
		k:      k,
		scores: nums[:k],
	}
}

func (kth *KthLargestRef) Add(val int) int {
	if len(kth.scores) == kth.k && val < kth.scores[len(kth.scores)-1] {
		// if we  already have k elems and this is not a candidate ignore
		return kth.scores[kth.k-1]
	}
	kth.scores = append(kth.scores, val)
	slices.SortFunc(kth.scores, func(a, b int) int {
		return b - a
	})
	// minor improvement, from O(nlogn) to O(klogk)
	// this is actually prob pretty good improvment for a "real" system
	kth.scores = kth.scores[0:kth.k]
	return kth.scores[kth.k-1]
}

type KthLargest struct {
	k      int
	scores *heap.IntHeap
}

func NewKthLargest(k int, initial []int) KthLargest {
	// create the initial heap
	h := heap.NewIntHeap(initial)
	// trim to k
	for h.Len() > k {
		goheap.Pop(h)
	}

	kth := KthLargest{
		k:      k,
		scores: h,
	}
	return kth
}

func (kth *KthLargest) String() string {
	return fmt.Sprintf("{%d, %v}", kth.k, *kth.scores)
}

func (kth *KthLargest) Add(val int) int {
	goheap.Push(kth.scores, val)
	for kth.scores.Len() > kth.k {
		goheap.Pop(kth.scores)
	}
	return kth.scores.Min()
}

// func percolateUp(heap []int, i int) {
// 	if heap[i] < heap[(i-1)/2] {
// 		swap(heap, i, (i-1)/2)
// 		i = (i - 1) / 2
// 		percolateUp(heap, i)
// 	}
// }

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type Pair struct {
	signal int
	kth    int
}

var KthLargestTests = []struct {
	k           int
	initial     []int
	addKthPairs []Pair
}{
	{3, []int{4, 5, 8, 2}, []Pair{
		//[8,5,4]
		{3, 4},
		// [8,5,5]
		{5, 5},
		// [10, 8, 5]
		{10, 5},
		// [10, 9, 8]
		{9, 8},
		// [10, 9, 8]
		{4, 8},
	}},
	{4, []int{7, 7, 7, 7, 8, 3},
		[]Pair{
			// [8,7,7,7]
			{2, 7},
			// [10, 8, 7, 7]
			{10, 7},
			// [10, 9, 8, 7]
			{9, 7},
			// [10, 9, 9, 8]
			{9, 8},
		},
	},
}

func TestKthLargest(t *testing.T) {
	for i, e := range KthLargestTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			tracker := NewKthLargest(e.k, e.initial)

			for _, pair := range e.addKthPairs {
				actual := tracker.Add(pair.signal)
				if actual != pair.kth {
					t.Fatalf("(%vth).Add(%v) = %v, want %v: %v", e.k, pair.signal, actual, pair.kth, tracker.scores)
				}
			}
		})
	}
}
