package arr

import (
	"regexp"
	"slices"
	"strings"
)

/*
Input: s = "IceCreAm"
			L	  R
		   "AceCreIm"
		      L  R

Output: "AceCreIm"

Explanation:

The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".

Example 2:

Input: s = "leetcode"

Output: "leotcede"



Constraints:

1 <= s.length <= 3 * 105
s consist of printable ASCII characters.
*/

func ReverseVowels(s string) string {
	L, R := 0, len(s)-1
	result := s

	for L < R {
		for !isVowel(s[L]) {
			L++
		}
		for !isVowel(s[R]) {
			R--
		}
		if L < R {
			result = swap(result, L, R)
			L++
			R--
		}
	}
	return result
}

func isVowel(ch byte) bool {
	switch ch {
	case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
		return true
	default:
		return false
	}
}
func swap(str string, i, j int) string {
	bytes := []byte(str)
	bytes[i], bytes[j] = bytes[j], bytes[i]
	return string(bytes)
}

func ReverseWords(s string) string {
	re := regexp.MustCompile(`\w+`)
	words := re.FindAllString(s, -1)
	slices.Reverse(words)
	return strings.Join(words, " ")
}
