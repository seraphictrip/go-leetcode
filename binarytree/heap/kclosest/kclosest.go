package kclosest

import (
	"container/heap"
	"math"
	"slices"
)

/*
Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane
 and an integer k, return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).

You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).



Example 1:


Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].
Example 2:

Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Explanation: The answer [[-2,4],[3,3]] would also be accepted.


Constraints:

1 <= k <= points.length <= 104
-104 <= xi, yi <= 104
*/

// Euclidean distance (i.e., √(x1 - x2)^2 + (y1 - y2)^2).
func EuclideanDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func KClosest(points [][]int, k int) [][]int {
	slices.SortFunc(points, func(a, b []int) int {
		// a - b (asc by euclidean distance)
		ad := EuclideanDistance(float64(a[0]), float64(a[1]), 0, 0)
		bd := EuclideanDistance(float64(b[0]), float64(b[1]), 0, 0)
		cmp := ad - bd
		if cmp < 0 {
			return -1
		}
		if cmp > 0 {
			return 1
		}
		return 0
	})
	return points[0:k]
}

type CoordinateHeap [][]int

func (h CoordinateHeap) Len() int { return len(h) }
func (h CoordinateHeap) Less(i, j int) bool {
	x := h[i]
	y := h[j]
	a := EuclideanDistance(float64(x[0]), float64(x[1]), 0, 0)
	b := EuclideanDistance(float64(y[0]), float64(y[1]), 0, 0)
	return a < b
}
func (h CoordinateHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *CoordinateHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *CoordinateHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func KClosestHeap(points [][]int, k int) [][]int {
	h := CoordinateHeap(points)
	// heapify O(n)
	heap.Init(&h)

	result := make([][]int, k)
	for i := range result {
		result[i] = heap.Pop(&h).([]int)
	}
	return result
}
