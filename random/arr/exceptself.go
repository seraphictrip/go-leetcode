package arr

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]


Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.


Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)


*/

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = 1
	suffix := make([]int, n)
	suffix[n-1] = 1
	for i := 1; i < n; i++ {
		j := n - i - 1
		prefix[i] = prefix[i-1] * nums[i-1]
		suffix[j] = suffix[j+1] * nums[j+1]
	}
	// for j := n - 2; j >= 0; j-- {
	// 	suffix[j] = suffix[j+1] * nums[j+1]
	// }
	for i := 0; i < n; i++ {
		prefix[i] = prefix[i] * suffix[i]
	}
	return prefix
}
