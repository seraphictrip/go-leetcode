package fibonacci

// 1, 1, 2, 3, 5, 8, 13...
func fibonacci(n int) int {
	memo := make(map[int]int, n)
	return fibMemo(n, memo)

}

func fibMemo(n int, memo map[int]int) int {
	if n < 1 {
		return 1
	}
	if n, ok := memo[n]; ok {
		return n
	}
	return fibMemo(n-1, memo) + fibMemo(n-2, memo)
}
