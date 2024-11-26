package bitmanipulation

func FindKOr(nums []int, k int) int {
	result := 0
	test := 1
	for i := 0; i < 32; i++ {
		count := 0
		for _, num := range nums {
			if test&num == test {
				count++
				if count >= k {
					result = result | test
					break
				}
			}
		}
		test = test << 1
	}
	return result
}
