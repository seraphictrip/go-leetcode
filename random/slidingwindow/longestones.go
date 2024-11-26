package slidingwindow

/*
Given a binary array nums and an integer k,
return the maximum number of consecutive 1's in the array if you can flip at most k 0's.



Example 1:

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
Example 2:

Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.


Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.
0 <= k <= nums.length

*/

// sliding window
// keep at most flippable 0s in window
// if get more, then move L until back in range

func LongestOnes(nums []int, k int) int {
	result := 0
	flippable := 0
	L := 0
	for R := 0; R < len(nums); R++ {
		if nums[R] == 0 {
			flippable++
			for flippable > k {
				if nums[L] == 0 {
					flippable--
				}
				L++
			}
		}
		result = max(result, R-L+1)
	}

	return result
}

func LongestOnes1(nums []int, k int) int {
	result := 0
	flippable := 0

	L, R := 0, 0

	for R < len(nums) {
		if nums[R] == 0 {
			flippable++
			for flippable > k {
				if nums[L] == 0 {
					flippable--
				}
				L++
			}
		}
		result = max(result, R-L+1)

		R++
	}

	return result
}
