package undirected

import "container/heap"

func MinCostConnectPoints(points [][]int) int {
	n := len(points)
	adj := MakeAdjList(points)
	var minheap MinHeap = make([]WeightedEdge, 0, n*n)
	// arbitrary start at point labeled 0
	for _, directedEdge := range adj[0] {
		src, weight, dest := 0, directedEdge[0], directedEdge[1]
		heap.Push(&minheap, WeightedEdge{src, dest, weight})
	}

	visited := make([]bool, len(points))
	visited[0] = true
	// we will add cheapest edges, as we build minimal spanning tree
	result := 0
	for len(minheap) > 0 {
		weightedEdge := heap.Pop(&minheap).(WeightedEdge)
		cur := weightedEdge.dest
		if visited[cur] {
			continue
		}
		result += weightedEdge.weight
		visited[cur] = true
		for _, neighbor := range adj[cur] {
			if !visited[neighbor[1]] {
				heap.Push(&minheap, WeightedEdge{cur, neighbor[1], neighbor[0]})
			}
		}
	}
	return result
}

/*

adj = map[int][2]int{
    0: [2]int{{}, {}, {}, {}, {}},
    1: [2]int{},
    2: [2]int{},
    3: [2]int{},
    4: [2]int{},
}
*/

func MakeAdjList(points [][]int) map[int][][2]int {
	adj := make(map[int][][2]int, len(points))

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			// add edge to me (i), and edge other direction as undirected
			adj[i] = append(adj[i], [2]int{weight(points[i], points[j]), j})
			adj[j] = append(adj[j], [2]int{weight(points[i], points[j]), i})
		}
	}
	return adj
}

// Manhattan distance between two points
// |x1-x2| + |y1-y2|
func weight(a, b []int) int {
	ax, ay := a[0], a[1]
	bx, by := b[0], b[1]
	// swap as needed so don't need to do abs
	// mostly just avoiding casting to float
	if ax < bx {
		ax, bx = bx, ax
	}
	if ay < by {
		ay, by = by, ay
	}
	return ax - bx + ay - by
}

type WeightedEdge struct {
	src    int
	dest   int
	weight int
}

type MinHeap []WeightedEdge

// Len is the number of elements in the collection.
func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i int, j int) bool {
	return h[i].weight < h[j].weight
}

// Swap swaps the elements with indexes i and j.
func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

// STACK PUSH, see heap.Push for heap push
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(WeightedEdge))
}

// STACK POP: see heap.Pop for heap Pop
func (h *MinHeap) Pop() any {
	old := *h
	topIndex := len(old) - 1
	top := old[topIndex]
	*h = old[:topIndex]
	return top
}
