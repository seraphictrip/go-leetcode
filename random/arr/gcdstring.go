package arr

/*
For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.



Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""


Constraints:

1 <= str1.length, str2.length <= 1000
str1 and str2 consist of English uppercase letters.
*/

func GcdOfStrings(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	if str1+str2 != str2+str1 {
		return ""
	}
	if m > n {
		m, n = n, m
		str1, str2 = str2, str1
	}

	for i := m; i >= 0; i-- {
		candidate := str1[:i]
		if m%i == 0 && n%i == 0 {
			return candidate
		}
	}
	return ""
}

func Times(str string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += str
	}
	return result
}

func GCD(a, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}
