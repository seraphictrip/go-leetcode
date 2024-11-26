package lcs

import "fmt"

// Q: Given two strings s1, s2, find the length of the longest common subsequence between the two strings

// This was me just solving...
func LCSBruteForce(s1, s2 string) int {
	subseq := make(map[string]bool)
	subseq[""] = true
	current := make([]string, 1, 10)
	current[0] = ""
	// find all subsequendes starting at i
	for i := len(s1) - 1; i >= 0; i-- {
		for _, cur := range current {
			next := fmt.Sprintf("%v%s", string(s1[i]), cur)
			if !subseq[next] {
				subseq[next] = true
				current = append(current, next)
			}
		}
	}

	lcs := ""
	lcslen := 0
	current = []string{""}
	for i := 0; i < len(s2); i++ {
		for _, cur := range current {
			next := fmt.Sprintf("%v%s", cur, string(s2[i]))
			if subseq[next] {
				if len(next) > lcslen {
					lcslen = len(next)
					lcs = next
					current = append(current, lcs)
				}
			}
		}
	}

	fmt.Println(current, subseq)
	return lcslen
}

func LCSBruteForceRecursion(s1, s2 string) int {
	var helper func(i, j int) int
	helper = func(i, j int) int {
		if i == len(s1) || j == len(s2) {
			return 0
		}
		if s1[i] == s2[j] {
			return 1 + helper(i+1, j+1)
		}
		return max(helper(i+1, j), helper(i, j+1))
	}

	return helper(0, 0)
}

func LCSMemo(s1, s2 string) int {
	var helper func(i, j int) int
	memo := make([][]int, len(s1))
	for i := range memo {
		memo[i] = make([]int, len(s2)+1)
		for j := 0; j <= len(s2); j++ {
			memo[i][j] = -1
		}
	}
	helper = func(i, j int) int {
		if i == len(s1) || j == len(s2) {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if s1[i] == s2[j] {
			memo[i][j] = 1 + helper(i+1, j+1)
			return memo[i][j]
		}
		memo[i][j] = max(helper(i+1, j), helper(i, j+1))
		return memo[i][j]
	}

	ans := helper(0, 0)
	fmt.Println(memo)
	return ans
}

func LCSBottomUp(s1, s2 string) int {
	var helper func(i, j int) int
	table := make([][]int, len(s1)+1)
	for i := range table {
		table[i] = make([]int, len(s2)+1)
	}
	helper = func(i, j int) int {
		if i == len(s1) || j == len(s2) {
			return 0
		}
		if s1[i] == s2[j] {
			table[i][j] = 1 + table[i+1][j+1]
			return table[i][j]
		}
		table[i][j] = max(table[i+1][j], table[i][j+1])
		return table[i][j]
	}

	for i := len(s1) - 1; i >= 0; i-- {
		for j := len(s2) - 1; j >= 0; j-- {
			table[i][j] = helper(i, j)
		}
	}
	return table[0][0]
}
