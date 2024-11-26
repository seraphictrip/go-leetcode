package binarytree

import (
	"cmp"
	"fmt"
	"slices"
)

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewTreeNode[T any](val T) *TreeNode[T] {
	return &TreeNode[T]{
		Val: val,
	}
}

func Height[T any](root *TreeNode[T]) int {
	if root == nil {
		return 0
	}
	return 1 + max(Height[T](root.Left), Height[T](root.Right))
}

// Traverse inorder
// left -> root -> right
// For BST this will be in sort order
func Inorder[T any](root *TreeNode[T], visit func(*TreeNode[T])) {
	if root == nil {
		return
	}
	Inorder(root.Left, visit)
	visit(root)
	Inorder(root.Right, visit)
}

func InorderRightToLeft[T any](root *TreeNode[T], visit func(*TreeNode[T])) {
	if root == nil {
		return
	}
	InorderRightToLeft(root.Right, visit)
	visit(root)
	InorderRightToLeft(root.Left, visit)
}

//				(3)
//			(2)    (5)
//	     (1)     (4)
//
// [1,2,3,4,5]
func InorderIterative[T any](root *TreeNode[T], visit func(*TreeNode[T])) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode[T], 0)

	cur := root
	for cur != nil || len(stack) != 0 {
		// 1. cur == root(3)
		// 2. cur == nil | stack = [3,2]
		// 3. cur = nil | stack = [3]
		// 4. cur = 5 | stack = []
		// 5  cur = nil | stack = [5]
		// Inorder is depth first, left->cur->right
		// so push unto stack all the way to the left
		for cur != nil {
			// push
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 1. [3, 2, 1]
		// 2. [3, 2]
		// 3. [3]
		// 4. [5, 4]
		// 5. [5]
		// pop
		// 1. (1)
		// 2. (2)
		// 3. (3)
		// 4. (4)
		// 5. (5)
		topIndex := len(stack) - 1
		cur = stack[topIndex]
		stack = stack[0:topIndex]
		// 1. [3, 2]
		// 2. [2]
		// 3. []
		// 4. [5]
		// 5 []
		visit(cur) // 1. 1 // 2. 2 // 3. 3 // 4. 4 // 5. 5
		cur = cur.Right
	}
}

func PreorderIterative[T any](root *TreeNode[T], visit func(*TreeNode[T])) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode[T], 0)

	// Push
	stack = append(stack, root)
	// !IsEmpty
	for len(stack) != 0 {
		// Pop
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		visit(top)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}

	}

}

type stack[T any] struct {
	stack []T
}

func (s *stack[T]) Push(item T) {
	s.stack = append(s.stack, item)
}

func (s *stack[T]) Pop() T {
	topIndex := len(s.stack) - 1
	result := s.stack[topIndex]
	s.stack = s.stack[0:topIndex]
	return result
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

// Traverse Preorder
// root -> left -> right
func Preorder[T any](root *TreeNode[T], visit func(*TreeNode[T])) {
	if root == nil {
		return
	}
	visit(root)
	Preorder(root.Left, visit)
	Preorder(root.Right, visit)
}

// Traverse Postorder
// left->right->root
func Postorder[T any](root *TreeNode[T], visit func(*TreeNode[T])) {
	if root == nil {
		return
	}
	Postorder(root.Left, visit)
	Postorder(root.Right, visit)
	visit(root)
}

//				(3)
//			(2)    (5)
//	     (1)     (4)
//
// [1,2,4,5,3]
func PostorderIterative[T any](root *TreeNode[T], visit func(*TreeNode[T])) {
	stack := &stack[*TreeNode[T]]{}

	cur := root
	for cur != nil || !stack.IsEmpty() {
		for cur.Left != nil {
			stack.Push(cur.Left)
			cur = cur.Left
		}

		for cur.Right != nil {
			stack.Push(cur.Right)
			cur = cur.Right
		}
		visit(cur)
	}
}

// Traverse levelorder
// h0, h1....
func Levelorder[T any](root *TreeNode[T], visit func(*TreeNode[T])) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode[T], 0)

	queue = append(queue, root)

	for len(queue) != 0 {
		cur := queue[0]
		if len(queue) == 1 {
			queue = nil
		} else {
			queue = queue[1:]
		}
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}

		visit(cur)

	}
}

func Collect[T any](root *TreeNode[T]) []T {
	result := make([]T, 0)
	Inorder(root, func(node *TreeNode[T]) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
	})
	return result
}

// BST are just BinaryTrees with additional constraints
// Everything to the Left is less than Parent, everything to the right is greater than parent
//

func Search[T cmp.Ordered](bst *TreeNode[T], target T) bool {
	if bst == nil {
		return false
	}
	if bst.Val == target {
		return true
	}
	if target < bst.Val {
		return Search(bst.Left, target)
	}

	if target > bst.Val {
		return Search(bst.Right, target)
	}
	// shouldn't be able to get here, but rather check for equal
	return false
}

// BST Insert
// Insert will insert the node in an appropriate place to maintain
// BST property, of all nodes to left are less than and all nodes to right are greater
// Insert does not maintain a balanced BST, but does maintain the BST property
// all values to the left are less than, all values right are greater than node val
//
//	(4)
var count = 0

func InsertBST[T cmp.Ordered](root *TreeNode[T], val T) *TreeNode[T] {
	fmt.Println(count, val, root)
	count++
	if root == nil {
		return NewTreeNode(val)
	}
	if val > root.Val {
		root.Right = InsertBST(root.Right, val)
	}
	if val < root.Val {
		root.Left = InsertBST(root.Left, val)
	}
	return root
}

// Remove the node with value val
// if node does not exist do nothing
// else swap with appropriate child value
// O(logN), traverse height of tree to remove, traverse again if need to swap
func Remove[T cmp.Ordered](root *TreeNode[T], val T) *TreeNode[T] {
	if root == nil {
		return nil
	}

	if val < root.Val {
		root.Left = Remove(root.Left, val)
	} else if val > root.Val {
		root.Right = Remove(root.Right, val)
	} else {
		// val == root.Val
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// Propagate up min value from Right subtree
			minNode := MinValueNode(root.Right)
			root.Val = minNode.Val
			// remove node we propagated up (it will be a leaf node or have only a right child)
			root.Right = Remove(root.Right, minNode.Val)
		}
	}

	return root
}

// Return the node with the min value in a BST
// Behavior is undefined if called on a BinaryTree that is not a BST
func MinValueNode[T cmp.Ordered](root *TreeNode[T]) *TreeNode[T] {
	// recursive
	// if root == nil {
	// 	return nil
	// }
	// if root.Left == nil {
	// 	return root
	// }
	// return MinValueNode(root.Left)
	// iterative
	cur := root
	for cur != nil && cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

// Return the node with the max value in a BST
// Behavior is undefined if called on a BinaryTree that is not a BST
func MaxValueNode[T cmp.Ordered](root *TreeNode[T]) *TreeNode[T] {
	cur := root
	for cur != nil && cur.Right != nil {
		cur = cur.Right
	}
	return cur
}

/*
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.



Example 1:


Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
Example 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]


Constraints:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.

*/

// This is only adjecent to problem, but intersting to solve
// to up my understanding
// O(n^2) in worst case (but that prob means source wasn't balanced tree)
// probably O(nlogn) if source is blanced tree
func MakeBSTFromPreorderNaive[T cmp.Ordered](preorder []T) *TreeNode[T] {
	if len(preorder) == 0 {
		return nil
	}
	root := NewTreeNode(preorder[0])

	for i := 1; i < len(preorder); i++ {
		InsertBST(root, preorder[i])
	}
	return root
}

// Preorder orders partition *[root][less][greater]
// this partitioning [root][less][greater] is the big take away
// and it maps directly onto BST+preorder definition
func MakeBSTFromPreorder[T cmp.Ordered](preorder []T) *TreeNode[T] {
	fmt.Println(count, preorder)
	count++
	n := len(preorder)
	if n == 0 {
		// base case, empty array
		return nil
	}
	// peel off root
	root := NewTreeNode(preorder[0])
	// find index of firstGreater to put on right
	firstGreater := slices.IndexFunc(preorder, func(n T) bool {
		return n > root.Val
	})

	// if a greater element was found recursively call
	if firstGreater != -1 {
		root.Left = MakeBSTFromPreorder(preorder[1:firstGreater])
		root.Right = MakeBSTFromPreorder(preorder[firstGreater:])
	} else {
		// otherwise everything goes on the left
		root.Left = MakeBSTFromPreorder(preorder[1:])
	}
	return root
}

//			(3)
//	        /   \
//			(1)  	(5)
//	    /   \    /  \
//	   (0)   (2)(4)  (6)
//
// []int{0, 2, 1, 4, 6, 5, 3}},
// [<(less)][>(greater)][root]
func MakeBSTFromPostOrderNaive[T cmp.Ordered](postorder []T) *TreeNode[T] {
	// iterate backwards, seems like it would be linear, but I'm a fool about such things
	// so lets explore
	n := len(postorder)
	if n == 0 {
		return nil
	}
	root := NewTreeNode(postorder[n-1])
	for i := n - 2; i >= 0; i-- {
		// same problem as before, start at beginning always
		InsertBST(root, postorder[i])
	}

	return root
}

// []int{0, 2, 1, 4, 6, 5, 3}},
// [<(less)][>(greater)][root]
func MakeBSTFromPostOrder[T cmp.Ordered](postorder []T) *TreeNode[T] {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	// peel root off end
	root := NewTreeNode(postorder[n-1])

	// find first element of "right" partition
	firstGreater := slices.IndexFunc(postorder, func(n T) bool {
		return n > root.Val
	})

	if firstGreater != -1 {
		// recursively build left and right tree
		root.Left = MakeBSTFromPostOrder(postorder[:firstGreater])
		root.Right = MakeBSTFromPostOrder(postorder[firstGreater : n-1])
	} else {
		// if no right partiion, just do "left" (I am the complement)
		root.Left = MakeBSTFromPostOrder(postorder[:n-1])
	}
	return root
}

// [ <P][P][>P]
func MakeBSTFromInorder[T cmp.Ordered](inorder []T) *TreeNode[T] {
	n := len(inorder)

	if n == 0 {
		return nil
	}
	mid := n / 2
	root := NewTreeNode(inorder[mid])
	root.Left = MakeBSTFromInorder(inorder[:mid])
	if mid+1 < n {
		root.Right = MakeBSTFromInorder(inorder[mid+1:])
	}
	return root
}

// [ <P ] [P] [ >P]
func Insert[T cmp.Ordered](root *TreeNode[T], val T) *TreeNode[T] {
	if root == nil {
		return NewTreeNode(val)
	}
	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else {
		// val > root.Val
		root.Right = Insert(root.Right, val)
	}
	return root

}
