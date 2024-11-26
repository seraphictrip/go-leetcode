package knapsack

import "fmt"

/*
You are given an array of binary strings strs and two integers m and n.

Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.

A set x is a subset of a set y if all elements of x are also elements of y.



Example 1:

Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3
Output: 4
Explanation: The largest subset with at most 5 0's and 3 1's is {"10", "0001", "1", "0"}, so the answer is 4.
Other valid but smaller subsets include {"0001", "1"} and {"10", "1", "0"}.
{"111001"} is an invalid subset because it contains 4 1's, greater than the maximum of 3.
Example 2:

Input: strs = ["10","0","1"], m = 1, n = 1
Output: 2
Explanation: The largest subset is {"0", "1"}, so the answer is 2.


Constraints:

1 <= strs.length <= 600
1 <= strs[i].length <= 100
strs[i] consists only of digits '0' and '1'.
1 <= m, n <= 100
*/

func FindMaxForm(strs []string, m int, n int) int {
	compact := make([][2]int, len(strs))
	for i := 0; i < len(strs); i++ {
		pair := [2]int{}
		for _, ch := range strs[i] {
			if ch == '0' {
				pair[0]++
			} else {
				pair[1]++
			}
		}
		compact[i] = pair
	}

	var helper func(i, zeros, ones, count int) int
	helper = func(i, zeros, ones, count int) int {
		if i == len(compact) {
			if zeros >= 0 && ones >= 0 {
				return count
			}
			return 0
		}
		take := helper(i+1, zeros-compact[i][0], ones-compact[i][1], count+1)
		skip := helper(i+1, zeros, ones, count)
		return max(take, skip)
	}
	return helper(0, m, n, 0)
}

func compactRepresentation(strs []string) [][2]int {
	compact := make([][2]int, len(strs))
	for i := 0; i < len(strs); i++ {
		pair := [2]int{}
		for _, ch := range strs[i] {
			if ch == '0' {
				pair[0]++
			} else {
				pair[1]++
			}
		}
		compact[i] = pair
	}
	return compact
}

func FindMaxFormMemo(strs []string, m int, n int) int {
	compact := compactRepresentation(strs)

	memo := make(map[[3]int]int, m*n*len(strs))
	var helper func(i, zeros, ones int) int
	helper = func(i, zeros, ones int) int {
		if i == len(compact) {
			return 0
		}
		key := [3]int{i, zeros, ones}
		if val, ok := memo[key]; ok {
			return val
		}
		take := 0
		zCnt, oCnt := zeros-compact[i][0], ones-compact[i][1]
		if zCnt >= 0 && oCnt >= 0 {
			take = 1 + helper(i+1, zCnt, oCnt)
		}

		skip := helper(i+1, zeros, ones)
		memo[key] = max(take, skip)
		return memo[key]
	}
	ans := helper(0, m, n)
	return ans
}

func FindMaxFormBottomUp(strs []string, m int, n int) int {
	compact := compactRepresentation(strs)

	table := make(map[[3]int]int, m*n*len(strs))
	// ORDER:
	// i, zeros, ones require:
	// SKIP:
	//		i+1
	//		Bigger before smaller
	//	TAKE:
	//		i+1, zeros-usedZeros, ones-usedOnes
	//		smaller before bigger for zeros and ones
	helper := func(i, zeros, ones int) int {
		// if i == len(compact) {
		// 	return 0
		// }
		key := [3]int{i, zeros, ones}
		// if val, ok := table[key]; ok {
		// 	return val
		// }
		take := 0
		zCnt, oCnt := zeros-compact[i][0], ones-compact[i][1]
		if zCnt >= 0 && oCnt >= 0 {
			take = 1 + table[[3]int{i + 1, zCnt, oCnt}]
		}

		skip := table[[3]int{i + 1, zeros, ones}]
		table[key] = max(take, skip)
		return table[key]
	}
	for z := 0; z <= m; z++ {
		for o := 0; o <= n; o++ {
			for i := len(strs) - 1; i >= 0; i-- {
				table[[3]int{i, z, o}] = helper(i, z, o)
			}
		}
	}
	fmt.Println(len(table), (m+1)*(n+1)*len(strs))
	return table[[3]int{0, m, n}]
}
