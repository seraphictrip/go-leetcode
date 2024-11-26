package clonegraph_test

import (
	"dsa/adjacencylist/clonegraph"
	"fmt"
	"strconv"
	"testing"
)

var CloneGraphTests = []struct {
}{}

func TestCloneGraph(t *testing.T) {
	for i, e := range CloneGraphTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			fmt.Println(e)
		})
	}
}

/*
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
edges: [[1,2], [1,4],

	[2, 1],[2,3]
	[3,2], [3,4]
	[4,1], [4,3]
	]
*/
func TestClone(t *testing.T) {
	one := clonegraph.NewNode(1)
	two := clonegraph.NewNode(2)
	three := clonegraph.NewNode(3)
	four := clonegraph.NewNode(4)
	one.Neighbors = []*clonegraph.Node{two, four}
	two.Neighbors = []*clonegraph.Node{one, three}
	three.Neighbors = []*clonegraph.Node{two, four}
	four.Neighbors = []*clonegraph.Node{one, three}

	fmt.Println(clonegraph.CloneGraph(one))

}
