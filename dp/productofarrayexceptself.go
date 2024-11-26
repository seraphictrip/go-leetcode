package dp

import (
	"fmt"
	"math"
)

/*
Given an integer array nums, return an array answer such that answer[i] is equal
to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]

[1,1,2,6]
[24,12,4,1]

Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]


Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.


Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
*/

func ProductExceptSelfNaive(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		val := 1
		for j := 0; j < n; j++ {
			if j != i {
				val *= nums[j]
			}
		}
		result[i] = val
	}

	return result
}

// use prefix/suffix products and then combine
func ProductExceptSelf(nums []int) []int {
	n := len(nums)

	if n == 0 {
		return nums
	}

	prefix := make([]int, n)
	suffix := make([]int, n)
	prefix[0] = 1
	suffix[n-1] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
		suffix[n-i-1] = suffix[n-i] * nums[n-i]
	}
	fmt.Println(prefix, suffix)
	return zip(prefix, suffix, func(x, y int) int { return x * y })

}

func zip[T any](a []T, b []T, with func(x, y T) T) []T {
	result := make([]T, len(a))
	for i := range a {
		result[i] = with(a[i], b[i])
	}
	return result
}

/*
Given an integer array nums, find the
subarray
 with the largest sum, and return its sum.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
[5] = 5
[5,4] = 9
[5,4, -1] = 8
[5,4,-1,7] = 15
[5,4,-1,7,8] = 23
[4] = 4
[4,-1] = 3
[4,-1,7] = 10
[4,-1,7,8] = 18
[-1] = -1
[-1,7] = 6
[-1,7,8] = 14
[7] = 7
[7,8] = 15
[8] = 8



Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.


Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
*/

func MaxSubArrayNaive(nums []int) int {
	// global maximum
	maxi := math.MinInt // setting to nums[0] also acceptable

	// outer loop for starting point
	for i := 0; i < len(nums); i++ {
		// runnign total, assuming take I
		currSum := 0
		// inner loop for ending point
		for j := i; j < len(nums); j++ {
			currSum += nums[j]
			// if we found a new max, store, otherwise keep processing
			maxi = max(maxi, currSum)
		}
	}
	return maxi
}

// [5,4,-1,7,8]
// INIT maxi = 5, maxEnding = 5
// i=1: maxEnding = 9, maxi = 9
// i=2: maxEnding = 8, maxi = 9
// i=3: maxEnding = 15, maxi = 15
// i=4: maxEnding = 23, maxi = 23

// [2, 3, -8, 7, -1, 2, 3]
// INIT: maxi = 2, maxEnding=2
// i=1: maxEnding=5, maxi=5
// i=2: maxEnding=-3, maxi=5
// i=3: maxEnding=7, maxi=7
// i=4: maxEnding=6, maxi=7
// i=5: maxEnding=8, maxi=8
// i=6: maxEnding=11, maxi=11
// Kadane's Algorithm: traverse over teh array from left to right
// at each element find the max subarry ENDING at each element
func MaxSubArray(nums []int) int {
	maxi := nums[0]
	maxEnding := nums[0]

	for i := 1; i < len(nums); i++ {
		// max ending is the running total of max ending, OR self
		maxEnding = max(maxEnding+nums[i], nums[i])
		// global max can be updated along the way
		maxi = max(maxi, maxEnding)
	}

	return maxi
}
