package dp

import "fmt"

/*
The thief has found himself a new place for his thievery again.
There is only one entrance to this area, called root.

Besides the root, each house has one and only one parent house.
After a tour, the smart thief realized that all houses in this place form a binary tree.
It will automatically contact the police if two directly-linked houses were broken into on
the same night.

Given the root of the binary tree, return the maximum amount of money the thief can rob
 without alerting the police.



Example 1:
					(3)
			(2)				(3)
				(3)				(1)
Input: root = [3,2,3,null,3,null,1]
// level order
[3, 5, 4]

Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
Example 2:
			(3)
		(4)		(5)
	(1)		(3)		(1)

Input: root = [3,4,5,1,3,null,1]
Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.


Constraints:

The number of nodes in the tree is in the range [1, 104].
0 <= Node.val <= 104


Being in the robber series this feels like a optimization problem, so sill tackle it that way first
but I think I want to look at level order as well
I think I can calc rob value at each level in linear time and then degrade to house robber 1



*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Rob3(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}
	// take
	take := root.Val
	if root.Left != nil {
		take += Rob3(root.Left.Left) + Rob3(root.Left.Right)
	}
	if root.Right != nil {
		take += Rob3(root.Right.Left) + Rob3(root.Right.Right)
	}

	// skip
	skip := Rob3(root.Left) + Rob3(root.Right)

	return max(take, skip)
}

func Rob3Memo(root *TreeNode, memo map[*TreeNode]int) int {
	// base case
	if root == nil {
		return 0
	}
	if cached, ok := memo[root]; ok {
		return cached
	}
	// take
	take := root.Val
	if root.Left != nil {
		take += Rob3Memo(root.Left.Left, memo) + Rob3Memo(root.Left.Right, memo)
	}
	if root.Right != nil {
		take += Rob3Memo(root.Right.Left, memo) + Rob3Memo(root.Right.Right, memo)
	}

	// skip
	skip := Rob3Memo(root.Left, memo) + Rob3Memo(root.Right, memo)

	memo[root] = max(take, skip)
	return memo[root]
}

func RobBFS(root *TreeNode) int {
	nums := make([]int, 0)
	BFS(root, func(node *TreeNode, level int) {
		if len(nums) < level+1 {
			nums = append(nums, node.Val)
		} else {
			nums[level] += node.Val
		}
	})
	fmt.Println(nums)
	return RobBottomUp(nums)
}

func BFS(root *TreeNode, visit func(node *TreeNode, level int)) {
	if root == nil {
		return
	}
	q := &queue{}
	q.Enqueue(root)
	level := 0
	for !q.IsEmpty() {
		n := q.Size()
		for i := 0; i < n; i++ {
			cur := q.Dequeue()
			visit(cur, level)
			if cur.Left != nil {
				q.Enqueue(cur.Left)
			}
			if cur.Right != nil {
				q.Enqueue(cur.Right)
			}
		}
		level++
	}
}

type queue struct {
	queue []*TreeNode
}

func (q *queue) Enqueue(node *TreeNode) {
	q.queue = append(q.queue, node)
}

func (q *queue) Dequeue() *TreeNode {
	first := q.queue[0]
	if len(q.queue) == 1 {
		q.queue = nil
	} else {
		q.queue = q.queue[1:]
	}
	return first
}

func (q *queue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *queue) Size() int {
	return len(q.queue)
}

func findValOrNextSmallest(bst *TreeNode, x int) int {
	// get the greatest value <= x in a binary search tree
	if bst == nil {
		return -1
	}

	if bst.Val == x {
		return x
	}
	if bst.Val > x {
		return findValOrNextSmallest(bst.Left, x)
	}
	rightBest := findValOrNextSmallest(bst.Right, x)
	if rightBest == -1 {
		return bst.Val
	}
	return rightBest
}
