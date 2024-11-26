package arr

import "math"

/*
Given an integer array nums, return true if there exists a triple of indices (i, j, k)
such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.



Example 1:

Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
Example 2:

Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
Example 3:

Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.


Constraints:

1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1


Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?
*/

func IncreasingTripletBruteForce(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	// nums[i] < nums[j] < nums[k]

	// Brute force
	for i := 0; i <= n-3; i++ {
		for j := i + 1; j <= n-2; j++ {
			if nums[j] > nums[i] {
				for k := j + 1; k < n; k++ {
					if nums[k] > nums[j] {
						return true
					}
				}
			}

		}
	}
	return false
}

// Single pass.. but I only got here cheating, how to get here>?
func IncreasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	// nums[i] < nums[j] < nums[k]

	// x < y < x
	x := math.MaxInt
	y := math.MaxInt

	for i := 0; i < n; i++ {
		if nums[i] < x {
			x = nums[i]
		} else if nums[i] < y && nums[i] > x {
			y = nums[i]
		} else if nums[i] > y {
			return true
		}
	}

	return false
}

// Time: O(n) (3n)
// Space: O(n)
func IncreasingTripletPrefix(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	mins := make([]int, n)
	maxes := make([]int, n)

	localmin := math.MaxInt
	localmax := math.MinInt
	for i := 0; i < n; i++ {
		// calculate left min
		localmin = min(localmin, nums[i])
		mins[i] = localmin
		// calc right max
		j := n - i - 1
		localmax = max(localmax, nums[j])
		maxes[j] = localmax
	}

	for i := 1; i < n; i++ {
		if mins[i] < nums[i] && nums[i] < maxes[i] {
			return true
		}
	}

	return false
}
