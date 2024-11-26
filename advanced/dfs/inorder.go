package dfs

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func Inorder(root *TreeNode) []int {
	return inorder(root, []int{})
}

func inorder(root *TreeNode, acc []int) []int {
	if root == nil {
		return acc
	}
	acc = inorder(root.Left, acc)
	acc = append(acc, root.Val)
	acc = inorder(root.Right, acc)

	return acc
}

func Preorder(root *TreeNode) []int {
	return preorder(root, []int{})
}

func preorder(root *TreeNode, acc []int) []int {
	if root == nil {
		return acc
	}
	acc = append(acc, root.Val)
	acc = preorder(root.Left, acc)
	acc = preorder(root.Right, acc)

	return acc
}

func InorderIterative(root *TreeNode) []int {

	result := make([]int, 0)
	var stack Stack = make([]*TreeNode, 0)
	cur := root

	for cur != nil || !stack.IsEmpty() {
		// go as far left as can
		if cur != nil {
			stack.Push(cur)
			cur = cur.Left
		} else {
			// process node
			node := stack.Pop()
			result = append(result, node.Val)
			// add right to be processed
			cur = node.Right
		}
	}

	return result
}

func CreateInorder(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	mid := n / 2

	root := NewTreeNode(nums[mid])
	root.Left = CreateInorder(nums[:mid])
	if mid+1 < n {
		root.Right = CreateInorder(nums[mid+1:])
	}

	return root
}

func CreatePreorder(nums []int) *TreeNode {
	idx := 0
	return createpreorder(nums, &idx, math.MinInt, math.MaxInt)
}

func createpreorder(nums []int, idx *int, min, max int) *TreeNode {
	if *idx >= len(nums) {
		return nil
	}
	key := nums[*idx]
	if key <= min || key >= max {
		return nil
	}
	root := NewTreeNode(key)
	*idx++

	root.Left = createpreorder(nums, idx, min, key)
	root.Right = createpreorder(nums, idx, key, max)
	return root
}

type Stack []*TreeNode

func (stack *Stack) Push(item *TreeNode) {
	*stack = append(*stack, item)
}

func (stack *Stack) Pop() *TreeNode {
	old := *stack
	val := old[len(old)-1]
	*stack = old[:len(old)-1]
	return val
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
