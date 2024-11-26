package subset

import (
	"math"
	"slices"
)

// [1,2,3]
// 								*
//				/							\
//			[]									[1]
//			/	\								/ 		 \
//		[]		[2]							[1]    		[1,2]
//		/\		/	\					/	\			/		\
//	 []  [3]  [2]   [2,3]			[1]		[1,3]	[1,2]		[1,2,3]

func PowerSet(nums []int) [][]int {
	powerset := make([][]int, 0, int(math.Pow(2, float64(len(nums)))))
	powerset = helper(0, nums, []int{}, powerset)
	return powerset
}

func PowerSet2(set []int) [][]int {
	powerset := make([][]int, 0, int(math.Pow(2, float64(len(set)))))
	subset := []int{}
	var helper func(i int)
	helper = func(i int) {
		if i >= len(set) {
			powerset = append(powerset, slices.Clone(subset))
			return
		}
		subset = append(subset, set[i])
		helper(i + 1)
		subset = subset[:len(subset)-1]
		helper(i + 1)
	}
	helper(0)
	return powerset
}

func helper(i int, set, subset []int, powerset [][]int) [][]int {
	if i >= len(set) {
		powerset = append(powerset, slices.Clone(subset))
		return powerset
	}
	// include
	subset = append(subset, set[i])
	powerset = helper(i+1, set, subset, powerset)
	// exclude
	subset = subset[:len(subset)-1]
	powerset = helper(i+1, set, subset, powerset)
	return powerset
}

func PowerMultiSet(set []int) [][]int {
	// we have to sort, so we can properly detect skips
	slices.Sort(set)
	powerset := make([][]int, 0, int(math.Pow(2, float64(len(set)))))
	subset := []int{}
	var helper func(i int)
	helper = func(i int) {
		if i >= len(set) {
			powerset = append(powerset, slices.Clone(subset))
			return
		}
		subset = append(subset, set[i])
		helper(i + 1)
		subset = subset[:len(subset)-1]

		// when we decide to skip, skip all instances, capture will happen in alternate branch
		for i+1 < len(set) && set[i] == set[i+1] {
			i++
		}
		helper(i + 1)
	}
	helper(0)
	return powerset
}
