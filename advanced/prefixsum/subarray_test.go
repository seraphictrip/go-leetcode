package prefixsum_test

import (
	"dsa/advanced/prefixsum"
	"strconv"
	"testing"
)

var SubarrayTests = []struct {
	nums     []int
	k        int
	expected int
}{
	// | prefixSums			| curSum	| curSum-k	| res |
	// |{0: 1}				| 0			| -2		| 0   |
	// | {0: 1}				| 1			| -1		| 0	  |
	// | {0: 1, 1: 1}		| 2			| 0			| 1	  |
	// | {0: 1, 1: 1, 2: 1} | 3			| 1			| 1	  |

	{[]int{1, 1, 1}, 2, 2},
	{[]int{1, 2, 3}, 3, 2},
	{[]int{1, -1, 1, 1, 1, 1}, 3, 4},
	// [1,2], [1,2,3,-3]
	// [2,3,-3,1]
	// [3], [3,-3, 1, 1, 1]
	// [-3, 1, 1, 1]
	// [1, 1, 1]
	// -
	// -
	// [4, 2, -3]
	// -
	// -
	{[]int{1, 2, 3, -3, 1, 1, 1, 4, 2, -3}, 3, 8},
}

func TestSubarray(t *testing.T) {
	for i, e := range SubarrayTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := prefixsum.SubarraySumParalellChannel(e.nums, e.k)
			if actual != e.expected {
				t.Fatalf("SubarraySum(%v, %v) = %v, want %v", e.nums, e.k, actual, e.expected)
			}
		})
	}
}
