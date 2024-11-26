package twopointers

import (
	"fmt"
	"math"
)

func CanJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	L, R := 0, 0
	for L <= R {
		if L+nums[L] > R {
			R = L + nums[L]
			if R >= len(nums)-1 {
				return true
			}
		}
		L++
	}
	return false

}

func Jump(nums []int) int {
	var helper func(i int) int
	goal := len(nums) - 1
	table := make([]int, len(nums))
	table[goal] = 0

	// ORDER:
	// i requires i + j, big before small
	helper = func(i int) int {
		if i >= goal {
			return 0
		}
		if i+nums[i] >= goal {
			return 1
		}
		minjumps := math.MaxInt / 2
		for j := 1; j <= nums[i]; j++ {
			count := 1 + table[i+j]
			if count < minjumps {
				minjumps = count
			}
		}
		table[i] = minjumps
		return table[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		table[i] = helper(i)
	}
	fmt.Println(table)
	return table[0]
}
