package dfs

/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.



Example 1:

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3


Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*/

type coordinate [2]int

func NumOfIslands(matrix [][]byte) int {
	ROWS, COLS := len(matrix), len(matrix[0])

	visited := make(map[coordinate]bool)
	islands := 0
	var dfs func(row, col int)

	dfs = func(row, col int) {
		if (row < 0 || row >= ROWS) || (col < 0 || col >= COLS) {
			// out of bounds
			return
		}
		coor := [2]int{row, col}
		if visited[coor] {
			// already visted
			return
		}
		if matrix[row][col] == '0' {
			// water
			return
		}
		visited[coor] = true
		// right
		dfs(row, col+1)
		// down
		dfs(row+1, col)
		// left
		dfs(row, col-1)
		// up
		dfs(row-1, col)
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			coor := [2]int{i, j}
			if matrix[i][j] == '1' && !visited[coor] {
				islands++
				dfs(i, j)
			}
		}
	}

	return islands
}
