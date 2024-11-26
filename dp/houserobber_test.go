package dp_test

import (
	"dsa/dp"
	"strconv"
	"testing"
)

var robTests = []struct {
	nums     []int
	expected int
}{
	{[]int{1, 2, 3, 1}, 4},
	{[]int{2, 7, 9, 3, 1}, 12},
	{[]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}, 4173},
}

func TestRob(t *testing.T) {
	for i, e := range robTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := dp.RobBottomUp(e.nums)
			if actual != e.expected {
				t.Fatalf("Rob(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}

var Rob3Tests = []struct {
	tree     *dp.TreeNode
	expected int
}{
	{
		buildTree([]int{3, 2, 3, -1, 3, -1, 1}),
		7,
	},
	{
		buildTree([]int{3, 4, 5, 1, 3, -1, 1}),
		9,
	},
	{
		buildTree([]int{79, 99, 77, -1, -1, -1, 69, -1, 60, 53, -1, 73, 11, -1, -1, -1, 62, 27, 62, -1, -1, 98, 50, -1, -1, 90, 48, 82, -1, -1, -1, 55, 64, -1, -1, 73, 56, 6, 47, -1, 93, -1, -1, 75, 44, 30, 82, -1, -1, -1, -1, -1, -1, 57, 36, 89, 42, -1, -1, 76, 10, -1, -1, -1, -1, -1, 32, 4, 18, -1, -1, 1, 7, -1, -1, 42, 64, -1, -1, 39, 76, -1, -1, 6, -1, 66, 8, 96, 91, 38, 38, -1, -1, -1, -1, 74, 42, -1, -1, -1, 10, 40, 5, -1, -1, -1, -1, 28, 8, 24, 47, -1, -1, -1, 17, 36, 50, 19, 63, 33, 89, -1, -1, -1, -1, -1, -1, -1, -1, 94, 72, -1, -1, 79, 25, -1, -1, 51, -1, 70, 84, 43, -1, 64, 35, -1, -1, -1, -1, 40, 78, -1, -1, 35, 42, 98, 96, -1, -1, 82, 26, -1, -1, -1, -1, 48, 91, -1, -1, 35, 93, 86, 42, -1, -1, -1, -1, 0, 61, -1, -1, 67, -1, 53, 48, -1, -1, 82, 30, -1, 97, -1, -1, -1, 1, -1, -1}),
		176,
	},
	{
		//			(2)
		//		(1)		(3)
		//			(4)
		buildTree([]int{2, 1, 3, -1, 4}),
		6,
	},
}

func TestRob3(t *testing.T) {
	for i, e := range Rob3Tests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := dp.RobBFS(e.tree)
			if actual != e.expected {
				t.Fatalf("Rob3(%v) = %v, want %v", e.tree, actual, e.expected)
			}
		})
	}
}

// all inputs should be positive, so use -1 to indicate missing node
func buildTree(input []int) *dp.TreeNode {
	n := len(input)
	nodes := make([]*dp.TreeNode, len(input))
	for i := range input {
		if input[i] >= 0 {
			nodes[i] = &dp.TreeNode{Val: input[i]}
		}
	}
	root := nodes[0]
	for i := range nodes {
		if nodes[i] == nil {
			// we are nil, we cant attach any children
			continue
		}
		l := i*2 + 1
		r := i*2 + 2
		if l < n {
			nodes[i].Left = nodes[l]
		}
		if r < n {
			nodes[i].Right = nodes[r]
		}
	}
	return root
}
