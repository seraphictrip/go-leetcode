package lcs

/*
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character


Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')


Constraints:

0 <= word1.length, word2.length <= 500
word1 and word2 consist of lowercase English letters.
*/

func MinDistance(word1 string, word2 string) int {
	var helper func(i, j int) int

	helper = func(i, j int) int {
		if i == len(word1) {
			return len(word2) - j
		}
		if j == len(word2) {
			return len(word1) - i
		}
		if word1[i] == word2[j] {
			// char at position match so sovle subproblem
			return helper(i+1, j+1)
		}
		// if we insert, pretend we inserted word2[j] to match, word one still at char and j progresses
		insert := 1 + helper(i, j+1)
		// if we delete a character sub "remove from word" so need to test next i, but same j
		delete := 1 + helper(i+1, j)
		// if replace pretend we replaced char to match, so can move forward
		replace := 1 + helper(i+1, j+1)

		return min(insert, delete, replace)
	}
	return helper(0, 0)
}

func MinDistanceMemo(word1 string, word2 string) int {
	var helper func(i, j int) int
	memo := make([][]int, len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	helper = func(i, j int) int {
		if i == len(word1) {
			return len(word2) - j
		}
		if j == len(word2) {
			return len(word1) - i
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if word1[i] == word2[j] {
			// char at position match so sovle subproblem
			memo[i][j] = helper(i+1, j+1)
			return memo[i][j]
		}
		// if we insert, pretend we inserted word2[j] to match, word one still at char and j progresses
		insert := 1 + helper(i, j+1)
		// if we delete a character sub "remove from word" so need to test next i, but same j
		delete := 1 + helper(i+1, j)
		// if replace pretend we replaced char to match, so can move forward
		replace := 1 + helper(i+1, j+1)

		memo[i][j] = min(insert, delete, replace)
		return memo[i][j]
	}
	return helper(0, 0)
}

func MinDistanceBottomUp(word1 string, word2 string) int {
	var helper func(i, j int) int
	table := make([][]int, len(word1)+1)
	for i := range table {
		table[i] = make([]int, len(word2)+1)
	}

	// ORDER:
	// i, j require
	// i + 1 and j+1 bigger before smaller
	helper = func(i, j int) int {
		if i == len(word1) {
			return len(word2) - j
		}
		if j == len(word2) {
			return len(word1) - i
		}
		// if table[i][j] != -1 {
		// 	return table[i][j]
		// }
		if word1[i] == word2[j] {
			// char at position match so sovle subproblem
			table[i][j] = table[i+1][j+1]
			return table[i][j]
		}
		// if we insert, pretend we inserted word2[j] to match, word one still at char and j progresses
		insert := 1 + table[i][j+1]
		// if we delete a character sub "remove from word" so need to test next i, but same j
		delete := 1 + table[i+1][j]
		// if replace pretend we replaced char to match, so can move forward
		replace := 1 + table[i+1][j+1]

		table[i][j] = min(insert, delete, replace)
		return table[i][j]
	}

	for i := len(word1); i >= 0; i-- {
		for j := len(word2); j >= 0; j-- {
			table[i][j] = helper(i, j)
		}
	}

	return table[0][0]
}
