package pitrav

import "slices"

/*
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree
 and inorder is the inorder traversal of the same tree, construct and return the binary tree.



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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//
//		(3)
//	(9)		(20)
//		(15)	(7)
//
// preorder: [P][LEFT][RIGHT]
// postorder: [LEFT][P][RIGHT]
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	// we know our root value from preorder
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	rootInorderIndex := slices.Index(inorder, rootVal)
	// inorder is [LEFT][P][RIGHT] so we can not get partitions from inorder
	inorderLeft := inorder[:rootInorderIndex]
	llen := len(inorderLeft)
	// using size of inorder partition we can get slice from preorder
	// as preorder is [P][LEFT][RIGHT]

	// add one as slice is eclusive [)
	// this makes this the first elem in right
	firstRight := 1 + llen

	preorderLeft := preorder[1:firstRight]
	root.Left = buildTree(preorderLeft, inorderLeft)

	// We can do the same for right, assuming we are not already at end of array
	if rootInorderIndex+1 < n {
		// [LEFT][P][RIGHT], so right is just everything to right
		inorderRight := inorder[rootInorderIndex+1:]
		// we alrady calculated partition boundary
		preorderRight := preorder[firstRight:]
		root.Right = buildTree(preorderRight, inorderRight)
	}
	return root
}

/*
Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.



Example 1:


Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]
Example 2:

Input: inorder = [-1], postorder = [-1]
Output: [-1]


Constraints:

1 <= inorder.length <= 3000
postorder.length == inorder.length
-3000 <= inorder[i], postorder[i] <= 3000
inorder and postorder consist of unique values.
Each value of postorder also appears in inorder.
inorder is guaranteed to be the inorder traversal of the tree.
postorder is guaranteed to be the postorder traversal of the tree.
*/
//			-1    0    1					-1	    1	 0
// inorder [LEFT][P][RIGHT] and post order [LEFT][RIGHT][P]
func BuildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	// get root val from postorder
	pVal := postorder[len(postorder)-1]
	// build rootNode
	root := &TreeNode{Val: pVal}
	// find root val inorder
	ioIndex := slices.Index(inorder, pVal)
	// Get partitions from inorder [LEFT][P][RIGHT]
	ioLeft := inorder[0:ioIndex]
	// Use size of partitions to get paritions in postorder
	firstRightIndex := len(ioLeft)
	poLeft := inorder[0:firstRightIndex]
	// [LEFT][RIGHT][P]

	// build left sub tree
	root.Left = buildTree(ioLeft, poLeft)

	if ioIndex+1 < n {
		// do same for right if we didn't run off edge
		ioRight := inorder[ioIndex+1:]
		poRight := postorder[firstRightIndex : len(postorder)-1]
		root.Right = buildTree(ioRight, poRight)
	}
	return root
}
