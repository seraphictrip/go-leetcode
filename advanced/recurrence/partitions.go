package recurrence

// Count the number of ways you can partition n objects using parts upto m
func Partitions(n, m int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if m == 0 {
		return 0
	}
	if m == 1 {
		return 1
	}
	return Partitions(n-m, m) + Partitions(n, m-1)
}
