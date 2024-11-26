package topological

import (
	"container/list"
	"errors"
	"slices"
)

var (
	ErrCycle = errors.New("cycle")
)

// topological sort or topological ordering of a directed graph is a linear ordering of its
// vertices such that for every directed edge (u,v) from vertex u to vertex v, u comes before v
//
//	in the ordering
func Sort(adj [][]int, V int) ([]int, error) {
	// slice to store indegree of each vertex
	// indegree is how many edges pointing to me
	indegree := make([]int, V)
	for _, list := range adj {
		// fill indegress using neighbors
		for _, vertex := range list {
			indegree[vertex]++
		}
	}
	// Queue to store vertices with indegree 0
	q := NewQueue[int]()
	for i := 0; i < V; i++ {
		// load all indegrees of zero to be processed
		// indegrees on remaining items will be
		// decremented as these nodes get processed
		if indegree[i] == 0 {
			q.Enqueue(i)
		}
	}

	// linear ordering of vertices such that for every directed edge (u,v)
	// from vertex u to vertex v, u comes before v
	result := make([]int, 0)

	for !q.IsEmpty() {
		node := q.Dequeue()
		// add item to linear ordering
		result = append(result, node)
		// dcrement indegrees, so we can
		// add new vertices to queue
		for _, adjacent := range adj[node] {
			indegree[adjacent]--
			if indegree[adjacent] == 0 {
				q.Enqueue(adjacent)
			}
		}
	}

	// Cycle Detection
	if len(result) != V {
		return nil, ErrCycle
	}
	return result, nil

}

// topological sort or topological ordering of a directed graph is a linear ordering of its
// vertices such that for every directed edge (u,v) from vertex u to vertex v, u comes before v
//
//	in the ordering
func SortGeneric[T comparable](graph map[T][]T, V int) ([]T, error) {
	// track indegrees, which are edges to me, think dependencies
	indegree := make(map[T]int, V)
	// fill out indegree using adjacency list representation of graph
	for node, neighbors := range graph {
		// 0s are important, so make sure each node is represented
		// in indegrees even if no edges poim to it
		if _, ok := indegree[node]; !ok {
			indegree[node] = 0
		}
		// capture indegrees
		for _, n := range neighbors {
			indegree[n]++
		}
	}
	q := NewQueue[T]()
	// enqueue everything with 0 indegrees
	for key, val := range indegree {
		if val == 0 {
			q.Enqueue(key)
		}
	}

	result := make([]T, 0, V)

	for !q.IsEmpty() {
		// load and process
		node := q.Dequeue()
		result = append(result, node)

		// use adj list to decrement in degree of nodes pointed to  by node we processes
		for _, adjacent := range graph[node] {
			indegree[adjacent]--
			if indegree[adjacent] == 0 {
				q.Enqueue(adjacent)
			}
		}

	}

	// Cycle Detection
	if len(result) != V {
		return nil, ErrCycle
	}
	return result, nil
}

// Kahn's algorithm/BFS
func TopologicalSort[T comparable](graph map[T][]T, V int) ([]T, error) {
	// track indegrees
	indegrees := make(map[T]int, V)
	for node, neighbors := range graph {
		if _, ok := indegrees[node]; !ok {
			// tracking 0s are important, so make sure we get current vertex
			// if not yet in map
			indegrees[node] = 0
		}
		for _, n := range neighbors {
			indegrees[n]++
		}
	}
	// queue up 0-ingrees
	q := NewQueue[T]()
	for key, val := range indegrees {
		if val == 0 {
			q.Enqueue(key)
		}
	}

	result := make([]T, 0, V)
	// process queue
	// 1. put node in result
	// 2. decrement indegree of adjacent nodes
	// 2a. queue if needed
	for !q.IsEmpty() {
		cur := q.Dequeue()
		result = append(result, cur)
		for _, adjacent := range graph[cur] {
			indegrees[adjacent]--
			if indegrees[adjacent] == 0 {
				q.Enqueue(adjacent)
			}
		}
	}

	// detect cycles
	if len(result) != V {
		return nil, ErrCycle
	}

	return result, nil

}

func TopoSortDFS[T comparable](graph map[T][]T, V int) []T {
	visited := map[T]bool{}
	stack := make([]T, 0, V)

	var dfs func(node T)
	dfs = func(node T) {
		visited[node] = true
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
		stack = append(stack, node)
	}

	for node := range graph {
		if !visited[node] {
			dfs(node)
		}
	}

	slices.Reverse(stack)
	return stack
}

type QueueInterface[T any] interface {
	Enqueue(T)
	Dequeue() T
	Len() int
	IsEmpty() bool
}

type Queue[T any] struct {
	l *list.List
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		l: list.New(),
	}
}

func (q Queue[T]) Enqueue(item T) {
	q.l.PushBack(item)
}

func (q Queue[T]) Dequeue() T {
	front := q.l.Front()
	val := q.l.Remove(front).(T)
	return val
}

func (q Queue[T]) Len() int {
	return q.l.Len()
}

func (q Queue[T]) IsEmpty() bool {
	return q.l.Len() == 0
}
