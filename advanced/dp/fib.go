package dp

import "fmt"

// 1, 1, 2, 3, 5, 8, 13...
func Fib(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// adding memoization isn't bad on 1 dimension array
func FibMemo(n int) int {
	memo := make(map[int]int, n)
	var fib func(n int, memo map[int]int) int

	fib = func(n int, memo map[int]int) int {
		if n <= 1 {
			return 1
		}
		if cached, ok := memo[n]; ok {
			return cached
		}
		one := fib(n-1, memo)
		two := fib(n-2, memo)

		memo[n] = one + two
		fmt.Printf("%v + %v = %v\n", one, two, memo[n])
		return memo[n]
	}
	return fib(n, memo)

}

func FibBottomUp(n int) int {
	table := make([]int, max(n+1, 2))
	table[0], table[1] = 1, 1
	for i := 2; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n]
}

// bottom up, but only store most recent values
// we basically just throw away values we no longer need
func FibDP(n int) int {
	table := [2]int{1, 1}
	for i := 2; i <= n; i++ {
		tmp := table[0]
		table[0] = table[1]
		table[1] = table[1] + tmp
	}
	return table[1]
}
