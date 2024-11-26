package recurrence

// Write a function taht takes two inputs, n and m
// and outputs the number of unique paths from the top
// left corner to bottom right corner of a NxM grid
// Constraints: one can only move down or right 1 unit at a time
func Paths(n, m int) int {
	return paths(n, m, 1, 1)
}

func paths(n, m, i, j int) int {
	if i > n || j > m {
		return 0
	}
	if i == n || j == m {
		return 1
	}
	return paths(n, m, i+1, j) + paths(n, m, i, j+1)
}
