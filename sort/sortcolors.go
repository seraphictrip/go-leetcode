package sort

/*
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.



Example 1:

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Example 2:

Input: nums = [2,0,1]
Output: [0,1,2]


Constraints:

n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.


Follow up: Could you come up with a one-pass algorithm using only constant extra space?

*/

func sortColors(nums []int) {
	// colors have been pre mapped, so I can use as is
	// 0, 1, and 2 to represent the color red, white, and blue, respectively.
	bijection := func(x int) int { return x }
	codomain := []int{0, 0, 0}

	BucketSort(nums, bijection, bijection, codomain)

}

// Bucket sort takes a list of inputs to be sorted and an
// injective function/one-to-one mapper
// the injective function maps distinct elements of its domain
//
//	to distinct elements of its codomain
//
// the range of the injective is equal to its codomain, which is to say
// that every element of the codomain is mapped to by some element of the domain
// this makes it surjective.  An injective that is also surjective is "bijective"
// the nice thing about something that is both injective and surjective (one-to-one and onto)
// is that it has an inverse f⁻¹ that maps codmain to domain
// in this implementation
// The codomain is the index of an array where we will keep frequency counts
// though we can change it if we provide the inverse mapper...
func BucketSort(domain []int, bijection func(input int) int, inverse func(input int) int, codomain []int) {
	for _, input := range domain {
		codomain[bijection(input)]++
	}

	i := 0
	for n := range codomain {
		for j := 0; j < codomain[n]; j++ {
			domain[i] = inverse(n)
			i++
		}
	}
}

// Bucket sort takes a list of inputs to be sorted and an
// injective function/one-to-one mapper
// the injective function maps distinct elements of its domain
//
//	to distinct elements of its codomain
//
// the range of the injective is equal to its codomain, which is to say
// that every element of the codomain is mapped to by some element of the domain
// this makes it surjective.  An injective that is also surjective is "bijective"
// the nice thing about something that is both injective and surjective (one-to-one and onto)
// is that it has an inverse f⁻¹ that maps codmain to domain
// in this implementation
// The codomain is an exhaustive list of the codomain, but this is more or less
// the constraint for when we choose to, or to not use bucket sort so is very fitting
func BucketSortGeneric[T any, V comparable](inputs []T, bijection func(input T) V, inverse func(input V) T, codomain []V) {
	// frequency count
	freq := make(map[V]int, len(codomain))
	// we could initialize, but 0 value is meaningful
	for _, input := range inputs {
		freq[bijection(input)]++
	}

	i := 0
	for _, y := range codomain {
		for j := 0; j < freq[y]; j++ {
			inputs[i] = inverse(y)
			i++
		}
	}
}
