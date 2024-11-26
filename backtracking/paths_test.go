package backtracking_test

import (
	bt "dsa/backtracking"
	"fmt"
	"strconv"
	"testing"
)

var (
	t1 = makeTestTree([]*bt.Node{bt.NewNode(4), bt.NewNode(0), bt.NewNode(1), nil, bt.NewNode(7), bt.NewNode(2), bt.NewNode(0)})
	t2 = makeTestTree([]*bt.Node{bt.NewNode(4),
		bt.NewNode(0), bt.NewNode(1),
		nil, bt.NewNode(7), bt.NewNode(0), bt.NewNode(2)})
	t3 = makeTestTree([]*bt.Node{bt.NewNode(4),
		bt.NewNode(0), bt.NewNode(1),
		nil, bt.NewNode(7), bt.NewNode(1), bt.NewNode(2),
		nil, nil, nil, nil, bt.NewNode(0),
	})
)

var HasValidPathTests = []struct {
	tree     *bt.Node
	expected bool
}{
	{},
	{t1, true},
	{t2, true},
	{t3, true},
}

func TestHasValidPath(t *testing.T) {
	for i, e := range HasValidPathTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := bt.HasValidPath(e.tree)
			if actual != e.expected {
				t.Fatalf("HasValidPath(%v) = %v, want %v", e.tree, actual, e.expected)
			}
			path := bt.Intstack([]int{})
			bt.GetPath(e.tree, &path)
			fmt.Println(path)
		})
	}
}

// TEST TREE
//			(4)
//		(0)		(1)
//		  (7)  (2)  (0)

func makeTestTree(nodes []*bt.Node) *bt.Node {
	n := len(nodes)
	if n == 0 {
		return nil
	}
	for i := 0; i < n; i++ {
		if nodes[i] == nil {
			continue
		}
		l, r := i*2+1, i*2+2
		if l < n {
			nodes[i].Left = nodes[l]
		}
		if r < n {
			nodes[i].Right = nodes[r]
		}
	}

	return nodes[0]

}
