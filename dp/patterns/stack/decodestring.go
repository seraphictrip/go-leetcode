package stack

import (
	"strconv"
	"strings"
	"unicode"
)

/*
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 105.



Example 1:

Input: s = "3[a]2[bc]"

		3a	c
a			b
3			2

Output: "aaabcbc"
Example 2:

Input: s = "3[a2[c]]"
Output: "accaccacc"
Example 3:

Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"


Constraints:

1 <= s.length <= 30
s consists of lowercase English letters, digits, and square brackets '[]'.
s is guaranteed to be a valid input.
All the integers in s are in the range [1, 300].
*/

func DecodeString(s string) string {
	stack := NewStack(len(s))
	curNum := 0
	curString := strings.Builder{}
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			// upto 300, so have to account for 10s place
			digit := int(ch - '0')
			curNum = curNum*10 + digit
		} else if ch == '[' {
			// add string so far to stack
			stack.Push(curString.String())
			// add num to stack
			stack.Push(curNum)
			curNum = 0
			curString.Reset()
		} else if ch == ']' {
			num := stack.Pop().(int)
			appnd := curString.String()
			prefix := stack.Pop().(string)
			curString.Reset()
			curString.WriteString(prefix)
			for i := 0; i < num; i++ {
				curString.WriteString(appnd)
			}
		} else {
			// build string
			curString.WriteRune(ch)
		}

	}
	return curString.String()
}

type Stack []any

func NewStack(cap int) Stack {
	return make([]any, 0, cap)
}

func (s *Stack) Push(val any) {
	*s = append(*s, val)
}

func (s *Stack) Pop() any {
	old := *s
	topIndex := len(old) - 1
	top := old[topIndex]
	*s = old[:topIndex]
	return top
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func EvalRPN(tokens []string) int {
	stack := NewStack(len(tokens))
	for _, token := range tokens {
		switch token {
		case "+":
			right := stack.Pop().(int)
			left := stack.Pop().(int)
			stack.Push(left + right)
		case "-":
			right := stack.Pop().(int)
			left := stack.Pop().(int)
			stack.Push(left - right)
		case "*":
			right := stack.Pop().(int)
			left := stack.Pop().(int)
			stack.Push(left * right)
		case "/":
			right := stack.Pop().(int)
			left := stack.Pop().(int)
			stack.Push(left / right)
		default:
			digit, _ := strconv.Atoi(token)
			stack.Push(digit)
		}
	}
	return stack.Pop().(int)
}
