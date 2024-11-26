package prefixsum

import (
	"dsa/patterns/functional"
	"slices"
)

// https://www.youtube.com/watch?v=DjYZk8nrXVY
/*
leetcode problems:
303 Range Sum Query
525: Contiguous array
560: Subarray Sum Equals K
*/

// INTUTIION: If I need to calculate running sums or parts of array
// a good first step is calculating prefix sum

// Single Query
// O(n) n = j-i
// Multiple Queries
// O(n*m): n length of arr, m is number queries
func FindSubArraySum(arr []int, i, j int) int {
	sum := 0
	for k := i; k <= j; k++ {
		sum += arr[k]
	}
	return sum
}

// P[i] = A[0]+A[1]+..A[N]
// SUM[i,j] = P[j] - P[i-1] (do some tests for this, this is eureka I don't yet have)
func PrefixSum(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	output := make([]int, len(arr))
	prefix := 0
	for i := 0; i < len(arr); i++ {
		output[i] += prefix + arr[i]
		prefix = output[i]
	}
	return output
}

func PrefixSum2(arr []int) []int {
	prefix := 0
	return functional.Map(arr, func(n int) int {
		prefix = n + prefix
		return prefix
	})
}

// Sum from right, useful if finding arr[0:i] + arr[i+1:])
func PostfixSum(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	output := make([]int, len(arr))
	post := 0
	for i := len(arr) - 1; i >= 0; i-- {
		output[i] = arr[i] + post
		post = output[i]
	}

	return output
}

func PostfixSum2(arr []int) []int {
	cp := make([]int, len(arr))
	copy(cp, arr)
	slices.Reverse(cp)
	inverted := PrefixSum(cp)
	slices.Reverse(inverted)
	return inverted
}

// orig: [1,2,3,4,5,6,7]
// P: [1,3,6,10,15,21,28] =
// Sum[i,j] = P
func SumFromPrefix(prefixes []int, i, j int) int {
	total := prefixes[j]
	// we are summing from i to j, so we can take out anything before i
	discount := 0
	if i > 0 {
		discount = prefixes[i-1]
	}
	return total - discount
}

// ...4186 4278 4371 4465 4560 4656 4753 4851 4950]
// ...4999999250000028 4999999350000021 4999999450000015 4999999550000010 4999999650000006 4999999750000003 4999999850000001 4999999950000000 5000000050000000]
