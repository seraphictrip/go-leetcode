package dp

import "fmt"

/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed, the only constraint stopping you from robbing
each of them is that adjacent houses have security systems connected and it will automatically
contact the police if two adjacent houses were broken into on the same night.

- can not rob adjacent houses,

skip take


Given an integer array nums representing the amount of money of each house,
return the maximum amount of money you can rob tonight without alerting the police.
- cost/opportunity in nums array
*/

func Rob(nums []int) int {
	memo := make(map[int]int, len(nums))
	result := robStartingWithTopDown(0, nums, memo)

	return result
}

func robStartingWith(i int, nums []int) int {
	if i > len(nums)-1 {
		// ran off edge
		return 0
	}
	return max(nums[i]+robStartingWith(i+2, nums), robStartingWith(i+1, nums))
}

func robStartingWithTopDown(i int, nums []int, memo map[int]int) int {
	if i > len(nums)-1 {
		// ran off edge
		return 0
	}
	if m, ok := memo[i]; ok {
		return m
	}
	memo[i] = max(nums[i]+robStartingWithTopDown(i+2, nums, memo), robStartingWithTopDown(i+1, nums, memo))
	return memo[i]
}

// {[]int{1, 2, 3, 1}, 4}
// [0,0,1,2,4,4]

// {[]int{2, 7, 9, 3, 1}, 12}
// [0,0,2,7,11,11,12]
func RobBottomUp(nums []int) int {
	dp := make([]int, len(nums)+2)

	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-2])
	}
	fmt.Println(dp)
	return dp[len(dp)-1]
}
