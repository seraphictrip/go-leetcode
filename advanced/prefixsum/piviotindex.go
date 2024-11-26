package prefixsum

func PivotIndex(nums []int) int {
	n := len(nums)
	prefix := make([]int, n)
	postfix := make([]int, n)
	prefix[0] = nums[0]
	postfix[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		j := n - i - 1
		prefix[i] = prefix[i-1] + nums[i]
		postfix[j] = postfix[j+1] + nums[j]
	}
	for i := 0; i < n; i++ {
		if prefix[i] == postfix[i] {
			return i
		}
	}
	return -1
}

func sum(arr []int) int {
	acc := 0
	for i := 0; i < len(arr); i++ {
		acc += arr[i]
	}
	return acc
}

func PivotIndexBruteForce(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if sum(nums[:i]) == sum(nums[i+1:]) {
			return i
		}
	}
	return -1
}

// NOTE: this is based off brute force, thus name, BUT
// is O(n), and O(1) memory, so more or less optimal
func PivotIndexBruteForce1(nums []int) int {
	left := 0
	// O(n) get complete sum, we will modify this as we move through array
	right := sum(nums)

	for i := 0; i < len(nums); i++ {
		// sum(left) == sum(right)
		if left == right-nums[i] {
			return i
		}
		// update left and right
		// left includes me for next round(s)
		left = left + nums[i]
		// right excludes me for next round(s)
		right = right - nums[i]
	}
	return -1
}
