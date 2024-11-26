package prefixsum

import (
	"sync"
	"sync/atomic"
)

/*
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.



Example 1:

Input: nums = [1,1,1], k = 2
prefix [1,2,3]
suffix [3,2,1]
sum: 3
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
prefix [1,3,6]
sum 6
Output: 2


Constraints:

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/

func SubarraySumBruteForce(nums []int, k int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		cur := 0
		for j := i; j < len(nums); j++ {
			cur += nums[j]
			if cur == k {
				result++
			}
		}
	}

	return result
}

// Intutition: what prefix sums, if any can we chop off to reach goal
func SubarraySum(nums []int, k int) int {
	res := 0
	curSum := 0
	prefixSums := make(map[int]int, len(nums))
	prefixSums[0] = 1
	// this is just calculating sum forward, so curSum would be prefixes[i] if we had precalculated
	for _, num := range nums {
		curSum = curSum + num
		if cnt, ok := prefixSums[curSum-k]; ok {
			res += cnt
		}
		prefixSums[curSum]++

	}
	return res
}

func SubarrySumPrefix(nums []int, k int) int {
	res := 0
	prefixes := prefixsum(nums)
	seen := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		if cnt, ok := seen[prefixes[i]-k]; ok {
			res += cnt
		}
		seen[prefixes[i]]++
	}
	return res

}

func prefixsum(nums []int) []int {
	prefix := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			prefix[i] = num
		} else {
			prefix[i] = prefix[i-1] + num
		}
	}
	return prefix
}

func SubarraySumParalellWaitGroupAndAtomic(nums []int, k int) int {
	var result atomic.Int32
	wg := sync.WaitGroup{}
	wg.Add(len(nums))
	for i := 0; i < len(nums); i++ {
		go func(i int) {
			cur := 0
			for j := i; j < len(nums); j++ {
				cur += nums[j]
				if cur == k {
					result.Add(1)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	return int(result.Load())
}

func SubarraySumParalellMutex(nums []int, k int) int {
	result := 0
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	wg.Add(len(nums))
	for i := 0; i < len(nums); i++ {
		go func(i int) {
			cur := 0
			for j := i; j < len(nums); j++ {
				cur += nums[j]
				if cur == k {
					// using a mutex to surround code that modifies similar to using atomic
					mutex.Lock()
					result++
					mutex.Unlock()
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	return result
}

func SubarraySumParalellChannel(nums []int, k int) int {
	result := 0
	ch := make(chan int)
	for i := 0; i < len(nums); i++ {
		go func(i int) {
			cur := 0
			for j := i; j < len(nums); j++ {
				cur += nums[j]
				if cur == k {
					ch <- 1
				}
			}
		}(i)
	}

	close(ch)
	for i := range ch {
		result += i
	}

	return result
}
