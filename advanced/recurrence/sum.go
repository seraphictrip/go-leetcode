package recurrence

func Sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + Sum(n-1)
}
