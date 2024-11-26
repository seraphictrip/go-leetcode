package laststone

import (
	"container/heap"
	"fmt"
	"slices"
)

/*
You are given an array of integers stones where stones[i] is the weight of the ith stone.

We are playing a game with the stones.
On each turn, we choose the heaviest two stones and smash them together.
Suppose the heaviest two stones have weights x and y with x <= y.
The result of this smash is:

If x == y, both stones are destroyed, and
If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no stones left, return 0.



Example 1:

Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
Example 2:

Input: stones = [1]
Output: 1


Constraints:

1 <= stones.length <= 30
1 <= stones[i] <= 1000
*/

func desc(a, b int) int {
	return a - b
}

func LastStoneWeightSearch(stones []int) int {
	slices.Sort(stones)
	for len(stones) > 1 {
		y := stones[len(stones)-1]
		x := stones[len(stones)-2]
		stones = stones[0 : len(stones)-2]
		fmt.Printf("%v vs %v: %v\n", y, x, y-x)
		if y-x > 0 {
			stones = append(stones, y-x)
			slices.Sort(stones)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

func LastStoneWeightRecursive(stones []int) int {
	n := len(stones)
	// base cases
	if n == 0 {
		return 0
	}
	if n == 1 {
		return stones[0]
	}
	// sort
	slices.Sort(stones)
	// choose stones
	yi := len(stones) - 1
	xi := len(stones) - 2
	y := stones[yi]
	x := stones[xi]

	stones = stones[:xi]
	if y-x > 0 {
		stones = append(stones, y-x)
	}
	return LastStoneWeightRecursive(stones)
}

// LastStoneWeight with max heap
func LastStoneWeight(stones []int) int {
	// heapify stones, parent GREATER THAN children
	h := MaxHeapify(stones)
	for h.Len() > 1 {
		// the heaviest two stones have weights x and y with x <= y.
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		if x > y {
			// with x <= y.
			x, y = y, x
		}
		// If x == y, both stones are destroyed
		if x == y {
			continue
		}
		// If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
		y -= x
		// push back onto the heap
		heap.Push(h, y)
	}
	if h.Len() == 1 {
		return heap.Pop(h).(int)
	}
	return 0

}

type MaxHeap []int

// Len is the number of elements in the collection.
func (h MaxHeap) Len() int {
	return len(h)
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
func (h MaxHeap) Less(i int, j int) bool {
	// NOTE: greater then, as max heap
	return h[i] > h[j]
}

// Swap swaps the elements with indexes i and j.
func (h MaxHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	lastIndex := len(old) - 1
	top := old[lastIndex]
	*h = old[:lastIndex]
	return top
}

func MaxHeapify(input []int) *MaxHeap {
	h := MaxHeap(input)
	heap.Init(&h)

	return &h
}
