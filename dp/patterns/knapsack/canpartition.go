package knapsack

/*
total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	goal := total / 2
	memo := make([][]bool, len(nums))
    for i := 0; i < len(memo); i++ {
        memo[i] = make([]bool, goal+1)
    }

	var helper func(i, a int) bool
	helper = func(i, a int) bool {
		if a == goal {
			return true
		}
		if i == len(nums) {
			return false
		}
		if memo[i][a] {
			return true
		}

        take := false
        if a+nums[i] <= goal {
            take = helper(i+1, a+nums[i])
        }

		skip := helper(i+1, a)
		memo[i][a] = take || skip
		return memo[i][a]
	}
	return helper(0, 0)
*/

func canPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	goal := total / 2

	table := make([][]bool, len(nums))

	var helper func(i, goal int) bool

	helper = func(i, goal int) bool {
		if goal == 0 {
			return true
		}
		if i >= len(nums) {
			return false
		}

		take := false
		if goal-nums[i] >= 0 {
			take = table[i+1][goal-nums[i]]
		}
		skip := table[i+1][goal]
		table[i][goal] = take || skip
		return table[i][goal]
	}

	for g := 0; g <= goal; g++ {
		for i := len(nums); i >= 0; i-- {
			table[i][g] = helper(i, g)
		}
	}

	return table[0][goal]
}

/*
intutitive solution
 var helper func(i, a, b int) bool
    helper = func(i, a, b int) bool {
        if i == len(nums) {
            return a == b
        }

        left := helper(i+1, a+nums[i], b)
        right := helper(i+1, a, b+nums[i])
        return left || right
    }
    return helper(0, 0, 0)
*/

// This was an obvious solution I had a hard time turning into dp or even memoizing
// will try again
func CanPartitionIntuitive(nums []int) bool {
	var helper func(i, a, b int) bool
	helper = func(i, a, b int) bool {
		if i == len(nums) {
			return a == b
		}

		left := helper(i+1, a+nums[i], b)
		right := helper(i+1, a, b+nums[i])
		return left || right
	}
	return helper(0, 0, 0)
}

func CanPartitionIntuitiveMemo(nums []int) bool {
	memo := make(map[[3]int]bool, len(nums))
	var helper func(i, a, b int) bool
	helper = func(i, a, b int) bool {
		if i == len(nums) {
			return a == b
		}
		key := [3]int{i, a, b}
		if val, ok := memo[key]; ok {
			return val
		}
		left := helper(i+1, a+nums[i], b)
		right := helper(i+1, a, b+nums[i])
		memo[key] = left || right
		return memo[key]
	}
	return helper(0, 0, 0)
}

func CanPartitionKnapsack(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	goal := total / 2
	var helper func(i, goal int) bool
	helper = func(i, goal int) bool {
		if goal == 0 {
			return true
		}
		if i >= len(nums) || goal < 0 {
			return false
		}
		take := false

		if goal-nums[i] >= 0 {
			take = helper(i+1, goal-nums[i])
		}

		skip := helper(i+1, goal)

		return take || skip
	}
	return helper(0, goal)
}

func CanPartitionKnapsackMemo(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	goal := total / 2
	memo := make(map[[2]int]bool, goal)
	var helper func(i, goal int) bool
	helper = func(i, goal int) bool {
		if goal == 0 {
			return true
		}
		if i >= len(nums) || goal < 0 {
			return false
		}
		key := [2]int{i, goal}
		if val, ok := memo[key]; ok {
			return val
		}
		take := false

		if goal-nums[i] >= 0 {
			take = helper(i+1, goal-nums[i])
		}

		skip := helper(i+1, goal)

		memo[key] = take || skip
		return memo[key]
	}
	ans := helper(0, goal)

	return ans
}

func CanPartitionHashSet(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	goal := total / 2
	set := make(map[int]bool, len(nums))
	setKeys := make([]int, 1, len(nums))
	setKeys[0] = 0
	set[0] = true
	for i := 0; i < len(nums); i++ {
		for _, key := range setKeys {
			newKey := key + nums[i]
			if newKey == goal {
				return true
			}
			if !set[newKey] {
				set[newKey] = true
				setKeys = append(setKeys, newKey)
			}
		}
	}
	return false
}
