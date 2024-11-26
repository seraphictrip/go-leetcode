package lcs

/*
Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.

An interleaving of two strings s and t is a configuration where s and t are divided into n and m
substrings
 respectively, such that:

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
Note: a + b is the concatenation of strings a and b.



Example 1:


Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true
Explanation: One way to obtain s3 is:
Split s1 into s1 = "aa" + "bc" + "c", and s2 into s2 = "dbbc" + "a".
Interleaving the two splits, we get "aa" + "dbbc" + "bc" + "a" + "c" = "aadbbcbcac".
Since s3 can be obtained by interleaving s1 and s2, we return true.
Example 2:

Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false
Explanation: Notice how it is impossible to interleave s2 with any other string to obtain s3.
Example 3:

Input: s1 = "", s2 = "", s3 = ""
Output: true


Constraints:

0 <= s1.length, s2.length <= 100
0 <= s3.length <= 200
s1, s2, and s3 consist of lowercase English letters.


Follow up: Could you solve it using only O(s2.length) additional memory space?
*/

func IsInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	var helper func(i, j, k int) bool
	helper = func(i, j, k int) bool {
		if i == len(s1) {
			return s2[j:] == s3[k:]
		}
		if j == len(s2) {
			return s1[i:] == s3[k:]
		}
		if s1[i] == s3[k] && s2[j] == s3[k] {
			return helper(i+1, j, k+1) || helper(i, j+1, k+1)
		} else if s1[i] == s3[k] {
			return helper(i+1, j, k+1)
		} else if s2[j] == s3[k] {
			return helper(i, j+1, k+1)
		}
		return false
	}
	return helper(0, 0, 0)
}

func IsInterleaveMemo(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	memo := make(map[[2]int]bool)
	var helper func(i, j int) bool
	helper = func(i, j int) bool {
		// k is a derived variable, it is always just i+j, so we can remove
		k := i + j
		key := [2]int{i, j}
		if val, ok := memo[key]; ok {
			return val
		}
		if i == len(s1) {
			memo[key] = s2[j:] == s3[k:]
			return memo[key]
		}
		if j == len(s2) {
			memo[key] = s1[i:] == s3[k:]
			return memo[key]
		}
		if s1[i] == s3[k] && s2[j] == s3[k] {
			memo[key] = helper(i+1, j) || helper(i, j+1)
		} else if s1[i] == s3[k] {
			memo[key] = helper(i+1, j)
		} else if s2[j] == s3[k] {
			memo[key] = helper(i, j+1)
		} else {
			memo[key] = false
		}
		return memo[key]
	}
	return helper(0, 0)
}

func IsInterleaveBottomUp(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	table := make([][]bool, len(s1)+1)
	for i := range table {
		table[i] = make([]bool, len(s2)+1)
	}
	var helper func(i, j int) bool
	// ORDER:
	// i, j requires:
	// i+1 bigger before smaller
	// j+1 bigger before smaller
	helper = func(i, j int) bool {
		// k is a derived variable, it is always just i+j, so we can remove
		k := i + j
		// if val, ok := memo[key]; ok {
		// 	return val
		// }
		if i == len(s1) {
			table[i][j] = s2[j:] == s3[k:]
			return table[i][j]
		}
		if j == len(s2) {
			table[i][j] = s1[i:] == s3[k:]
			return table[i][j]
		}
		if s1[i] == s3[k] && s2[j] == s3[k] {
			table[i][j] = table[i+1][j] || table[i][j+1]
		} else if s1[i] == s3[k] {
			table[i][j] = table[i+1][j]
		} else if s2[j] == s3[k] {
			table[i][j] = table[i][j+1]
		} else {
			table[i][j] = false
		}
		return table[i][j]
	}

	for i := len(s1); i >= 0; i-- {
		for j := len(s2); j >= 0; j-- {
			table[i][j] = helper(i, j)
		}
	}

	return table[0][0]
}
