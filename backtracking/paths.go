package backtracking

/*
Determine if a path exists from teh root of tree to a leaf node.  It may not contain any zeros
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

func HasValidPath(root *Node) bool {
	if root == nil || root.Val == 0 {
		return false
	}
	if root.Left == nil && root.Right == nil {
		// we are at a leaf and not zero, we win
		return true
	}
	if HasValidPath(root.Left) {
		return true
	}
	return HasValidPath(root.Right)
}

func GetPath(root *Node, path *Intstack) bool {
	if root == nil || root.Val == 0 {
		return false
	}
	path.Push(root.Val)
	if root.Left == nil && root.Right == nil {
		// leaf node
		return true
	}
	if GetPath(root.Left, path) {
		return true
	}
	if GetPath(root.Right, path) {
		return true
	}
	// backtrack
	path.Pop()
	return false
}

type Intstack []int

func (s *Intstack) Push(val int) {
	*s = append(*s, val)
}

func (s *Intstack) Pop() int {
	old := *s
	lastIndex := len(*s) - 1
	item := old[lastIndex]
	*s = old[:lastIndex]
	return item
}
