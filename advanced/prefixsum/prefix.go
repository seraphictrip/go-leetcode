package prefixsum

func PrefixSumInclusive(nums []int) []int {
	prefix := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefix[i] = nums[i]
		} else {
			prefix[i] = prefix[i-1] + nums[i]
		}
	}

	return prefix
}

func PrefixSumExclusive(nums []int) []int {
	prefix := make([]int, len(nums))
	prefix[0] = 0
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}

	return prefix
}
