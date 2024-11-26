package subset

import "slices"

// At it's core combinations are just subsets of subsets
// Choose(n, k) N choose K
func Choose(n int, k int) [][]int {
	combinations := make([][]int, 0)
	combo := []int{}
	var helper func(int)
	helper = func(i int) {
		// halt early if reach k
		if len(combo) == k {
			combinations = append(combinations, slices.Clone(combo))
			return
		}
		// make sure to halt if further than n
		if i > n {
			return
		}

		combo = append(combo, i)
		helper(i + 1)

		combo = combo[:len(combo)-1]
		helper(i + 1)
	}
	helper(1)
	return combinations
}

func ChooseV2(n, k int) [][]int {
	// n!/(k!(n-k!))
	// 5*4*3*2*1
	// 2*1*3*2*1
	combinations := make([][]int, 0, n) // let dynamically grow, tho could do better if need
	combo := []int{}
	var helper func(int)
	helper = func(i int) {
		if len(combo) == k {
			combinations = append(combinations, slices.Clone(combo))
		}
		if i > n {
			return
		}
		for j := i; j < n; j++ {
			combo = append(combo, j)
			helper(j + 1)
			combo = combo[:len(combo)-1]
		}
	}
	helper(1)
	return combinations
}
