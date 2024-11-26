package directed

import "container/heap"

// Given a connected graph respresented by a list of edges where
// edge[0] == src, edge[1] = dst, and edge[2] = weight,
// find teh shortest path from src to every otehr node in the graph.
// There are n nodes in the graph
// O(E*logV) or O(E*LogE)
func ShortestPath(edges [][]int, n, src int) map[int]int {
	adj := make(map[int][][2]int, n)

	for _, edge := range edges {
		src, dest, weight := edge[0], edge[1], edge[2]
		adj[src] = append(adj[src], [2]int{weight, dest})
	}

	shortest := make(map[int]int, n)
	minHeap := NewMinHeap(n * n)

	heap.Push(&minHeap, [2]int{0, src})

	for len(minHeap) != 0 {
		weighted := heap.Pop(&minHeap).([2]int)
		w, node := weighted[0], weighted[1]
		if _, ok := shortest[node]; ok {
			continue
		}
		shortest[node] = w

		for _, weighted2 := range adj[node] {
			w2, node2 := weighted2[0], weighted2[1]
			if _, ok := shortest[node2]; !ok {
				heap.Push(&minHeap, [2]int{w + w2, node2})
			}
		}

	}

	return shortest
}

type minheap [][2]int

func NewMinHeap(cap int) minheap {
	return make([][2]int, cap)
}

// Len is the number of elements in the collection.
func (h minheap) Len() int {
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
func (h minheap) Less(i int, j int) bool {
	return h[i][0] < h[j][0]
}

// Swap swaps the elements with indexes i and j.
func (h minheap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *minheap) Pop() any {
	old := *h
	topIndex := len(old) - 1
	top := old[topIndex]
	*h = old[:topIndex]
	return top
}
