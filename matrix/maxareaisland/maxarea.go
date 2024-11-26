package maxareaisland

// https://leetcode.com/problems/max-area-of-island/
/*
[	[0,0,1,0,0,0,0,1,0,0,0,0,0],
	[0,0,0,0,0,0,0,1,1,1,0,0,0],
	[0,1,1,0,1,0,0,0,0,0,0,0,0],
	[0,1,0,0,1,1,0,0,1,0,1,0,0],
	[0,1,0,0,1,1,0,0,1,1,1,0,0],
	[0,0,0,0,0,0,0,0,0,0,1,0,0],
	[0,0,0,0,0,0,0,1,1,1,0,0,0],
	[0,0,0,0,0,0,0,1,1,0,0,0,0]]
*/

func MaxAreaOfIsland(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	ROWS, COLS := len(grid), len(grid[0])
	visited := make(map[[2]int]bool)
	maxSize := 0
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if grid[i][j] == 1 && !visited[[2]int{i, j}] {
				// we found an island, we need to get size
				maxSize = max(maxSize, dfs(grid, i, j, visited))
			}
		}
	}
	return maxSize
}

func dfs(grid [][]int, r, c int, visited map[[2]int]bool) int {
	ROWS, COLS := len(grid), len(grid[0])
	coor := [2]int{r, c}
	if r < 0 || c < 0 {
		// out of bounds
		return 0
	}
	if r == ROWS || c == COLS {
		// out of bounds
		return 0
	}
	if grid[r][c] == 0 {
		// water
		return 0
	}
	if visited[coor] {
		return 0
	}

	visited[coor] = true
	area := 1
	area += dfs(grid, r+1, c, visited)
	area += dfs(grid, r-1, c, visited)
	area += dfs(grid, r, c+1, visited)
	area += dfs(grid, r, c-1, visited)

	return area
}
