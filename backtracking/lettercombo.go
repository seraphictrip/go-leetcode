package backtracking

import "fmt"

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.




Example 1:

Input: digits = "23"
					a
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Example 2:

Input: digits = ""
Output: []
Example 3:

Input: digits = "2"
				a b c
Output: ["a","b","c"]


Constraints:

0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].
*/
// DOMAIN: 2-9
// CODOMAIN: a-z
func LetterCombinations(digits string) []string {
	result := make([]string, 0)

	acc := ""
	var backtrack func(int)

	backtrack = func(start int) {
		if start >= len(digits) {
			if len(acc) == len(digits) {
				result = append(result, acc)
			}
			return
		}
		for i := start; i < len(digits); i++ {
			ms := getMappings(digits[i])
			for j := 0; j < len(ms); j++ {
				acc = fmt.Sprintf("%v%v", acc, string(ms[j]))
				backtrack(i + 1)
				acc = acc[:len(acc)-1]
			}
		}

	}
	backtrack(0)
	return result
}

func getMappings(digit byte) string {
	i := int(digit - '0')
	return mappings[i]
}

var mappings = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
