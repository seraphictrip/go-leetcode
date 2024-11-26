package recursion

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Factorial1(n int) int {
	result := 1
	for ; n > 0; n-- {
		return n * result
	}
	return result
}

func FactorialIterative(n int) uint64 {
	if n <= 1 {
		return 1
	}
	result := 1
	for i := n; i >= 1; i-- {
		result *= i
	}
	return uint64(result)
}
