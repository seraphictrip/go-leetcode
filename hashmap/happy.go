package hashmap

import (
	"fmt"
)

func IsHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 {
		if seen[n] {
			return false
		}
		seen[n] = true
		digits := GetDigits(n)
		n = 0
		for _, d := range digits {
			n += d * d
		}
		fmt.Println(n)
	}
	return true
}

func sum(digits []int) int {
	result := 0
	for _, num := range digits {
		result += num
	}
	return result
}

func GetDigits(n int) []int {
	digits := make([]int, 0, 3)
	cur := n
	for cur != 0 {
		last := cur % 10
		digits = append(digits, last)
		cur /= 10
	}
	return digits
}
