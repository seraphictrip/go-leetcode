package bfs_test

import (
	bt "dsa/binarytree"
	"dsa/binarytree/bfs"
	"fmt"
	"slices"
	"strconv"
	"testing"
)

var BFSTests = []struct {
	bst      *bt.TreeNode[int]
	expected []int
	height   int
}{
	{},
	//				(1)
	{bt.MakeBSTFromInorder([]int{1}), []int{1}, 1},
	//				(2)
	//			(1)
	{bt.MakeBSTFromInorder([]int{1, 2}), []int{2, 1}, 2},
	//				(2)
	//			(1)		(3)
	{bt.MakeBSTFromInorder([]int{1, 2, 3}), []int{2, 1, 3}, 2},
	//				(3)
	//			(2)		(4)
	//		(1)
	{bt.MakeBSTFromInorder([]int{1, 2, 3, 4}), []int{3, 2, 4, 1}, 3},
	// 				(3)
	//			(2)		(5)
	//		  (1)	  (4)
	{bt.MakeBSTFromInorder([]int{1, 2, 3, 4, 5}), []int{3, 2, 5, 1, 4}, 3},
	//				(4)
	//		(2)				(6)
	//	 (1)   (3)		 (5)
	{bt.MakeBSTFromInorder([]int{1, 2, 3, 4, 5, 6}), []int{4, 2, 6, 1, 3, 5}, 3},
	//				  (4)
	//			(2)			(6)
	//		(1)   (3)	(5)		(7)
	{bt.MakeBSTFromInorder([]int{1, 2, 3, 4, 5, 6, 7}), []int{4, 2, 6, 1, 3, 5, 7}, 3},
	//				  (5)
	//			(3)			(7)
	//		(2)    (4)	(6)		(8)
	//	  (1)
	{bt.MakeBSTFromInorder([]int{1, 2, 3, 4, 5, 6, 7, 8}), []int{5, 3, 7, 2, 4, 6, 8, 1}, 4},
}

func TestBFS(t *testing.T) {
	for i, e := range BFSTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			acc := make([]int, 0, len(e.expected))
			bfs.BFS(e.bst, func(node *bt.TreeNode[int]) {
				acc = append(acc, node.Val)
			})
			if !slices.Equal(acc, e.expected) {
				t.Fatalf("levelorder got %v, want %v", acc, e.expected)
			}
		})
	}
}

func TestBFSWithLevel(t *testing.T) {
	for i, e := range BFSTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			acc := make([]int, 0, len(e.expected))
			height := bfs.BFSWithLevel(e.bst, func(node *bt.TreeNode[int], level int) {
				fmt.Println(level)
				acc = append(acc, node.Val)
			})
			if !slices.Equal(acc, e.expected) {
				t.Fatalf("levelorder got %v, want %v", acc, e.expected)
			}
			if height != e.height {
				t.Fatalf("unexpected height: %v != %v", height, e.height)
			}
		})
	}
}
