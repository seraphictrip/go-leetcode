package binarytree

import (
	"cmp"
	"fmt"
	"math"
	"math/rand"
	"slices"
	"strconv"
	"testing"
)

var HeightTests = []struct {
	node     *TreeNode[int]
	expected int
}{
	// Height of balanced trees should be approx logN, but we want to count from 1
	// Floor(log2(n)+1)
	// log2(0) = Inf... we just handle this case
	{nil, 0},
	// log2(1) = 0; Floor(0+1) = 0
	{MakeBinaryTree([]int{1}), 1},
	// log2(2) = 1; Floor(1+1) = 2
	{MakeBinaryTree([]int{1, 2}), 2},
	// log2(3) = 1.585; Floor(1.585+1) = 2
	{MakeBinaryTree([]int{1, 2, 3}), 2},
	// log2(4) = 2; Floor(2+1) = 3
	{MakeBinaryTree([]int{1, 2, 3, 4}), 3},
	// log2(5) = 2.322; Floor(2.322 + 1) = 3
	{MakeBinaryTree([]int{1, 2, 3, 4, 5}), 3},
	// log2(8) = 3; Floor(3+1) = 4
	{MakeBinaryTree([]int{1, 2, 3, 4, 5, 6, 7, 8}), 4},
	// log2(16) = 4; Floor(4+1) = 5
	{MakeBalancedBST(Range(1, 16, 1)), 5},
	// log2(32) = 5; Floor(5+1) = 6
	{MakeBalancedBST(Range(1, 32, 1)), 6},
	// log2(64) = 6; Floor(5+1) = 7
	{MakeBalancedBST(Range(1, 64, 1)), 7},
	// log2(100) = 6.644; Floor(6.644+1) = 7
	{MakeBalancedBST(Range(1, 100, 1)), 7},
	// log2(128) = 7; Floor(7+1) = 8
	{MakeBalancedBST(Range(1, 128, 1)), 8},
	// log2(256) = 8; Floor(8+1) = 9
	{MakeBalancedBST(Range(1, 256, 1)), 9},
	// log2(512) = 9; Floor(9+1) = 10
	{MakeBalancedBST(Range(1, 512, 1)), 10},
	// log2(1024) = 10; Floor(10+1) = 11
	{MakeBalancedBST(Range(1, 1024, 1)), 11},
	// log2(2048) = 11; Floor(11+1) = 12
	{MakeBalancedBST(Range(1, 2048, 1)), 12},
	// log2(4096) = 12; Floor(12+1) = 13
	{MakeBalancedBST(Range(1, 4096, 1)), 13},
	// log2(8192) = 13; Floor(13+1) = 14
	{ToBST(Range(1, 8192, 1)), 14},
	// log2(16384) = 14; Floor(14+1) = 15
	{MakeBalancedBST(Range(1, 16384, 1)), 15},
	// log2(32768) = 15; Floor(15+1) = 16
	{MakeBalancedBST(Range(1, 32768, 1)), 16},
	// log2(65536) = 16; Floor(16+1) = 17
	{MakeBalancedBST(Range(1, 65536, 1)), 17},
	// log2(16777216) = 24; Floor(24+1) = 25
	{MakeBalancedBST(Range(1, 16777216, 1)), 25},
}

func TestHeight(t *testing.T) {
	for i, e := range HeightTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := Height(e.node)
			count := len(Collect(e.node))
			balancedHeight := math.Floor(math.Log2(float64(count)) + 1)
			fmt.Printf("count: %v log2: %0.3f blancedheight: %v\n", count, math.Log2(float64(count)), balancedHeight)
			if actual != e.expected {
				t.Fatalf("Height(%v) = %v, want %v", e.node, actual, e.expected)
			}
		})
	}
}

var InorderTests = []struct {
	root     *TreeNode[int]
	expected []int
}{
	{},
	{ToBST([]int{1}), []int{1}},
	// inorder is just same order as array rep
	{MakeBinaryTree([]int{5, 4, 3, 1, 2}), []int{5, 4, 3, 1, 2}},
	{ToBST(Range(0, 100, 1)), Range(0, 100, 1)},
}

func TestInorder(t *testing.T) {
	for i, e := range InorderTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			acc := make([]int, 0, len(e.expected))
			rev := make([]int, 0, len(e.expected))
			InorderIterative[int](e.root, func(node *TreeNode[int]) {
				acc = append(acc, node.Val)
			})
			InorderRightToLeft[int](e.root, func(node *TreeNode[int]) {
				rev = append(rev, node.Val)
			})
			if !slices.Equal(acc, e.expected) {
				t.Fatalf("%v != %v", acc, e.expected)
			}
			slices.Reverse(e.expected)
			if !slices.Equal(rev, e.expected) {
				t.Fatalf("RtoL no bueno, got %v, want %v", rev, e.expected)
			}
		})
	}
}

var PreorderTests = []struct {
	root     *TreeNode[int]
	expected []int
}{
	{ToBST([]int{1}), []int{1}},
	{ToBST([]int{1, 2}), []int{2, 1}},
	//         (2)
	//       (1)  (3)
	{ToBST([]int{1, 2, 3}), []int{2, 1, 3}},
	// 			(3)
	//		  (2)  (5)
	//      (1)   (4)
	{ToBST([]int{1, 2, 3, 4, 5}), []int{3, 2, 1, 5, 4}},
	// 			(4)
	//		  (2)  	(6)
	//      (1) (3)(5)
	{ToBST([]int{1, 2, 3, 4, 5, 6}), []int{4, 2, 1, 3, 6, 5}},
	// 			(4)
	//         /   \
	//		(2)  	(6)
	//     /   \    /  \
	//    (1)   (3)(5)  (7)
	{ToBST([]int{1, 2, 3, 4, 5, 6, 7}), []int{4, 2, 1, 3, 6, 5, 7}},
	// 			(5)
	//		(3)  	(7)
	//    (2)  (4) (6)  (8)
	//  (1)
	{ToBST([]int{1, 2, 3, 4, 5, 6, 7, 8}), []int{5, 3, 2, 1, 4, 7, 6, 8}},
	//					(5)
	//			(3)				(8)
	//		(2)		(4)		(7)		(9)
	//	  (1)	 		   (6)
	//
	{ToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), []int{5, 3, 2, 1, 4, 8, 7, 6, 9}},
	{ToBST(Range(1, 10, 1)), []int{6, 3, 2, 1, 5, 4, 9, 8, 7, 10}},
	{ToBST(Range(1, 16, 1)), []int{9, 5, 3, 2, 1, 4, 7, 6, 8, 13, 11, 10, 12, 15, 14, 16}},
}

func TestPreorder(t *testing.T) {
	for i, e := range PreorderTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			acc := make([]int, 0, len(e.expected))
			Preorder[int](e.root, func(node *TreeNode[int]) {
				acc = append(acc, node.Val)
			})
			if !slices.Equal(acc, e.expected) {
				t.Fatalf("%v != %v", acc, e.expected)
			}
			acc = make([]int, 0, len(e.expected))
			PreorderIterative(e.root, func(node *TreeNode[int]) {
				acc = append(acc, node.Val)
			})
			if !slices.Equal(acc, e.expected) {
				t.Fatalf("%v != %v", acc, e.expected)
			}
		})
	}
}

// post order partitions [less][greater][root]
// this maps directly onto BST+PostORder
var PostOrderTests = []struct {
	root     *TreeNode[int]
	expected []int
}{
	{},
	// 			(1)
	{ToBST([]int{1}), []int{1}},
	//			(2)
	//		(1)
	{ToBST([]int{1, 2}), []int{1, 2}},
	//			(2)
	//		(1)		(3)
	{ToBST([]int{1, 2, 3}), []int{1, 3, 2}},
	//			(3)
	//		(2)		(5)
	//    (1)	  (4)
	{ToBST([]int{1, 2, 3, 4, 5}), []int{1, 2, 4, 5, 3}},
	//			(3)
	//		(1)			(5)
	//	(0)   (2)	  (4)  (6)
	{ToBST([]int{0, 1, 2, 3, 4, 5, 6}), []int{0, 2, 1, 4, 6, 5, 3}},
	// 			(3)
	//         /   \
	//		(1)  	(5)
	//     /   \    /  \
	//    (0)   (2)(4)  (6)
	{ToBST([]int{0, 1, 2, 3, 4, 5, 6}), []int{0, 2, 1, 4, 6, 5, 3}},
	{ToBST([]int{0, 1, 2, 3, 4, 5, 6, 7}), []int{0, 1, 3, 2, 5, 7, 6, 4}},
	{ToBST(Range(0, 16, 1)), []int{0, 1, 3, 2, 5, 7, 6, 4, 9, 10, 12, 11, 14, 16, 15, 13, 8}},
}

func TestPostOrder(t *testing.T) {
	for i, e := range PostOrderTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			acc := make([]int, 0, len(e.expected))
			Postorder[int](e.root, func(node *TreeNode[int]) {
				acc = append(acc, node.Val)
			})
			if !slices.Equal(acc, e.expected) {
				t.Fatalf("%v != %v", acc, e.expected)
			}
		})
	}
}

var LevelOrderTests = []struct {
	root     *TreeNode[int]
	expected []int
}{
	{},
	// 			(1)
	{ToBST([]int{1}), []int{1}},
	//			(2)
	//		(1)
	{ToBST([]int{1, 2}), []int{2, 1}},
	// 	//			(2)
	// 	//		(1)		(3)
	{ToBST([]int{1, 2, 3}), []int{2, 1, 3}},
	//			(3)
	//		(2)		(5)
	//    (1)	  (4)
	{ToBST([]int{1, 2, 3, 4, 5}), []int{3, 2, 5, 1, 4}},
	// 	//			(3)
	// 	//		(1)			(5)
	// 	//	(0)   (2)	  (4)  (6)
	{ToBST([]int{0, 1, 2, 3, 4, 5, 6}), []int{3, 1, 5, 0, 2, 4, 6}},
	// 	// 			(3)
	// 	//         /   \
	// 	//		(1)  	(5)
	// 	//     /   \    /  \
	// 	//    (0)   (2)(4)  (6)
	{ToBST([]int{0, 1, 2, 3, 4, 5, 6}), []int{3, 1, 5, 0, 2, 4, 6}},
	//					(5)
	//			(3)				(8)
	//		(2)		(4)		(7)		(9)
	//	  (1)	 		   (6)
	//
	{ToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), []int{5, 3, 8, 2, 4, 7, 9, 1, 6}},
}

func TestLevelOrder(t *testing.T) {
	for i, e := range LevelOrderTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			acc := make([]int, 0, len(e.expected))
			Levelorder[int](e.root, func(node *TreeNode[int]) {
				acc = append(acc, node.Val)
			})
			if !slices.Equal(acc, e.expected) {
				t.Fatalf("%v != %v", acc, e.expected)
			}
		})
	}
}

// Given an array, make a Binary Tree by taking middle element
// and putting everything to left of it to left, and everything to the right of it
// to the right
// this results in a balanced bst IF the input array is sorted
// O(n)
//
//	  			(3)
//				(2)		(5)
//	      (1)	(4)
func MakeBinaryTree(arr []int) *TreeNode[int] {
	if len(arr) == 0 {
		return nil
	}
	mid := len(arr) / 2
	root := NewTreeNode(arr[mid])
	if mid >= 0 {
		root.Left = MakeBinaryTree(arr[:mid])
	}
	if mid+1 < len(arr) {
		root.Right = MakeBinaryTree(arr[mid+1:])
	}

	return root
}

func MakeBST(arr []int) *TreeNode[int] {
	if len(arr) == 0 {
		return nil
	}
	shuffle(arr)
	root := InsertBST(nil, arr[0])
	for i := 1; i < len(arr); i++ {
		InsertBST(root, arr[i])
	}
	return root
}

// MakeBalancedBST first sorts, and then calls MakeBinaryTree
// O(nlogn)
func MakeBalancedBST(arr []int) *TreeNode[int] {
	slices.Sort(arr)
	return MakeBinaryTree(arr)
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func shuffle(nums []int) {
	for i := range nums {
		swap(nums, i, rand.Intn(len(nums)))
	}
}

func Range(start, end, inc int) []int {
	if end < start || inc <= 0 {
		panic("does not converge")
	}
	result := make([]int, 0)
	cur := start
	for cur <= end {
		result = append(result, cur)
		cur += inc
	}
	return result
}

func ToBST[T cmp.Ordered](sortedarr []T) *TreeNode[T] {
	n := len(sortedarr)
	if n == 0 {
		return nil
	}
	mid := n / 2
	root := NewTreeNode(sortedarr[mid])
	root.Left = ToBST(sortedarr[:mid])
	if mid+1 < n {
		root.Right = ToBST(sortedarr[mid+1:])
	}
	return root
}

var MakeBSTFromPreorderNaiveTests = []struct {
	preorder []int
}{
	{},
	{
		[]int{1},
	},
	// Pre-order partitions a BST
	// [root][left][right]
	{[]int{2, 1, 3}},
	{[]int{3, 2, 1, 5, 4}},
	{[]int{4, 2, 1, 3, 6, 5}},
	{[]int{4, 2, 1, 3, 6, 5, 7}},
	{[]int{9, 5, 3, 2, 1, 4, 7, 6, 8, 13, 11, 10, 12, 15, 14, 16}},
	{Range(0, 16, 1)},
}

func TestMakeBSTFromPreorderNaive(t *testing.T) {
	for i, e := range MakeBSTFromPreorderNaiveTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			inorder := slices.Clone(e.preorder)
			slices.Sort(inorder)
			bst := MakeBSTFromPreorder(e.preorder)
			acc := make([]int, 0, len(e.preorder))
			Preorder(bst, func(node *TreeNode[int]) {
				acc = append(acc, node.Val)
			})
			if !slices.Equal(acc, e.preorder) {
				t.Fatalf("not preorder?: got %v, want %v", acc, e.preorder)
			}
			slices.Sort(acc)
			if !slices.Equal(acc, inorder) {
				t.Fatalf("not inorder?: got %v, want %v", acc, inorder)
			}
		})
	}
}

var MakeBSTFromPostorderNaiveTests = []struct {
	postorder []int
}{
	// {},
	// {
	// 	[]int{1},
	// },
	// // post-order partitions a BST
	// // [left][right][root]
	// // {[]int{1, 3, 2}},
	// //			(3)
	// //		(2)		(5)
	// //    (1)	  (4)
	// {[]int{1, 2, 4, 5, 3}},
	// {[]int{0, 2, 1, 4, 6, 5, 3}},
	// {[]int{0, 1, 3, 2, 5, 7, 6, 4}},
	{[]int{0, 1, 3, 2, 5, 7, 6, 4, 9, 10, 12, 11, 14, 16, 15, 13, 8}},
	// {Range(0, 16, 1)},
}

func TestMakeBSTFromPostorderNaive(t *testing.T) {
	for i, e := range MakeBSTFromPostorderNaiveTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			inorder := slices.Clone(e.postorder)
			slices.Sort(inorder)
			bst := MakeBSTFromPostOrder(e.postorder)
			acc := make([]int, 0, len(e.postorder))
			Postorder(bst, func(node *TreeNode[int]) {
				acc = append(acc, node.Val)
			})
			if !slices.Equal(acc, e.postorder) {
				t.Fatalf("not postorder?: got %v, want %v", acc, e.postorder)
			}
			slices.Sort(acc)
			if !slices.Equal(acc, inorder) {
				t.Fatalf("not inorder?: got %v, want %v", acc, inorder)
			}
		})
	}
}
