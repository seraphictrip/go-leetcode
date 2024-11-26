package lcs

import "fmt"

/*
Given two strings s and t, return the number of distinct subsequences of s which equals t.

The test cases are generated so that the answer fits on a 32-bit signed integer.



Example 1:

Input: s = "rabbbit", t = "rabbit"
Output: 3
Explanation:
As shown below, there are 3 ways you can generate "rabbit" from s.
rabbbit
rabbbit
rabbbit
Example 2:

Input: s = "babgbag", t = "bag"
Output: 5
Explanation:
As shown below, there are 5 ways you can generate "bag" from s.
babgbag
babgbag
babgbag
babgbag
babgbag
*/

func NumDistinct(s string, t string) int {
	var helper func(i, j int) int
	helper = func(i, j int) int {
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 0
		}
		if s[i] != t[j] {
			return helper(i+1, j)
		} else {
			return helper(i+1, j+1) + helper(i+1, j)
		}
	}
	return helper(0, 0)
}

func NumDistinctMemo(s string, t string) int {
	var helper func(i, j int) int
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(t)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	helper = func(i, j int) int {
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if s[i] != t[j] {
			memo[i][j] = helper(i+1, j)
		} else {
			memo[i][j] = helper(i+1, j+1) + helper(i+1, j)
		}
		return memo[i][j]
	}
	return helper(0, 0)
}

func NumDistinctBottomUp(s string, t string) int {

	table := make([][]int, len(s)+1)
	for i := range table {
		table[i] = make([]int, len(t)+1)
	}
	// ORDER:
	// i, j requries
	// i+1 large before small
	// j+1 large before small

	helper := func(i, j int) int {
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 0
		}
		if s[i] != t[j] {
			table[i][j] = table[i+1][j]
		} else {
			table[i][j] = table[i+1][j+1] + table[i+1][j]
		}
		return table[i][j]
	}
	for i := len(s); i >= 0; i-- {
		for j := len(t); j >= 0; j-- {
			table[i][j] = helper(i, j)
		}
	}
	fmt.Println(table)
	return table[0][0]
}
