package divide

/*
We can construct a Quad-Tree from a two-dimensional area using the following steps:

If the current grid has the same value (i.e all 1's or all 0's) set isLeaf True and set val
to the value of the grid and set the four children to Null and stop.
If the current grid has different values, set isLeaf to False and set val to any value and divide the current grid into four sub-grids as shown in the photo.
Recurse for each of the children with the proper sub-grid.
*/

// Definition for a QuadTree node.
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var helper func(row0, col0, width int) *Node
	helper = func(row0, col0, width int) *Node {
		if width == 1 {
			return &Node{
				Val:    grid[row0][col0] == 1,
				IsLeaf: true,
			}
		}
		val := grid[row0][col0]
		singleNode := true
		for i := row0; i < row0+width; i++ {
			if !singleNode {
				break
			}
			for j := col0; j < col0+width; j++ {
				if grid[i][j] != val {
					singleNode = false
					break
				}
			}
		}
		var node *Node
		if singleNode {
			node = &Node{
				Val:    val == 1,
				IsLeaf: true,
			}
		} else {
			w := width / 2
			node = &Node{
				Val:         true,
				IsLeaf:      false,
				TopLeft:     helper(row0, col0, w),
				TopRight:    helper(row0, col0+w, w),
				BottomLeft:  helper(row0+w, col0, w),
				BottomRight: helper(row0+w, col0+w, w),
			}
		}
		return node
	}

	return helper(0, 0, len(grid))
}
