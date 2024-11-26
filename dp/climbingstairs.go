package dp

import "fmt"

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step


Constraints:

1 <= n <= 45
*/

func climbStairsRecursive(n int) int {
	if n < 0 {
		// I've overstepped
		return 0
	}
	if n == 0 {
		// I've landed
		return 1
	}
	return climbStairsRecursive(n-1) + climbStairsRecursive(n-2)
}

// TOP DOWN
func ClimbStairsTopDown(n int) int {
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	result := climbStairsMemo(n, memo)
	fmt.Println(memo)
	return result
}
func climbStairsMemo(n int, memo []int) int {
	if n < 0 {
		// I've overstepped
		return 0
	}
	if n == 0 {
		// I've landed
		return 1
	}
	if memo[n] != -1 {
		return memo[n]
	}
	memo[n] = climbStairsMemo(n-1, memo) + climbStairsMemo(n-2, memo)
	return memo[n]
}

// bottom up
func ClimbStairs(n int) int {
	table := make([]int, max(n+1, 2))
	table[0] = 1
	table[1] = 1

	for i := 2; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}

	return table[n]
}

func ClimbingStairsConstantSpace(n int) int {
	table := []int{1, 1}
	table[0] = 1
	table[1] = 1

	for i := 2; i <= n; i++ {
		tmp := table[1]
		table[1] = table[0] + table[1]
		table[0] = tmp
	}

	return table[1]
}
