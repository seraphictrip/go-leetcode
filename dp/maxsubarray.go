package dp

import "math"

/*
Given an integer array nums, find a
subarray
 that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.



Example 1:

Input: nums = [2,3,-2,4]
[2] = 2
[2,3] = 6
[2,3,-2] = -12

Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.


Constraints:

1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
The product of any subarray of nums is guaranteed to fit in a 32-bit intege4

We are doing multiplication so 1 is Identity (as opposed to 0)

can either take or skip, if skip can not carry forward

*/

func MaxProduct(nums []int) int {
	n := len(nums)
	maxi := math.MinInt
	prefix, suffix := 1, 1

	for i := 0; i < n; i++ {
		// if we zero'd out, we need to start fresh
		if prefix == 0 {
			prefix = 1
		}
		if suffix == 0 {
			suffix = 1
		}
		prefix = prefix * nums[i]
		suffix = suffix * nums[n-i-1]

		maxi = max(maxi, max(prefix, suffix))

	}
	return maxi
}
