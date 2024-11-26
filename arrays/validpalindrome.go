package arrays

import "unicode"

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

func isPalindrome(s string) bool {
	L, R := 0, len(s)
	for L < R {
		for !isValidChar(s[L]) {
			L++
		}
		for !isValidChar(s[R]) {
			R--
		}
		if unicode.ToLower(rune(s[L])) != unicode.ToLower(rune(s[R])) {
			return false
		}
	}
	return true
}

func isValidChar(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || unicode.IsDigit(rune(ch))
}

/*
Low latency
security
scalability
consistency
availability

*/
