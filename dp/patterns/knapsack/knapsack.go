package knapsack

import (
	"fmt"
	"math"
)

// 0/1 knapsack pattern

// Given a list of N items, and a backpack with a limited capacitiy , return the maximum total profit that ca be contained in the backpack.
// The ith items profit is profit[i], and it's weight is weight[i].  Assuem you can only add each iteme to the bag at most one time.

// PROFIT: [4,4,7,1]
// WEIGHT: [5,2,3,1]
// CAP: 8

// 0/1 because include or exclude
func KnapsackBruteForce(profit, weight []int, cap int) int {
	var helper func(i int, cap int) int
	// ORDER:
	// terms i, cap requires
	//		take: i+1, cap-weight[i]
	helper = func(i, cap int) int {
		if cap < 0 {
			return math.MinInt
		}
		if cap == 0 || i == len(profit) {
			return 0
		}
		take := profit[i] + helper(i+1, cap-weight[i])
		skip := helper(i+1, cap)
		return max(take, skip)
	}

	return helper(0, cap)
}

func KnapsackTopDown(profit, weight []int, cap int) int {
	memo := make([][]int, len(profit))
	for i := range len(profit) {
		// init to -1
		memo[i] = make([]int, cap+1)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}
	var helper func(i int, cap int) int
	helper = func(i, cap int) int {
		if cap < 0 {
			return math.MinInt
		}
		if cap == 0 || i == len(profit) {
			return 0
		}
		if memo[i][cap] != -1 {
			return memo[i][cap]
		}
		take := profit[i] + helper(i+1, cap-weight[i])
		skip := helper(i+1, cap)
		memo[i][cap] = max(take, skip)
		return memo[i][cap]
	}

	ans := helper(0, cap)
	return ans
}

func KnapsackBottomUp(profit, weight []int, cap int) int {
	N, M := len(profit), cap
	table := make([][]int, len(profit))
	for i := range len(profit) {
		// init to -1
		table[i] = make([]int, cap+1)
		for j := 0; j < len(table[i]); j++ {
			table[i][j] = 0
		}
	}

	for j := 0; j < M+1; j++ {
		if weight[0] <= j {
			table[0][j] = profit[0]
		}
	}

	for i := 1; i < N; i++ {
		for j := 1; j < M+1; j++ {
			skip := table[i-1][j]
			include := 0
			if j-weight[i] >= 0 {
				include = profit[i] + table[i-1][j-weight[i]]
			}
			table[i][j] = max(include, skip)
		}
	}
	fmt.Println(table)
	return table[N-1][M]

}

func UnboundedKnapsackBruteForce(profit, weight []int, cap int) int {
	N := len(profit)
	var helper func(i, cap int) int

	helper = func(i, cap int) int {
		if i == N {
			return 0
		}
		take := 0
		if cap-weight[i] >= 0 {
			take = profit[i] + helper(i, cap-weight[i])
		}
		skip := helper(i+1, cap)
		return max(take, skip)
	}
	return helper(0, cap)
}

func UnboundedKnapsackTopDown(profit, weight []int, cap int) int {
	N := len(profit)
	M := cap

	memo := buildInitialMemo(N, M+1, -1)

	var helper func(i, cap int) int

	helper = func(i, cap int) int {
		if i == N {
			return 0
		}
		if memo[i][cap] != -1 {
			return memo[i][cap]
		}
		take := 0
		if cap-weight[i] >= 0 {
			take = profit[i] + helper(i, cap-weight[i])
		}
		skip := helper(i+1, cap)
		memo[i][cap] = max(take, skip)
		return memo[i][cap]
	}
	return helper(0, cap)
}

func UnboundedKnapsackBottomUp(profit, weight []int, cap int) int {
	N := len(profit)
	M := cap

	table := buildInitialMemo(N+1, M+1, -1)

	// ORDER:
	// 		i, cap requires:
	//	TAKE:
	//		smaller cap before larger cap
	//	SKIP:
	//		larger i before smaller i

	var helper func(i, cap int) int

	helper = func(i, cap int) int {
		if i == N {
			return 0
		}
		// if table[i][cap] != -1 {
		// 	return table[i][cap]
		// }
		take := 0
		if cap-weight[i] >= 0 {
			take = profit[i] + table[i][cap-weight[i]]
		}
		skip := table[i+1][cap]
		table[i][cap] = max(take, skip)
		return table[i][cap]
	}

	for c := 0; c <= cap; c++ {
		for i := len(profit); i >= 0; i-- {
			table[i][c] = helper(i, c)
		}
	}

	fmt.Println(table)

	return table[0][cap]
}

func buildInitialMemo(N, M, defaultValue int) [][]int {
	memo := make([][]int, N)
	for i := 0; i < N; i++ {
		memo[i] = make([]int, M)
		for j := 0; j < M; j++ {
			memo[i][j] = defaultValue
		}
	}
	return memo
}

// unbounded knapsack
// see: https://www.youtube.com/watch?v=NA7u5GTh6fw&t=558s
// this walks through steps for dp + unwrapping to iterative
//
// [1,2,5]
// Take ist coin:
// 		- have to make amount - coin[i] = newAmount
// skip ith coin:
//		- have to move to the ith + 1 coin
//			5
//		4		5
//	  3   4

//	 0 : [0,1,1,1,1,1]
//	 1 : [0,0,2,3,0,0]
//		2 : [0,0,0,0,0,0]
func CoinChange(amount int, coins []int) int {
	var helper func(i, amount int) int
	helper = func(i, amount int) int {
		if i == len(coins) || amount < 0 {
			return 0
		}
		if amount == 0 {
			return 1
		}
		take := 0
		if amount-coins[i] >= 0 {
			take = helper(i, amount-coins[i])
		}
		skip := helper(i+1, amount)
		return take + skip
	}
	return helper(0, amount)
}

func CoinChangeTopDown(amount int, coins []int) int {
	N, M := len(coins), amount
	var helper func(i, amount int) int
	memo := make([][]int, N)
	for i := range memo {
		memo[i] = make([]int, M+1)
		for j := range memo[i] {
			if j == 0 {
				memo[i][j] = 0
			} else {
				memo[i][j] = -1
			}
		}
	}
	count := 0
	helper = func(i, amount int) int {
		fmt.Printf("%v: helper(%v, %v)\n", count, i, amount)
		count++
		if i == len(coins) || amount < 0 {
			return 0
		}
		if amount == 0 {
			return 1
		}
		if memo[i][amount] != -1 {
			return memo[i][amount]
		}
		take := 0
		if amount-coins[i] >= 0 {
			take = helper(i, amount-coins[i])
		}

		skip := helper(i+1, amount)
		memo[i][amount] = take + skip
		return memo[i][amount]
	}
	ans := helper(0, amount)
	return ans
}

/*
ORDER:

		terms i, amount requires:
			Take: i, amount - coins[i]
				- i does not change so amount depends on amount - coins[i]
				- smaller values before larger for amount

			Skip: i+1, amount
				- amount does not change so i depends on i
				-bigger values before smaller values for i

	[0,1,1,1,1,1]
	[0,1,2,2,3,0]
	[0,0,0,0,0,0]
	[0,0,0,0,0,0]
*/
func CoinChangeBottomUp(amount int, coins []int) int {
	N, M := len(coins), amount
	table := make([][]int, N+1)
	for i := range table {
		table[i] = make([]int, M+1)
	}

	helper := func(i, x int) int {
		if i == len(coins) {
			if x == 0 {
				return 1
			}
			return 0
		}
		take := 0
		if coins[i] <= x {
			take = table[i][x-coins[i]]
		}
		skip := table[i+1][x]
		return take + skip
	}

	for x := 0; x <= amount; x++ {
		for i := len(coins); i >= 0; i-- {
			table[i][x] = helper(i, x)
		}
	}
	return table[0][amount]
}

func CoinChange1(coins []int, amount int) int {
	var helper func(i, x int) int
	helper = func(i, x int) int {
		if i >= len(coins) || x < 0 {
			// I might add one, so give self space to not overflow
			return math.MaxInt - 1
		}
		if x == 0 {
			return 0
		}
		take := math.MaxInt - 1
		if x-coins[i] >= 0 {
			take = 1 + helper(i, x-coins[i])
		}
		skip := helper(i+1, x)
		return min(take, skip)
	}

	ans := helper(0, amount)
	if ans > 1e4 {
		return -1
	}
	return ans
}

func CoinChange1TopDown(coins []int, amount int) int {
	var helper func(i, x int) int
	memo := buildInitialMemo(len(coins), amount+1, -1)
	helper = func(i, x int) int {
		if i >= len(coins) || x < 0 {
			// I might add one, so give self space to not overflow
			return math.MaxInt - 1
		}
		if x == 0 {
			return 0
		}

		if memo[i][x] != -1 {
			return memo[i][x]
		}

		take := math.MaxInt - 1
		if x-coins[i] >= 0 {
			take = 1 + helper(i, x-coins[i])
		}
		skip := helper(i+1, x)
		memo[i][x] = min(take, skip)
		return memo[i][x]
	}

	ans := helper(0, amount)
	if ans > 1e4 {
		return -1
	}
	return ans
}

func CoinChange1BottomUp(coins []int, amount int) int {
	var helper func(i, x int) int
	table := buildInitialMemo(len(coins)+1, amount+1, 0)
	// ORDER: i and x require
	//	TAKE:
	//		x-coins[i]
	//		smaller before bigger
	//	SKIP:
	//		i+1
	//		bigger before smaller for i
	helper = func(i, x int) int {
		if i >= len(coins) || x < 0 {
			// I might add one, so give self space to not overflow
			return math.MaxInt - 1
		}
		if x == 0 {
			return 0
		}

		take := math.MaxInt - 1
		if x-coins[i] >= 0 {
			take = 1 + table[i][x-coins[i]]
		}
		skip := table[i+1][x]
		table[i][x] = min(take, skip)
		return table[i][x]
	}

	for x := 0; x <= amount; x++ {
		for i := len(coins); i >= 0; i-- {
			table[i][x] = helper(i, x)
		}
	}

	ans := table[0][amount]
	if ans > 1e4 {
		return -1
	}
	return ans
}
