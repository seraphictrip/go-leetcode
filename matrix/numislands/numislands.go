package numislands

/*
https://leetcode.com/problems/number-of-islands/description/
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water),
return the number of islands.

An island is surrounded by water and is formed by connecting

adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.



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

func NumIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	ROWS, COLS := len(grid), len(grid[0])
	visited := make(map[[2]int]bool)
	islands := 0
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if grid[i][j] == '1' && !visited[[2]int{i, j}] {
				islands++
				dfs(grid, i, j, visited)
			}
		}
	}
	return islands
}

func dfs(grid [][]byte, r, c int, visited map[[2]int]bool) {
	ROWS, COLS := len(grid), len(grid[0])
	// run off top or left
	if r < 0 || c < 0 {
		return
	}
	// run off bottom or right
	if r == ROWS || c == COLS {
		return
	}
	coor := [2]int{r, c}
	// already visited
	if visited[coor] {
		return
	}
	if grid[r][c] != '1' {
		return
	}
	visited[coor] = true
	dfs(grid, r+1, c, visited)
	dfs(grid, r-1, c, visited)
	dfs(grid, r, c+1, visited)
	dfs(grid, r, c-1, visited)

}
