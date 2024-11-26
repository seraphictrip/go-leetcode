package twopointers

/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original
string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).



Example 1:

Input:
s = "abc"
	   j
t = "ahbgdc"
	      i

Output: true
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false


Constraints:

0 <= s.length <= 100
0 <= t.length <= 104
s and t consist only of lowercase English letters.
*/

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for ; i < len(t); i++ {
		if t[i] == s[j] {
			j++
			if j == len(s) {
				// found all of s
				return true
			}
		}
	}
	return false
}

/*
s = "abc"

	i

t = "ahbgdc"

	j
*/
func IsSubsequenceBruteForce(s, t string) bool {
	j := 0
	for i := 0; i < len(t); i++ {
		if j >= len(s) {
			return true
		}
		if t[i] == s[j] {
			j++
		}
	}
	return j >= len(s)
}
