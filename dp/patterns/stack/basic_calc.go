package stack

import (
	"unicode"
)

/*
Given a string s which represents an expression, evaluate this expression and return its value.

The integer division should truncate toward zero.

You may assume that the given expression is always valid. All intermediate results will be in the range of [-231, 231 - 1].

Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().



Example 1:

Input: s = "3+2*2"
Output: 7
Example 2:

Input: s = " 3/2 "
Output: 1
Example 3:

Input: s = " 3+5 / 2 "
Output: 5


Constraints:

1 <= s.length <= 3 * 105
s consists of integers and operators ('+', '-', '*', '/') separated by some number of spaces.
s represents a valid expression.
All the integers in the expression are non-negative integers in the range [0, 231 - 1].
The answer is guaranteed to fit in a 32-bit integer.
*/

func Calculate(s string) int {
	vals := NewGenericStack[int](10)
	operators := NewGenericStack[byte](10)
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			// build out our digit
			n := 0
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				n = n*10 + (int(s[i] - '0'))
				i++
			}
			i--
			vals.Push(n)
		} else if s[i] == '+' || s[i] == '-' {
			// ensure order of operation by evaluating everything to left first
			for !operators.IsEmpty() {
				op := operators.Pop()
				right := vals.Pop()
				left := vals.Pop()
				evaluated := eval(left, right, op)
				vals.Push(evaluated)
			}
			operators.Push(s[i])

		} else if s[i] == '*' || s[i] == '/' {
			// ensure evaluated left to right by doing any thing of equal precedence to left
			for !operators.IsEmpty() && (operators.Peek() == '*' || operators.Peek() == '/') {
				op := operators.Pop()
				right := vals.Pop()
				left := vals.Pop()
				evaluated := eval(left, right, op)
				vals.Push(evaluated)
			}
			operators.Push(s[i])
		}

	}
	for !operators.IsEmpty() {
		op := operators.Pop()
		right := vals.Pop()
		left := vals.Pop()
		evaluated := eval(left, right, op)
		vals.Push(evaluated)
	}
	return vals.Pop()
}

func eval(left int, right int, op byte) int {
	evaluated := 0
	switch op {
	case '+':
		evaluated = left + right
	case '-':
		evaluated = left - right
	case '*':
		evaluated = left * right
	case '/':
		evaluated = left / right
	}
	return evaluated
}

type GenericStack[T any] []T

func NewGenericStack[T any](cap int) GenericStack[T] {
	return make([]T, 0, cap)
}

func (s *GenericStack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *GenericStack[T]) Pop() T {
	old := *s
	topIndex := len(old) - 1
	top := old[topIndex]
	*s = old[:topIndex]
	return top
}

func (s GenericStack[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s GenericStack[T]) Peek() T {
	topIndex := len(s) - 1
	return s[topIndex]
}
