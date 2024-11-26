package clonegraph

/*
Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}


Test case format:

For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.

An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.



Example 1:


Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
edges: [[1,2], [1,4],
		[2, 1],[2,3]
		[3,2], [3,4]
		[4,1], [4,3]
		]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
Example 2:


Input: adjList = [[]]
Output: [[]]
Explanation: Note that the input contains one empty list. The graph consists of only one node with val = 1 and it does not have any neighbors.
Example 3:

Input: adjList = []
Output: []
Explanation: This an empty graph, it does not have any nodes.


Constraints:

The number of nodes in the graph is in the range [0, 100].
1 <= Node.val <= 100
Node.val is unique for each node.
There are no repeated edges and no self-loops in the graph.
The Graph is connected and all nodes can be visited starting from the given node.
*/

type Node struct {
	Val       int
	Neighbors []*Node
}

func NewNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

func CloneGraph(node *Node) *Node {
	acc := make(map[*Node]*Node)
	var dfs func(n *Node)
	dfs = func(n *Node) {
		if acc[n] != nil {
			// already visited
			return
		}
		acc[n] = NewNode(n.Val)
		for _, neighbor := range n.Neighbors {
			dfs(neighbor)
			acc[n].Neighbors = append(acc[n].Neighbors, acc[neighbor])
		}
	}
	dfs(node)
	return acc[node]
}

func fmap[T any, V any](input []T, fn func(T) V) []V {
	acc := make([]V, len(input))
	for i := range input {
		acc[i] = fn(input[i])
	}
	return acc
}

// func BuildAdjacencyList(node *Node) [][]int {

// }

// func FromAdjacencyList(ls [][]int) *Node {
// 	m :
// 	for i := range ls {

// 	}
// }
/*
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
*/
