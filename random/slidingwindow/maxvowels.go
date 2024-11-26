package slidingwindow

/*
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.


* Sliding window
* fixed size
* count vowels

Example 1:

Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.
Example 2:

Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.
Example 3:

Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.


Constraints:

1 <= s.length <= 105
s consists of lowercase English letters.
1 <= k <= s.length
*/

func MaxVowels(s string, k int) int {
	result := 0
	count := 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}
	result = count
	L, R := 1, k
	for R < len(s) {
		if isVowel(s[L-1]) {
			count--
		}
		if isVowel(s[R]) {
			count++
		}
		result = max(result, count)
		L++
		R++
	}

	return result
}

func isVowel(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
