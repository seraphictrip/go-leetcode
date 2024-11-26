package adjacencylist

import "container/list"

type GraphNode[T any] struct {
	val       T
	neighbors []*GraphNode[T]
}

type Adjlist[T comparable] map[T][]T

// Given a list of directed edges, build an adjacency list
func MakeAdjacencyList[T comparable](edges [][]T) Adjlist[T] {
	adj := make(map[T][]T)
	for _, edge := range edges {
		src := edge[0]
		dst := edge[1]
		if _, ok := adj[src]; !ok {
			adj[src] = make([]T, 0)
		}
		if _, ok := adj[dst]; !ok {
			adj[dst] = make([]T, 0)
		}
		adj[src] = append(adj[src], dst)
	}
	return adj
}

// DFS (backtracking)
func CountPaths[T comparable](node T, target T, adjlist Adjlist[T], visited map[T]bool) int {
	if visited[node] {
		return 0
	}
	if node == target {
		return 1
	}
	count := 0
	visited[node] = true
	for _, neighbor := range adjlist[node] {
		count += CountPaths(neighbor, target, adjlist, visited)
	}
	visited[node] = false
	return count
}

func CountPathsBFS[T comparable](graph Adjlist[T], src, target T) int {
	acc := 0
	q := NewQueue[T]()
	q.Enqueue(src)

	for !q.IsEmpty() {
		n := q.Len()
		for n > 0 {
			cur := q.Dequeue()
			if cur == target {
				acc++
			}
			for _, neighbor := range graph[cur] {
				q.Enqueue(neighbor)
			}
			n--
		}
	}
	return acc
}

func GetPaths[T comparable](graph Adjlist[T], src, target T) [][]T {
	acc := make([][]T, 0)
	visited := make(map[T]bool)

	var dfs func(path []T, node T)
	dfs = func(path []T, node T) {
		// if already visited bail
		if visited[node] {
			return
		}
		// update path
		path = append(path, node)
		// check if at target
		if node == target {
			// add path to accumulator
			completedPath := make([]T, len(path))
			copy(completedPath, path)
			acc = append(acc, completedPath)
		}
		// mark visited
		visited[node] = true
		// visit all neighbors
		for _, neighbor := range graph[node] {
			dfs(path, neighbor)
		}
		visited[node] = false

	}

	dfs([]T{}, src)
	return acc
}

func ShortestPath[T comparable](graph Adjlist[T], src, target T) int {
	acc := 0
	q := NewQueue[T]()
	q.Enqueue(src)
	visited := make(map[T]bool)
	visited[src] = true

	for !q.IsEmpty() {
		n := q.Len()
		for n > 0 {
			cur := q.Dequeue()
			if cur == target {
				return acc
			}
			for _, neighbor := range graph[cur] {
				if !visited[neighbor] {
					q.Enqueue(neighbor)
					visited[neighbor] = true
				}

			}
			n--
		}
		acc++
	}

	return -1
}

// dfs template
/*
func Caller(graph, ...other) any {
	acc := any

	var dfs func (node ...other)
	dfs = func(node, ...other){
		// do whatever is needed on graph
		// accumulate to acc as needed
	}

	dfs(root)
	return acc
}
*/

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
