package arrandhash

func TwoSumOnePass(nums []int, target int) []int {
	// num => index
	seen := make(map[int]int, len(nums))

	for i, num := range nums {
		if index, ok := seen[target-num]; ok {
			return []int{i, index}
		}
		seen[num] = i
	}
	return nil
}

func TwoSumHashMap(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, num := range nums {
		seen[num] = i
	}
	for i, num := range nums {
		// NOTE: this actually makes more difficult, have to check not self
		if index, ok := seen[target-num]; ok && i != index {
			return []int{i, index}
		}
	}
	return nil
}

// O(n^2)
func TwoSumBruteForce(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// we can use intution that previous pass would have already caught to not have to start at 0
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
