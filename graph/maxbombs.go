package graph

import (
	"math"
)

func MaximumDetonation(bombs [][]int) int {
	graph := MakeGraph(bombs)
	result := 0

	var dfs func(int, int, []bool) int

	dfs = func(node, count int, visited []bool) int {
		if visited[node] {
			return count
		}
		visited[node] = true
		count++
		if count > result {
			result = count
		}
		for _, neighbor := range graph[node] {
			count = dfs(neighbor, count, visited)
		}
		// backtrack
		// visited[node] = false
		return count
	}
	// find max connectedness
	for i := range graph {
		visited := make([]bool, len(bombs))
		dfs(i, 0, visited)
	}
	return result
}

func MakeGraph(bombs [][]int) [][]int {
	graph := make([][]int, len(bombs))
	// for each point, compare range to distance
	// if range > distance connect
	for i, p1 := range bombs {
		x1, y1, rng := p1[0], p1[1], p1[2]
		edges := make([]int, 0, len(bombs))
		for j, p2 := range bombs {
			x2, y2 := p2[0], p2[1]
			if i == j {
				continue
			}
			d := distance(x1, x2, y1, y2)
			if float64(rng) >= d {
				edges = append(edges, j)
			}
		}
		graph[i] = edges
	}

	return graph
}

func distance(x1, x2, y1, y2 int) float64 {
	a := math.Abs(float64(x1) - float64(x2))
	b := math.Abs(float64(y1) - float64(y2))
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return c
}
