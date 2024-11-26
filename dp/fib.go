package dp

// Recursive/Brute force solution
//
//					(5)
//		(4)					(3)
//	 (2)  	(3)  		(1)		(2)
//
// (1) (0) (2) (1)			(1)		(0)
//
// F(0) = 0, F(1) = 1
// F(n) = F(n-1) + F(n-2)
func Fib1(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib1(n-1) + Fib1(n-2)
}

func FibMemo(n int) int {
	memo := map[int]int{
		0: 0,
		1: 1,
	}
	return fibMemo(n, memo)
}

func fibMemo(n int, memo map[int]int) int {
	// check cache
	if cached, ok := memo[n]; ok {
		return cached
	}
	// cachce before return
	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}

// Bottom up, we build the table of values from the beginning
func FibIter(n int) int {
	if n <= 1 {
		return n
	}
	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		tmp := dp[1]
		dp[1] = dp[0] + dp[1]
		dp[0] = tmp
	}
	return dp[1]
}

func FibBottomUp(n int) int {
	if n <= 1 {
		return n
	}
	// version above is "better", but this should be more illustrative hopefully?
	table := make([]int, n+1)
	table[0] = 0
	table[1] = 1
	for i := 2; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n]
}

// Dynamic Programming

// Start with a decision Tree

// Identify sub problems

// https://blog.moertel.com/posts/2013-05-11-recursive-to-iterative.html
/*
The Simple Method
This translation method works on many simple recursive functions. When it works, it works well, and the results are lean and fast. I generally try it first and consider more complicated methods only when it fails.

In a nutshell:

1. Study the function.
2. Convert all recursive calls into tail calls. (If you can’t, stop. Try another method.)
3. Introduce a one-shot loop around the function body.
4. Convert tail calls into continue statements.
5. Tidy up.
An important property of this method is that it’s incrementally correct – after every step you have a function that’s equivalent to the original. So if you have unit tests, you can run them after each and every step to make sure you didn’t make a mistake.

Let’s see the method in action.

*/

// http://en.wikipedia.org/wiki/Mathematical_induction

// TOP DOWN = MEMOIZATION

// BOTTOM UP (usually iterative)

// 1. Study original
func factorial(n int) int {
	if n < 2 {
		return n
	}
	return n * factorial(n-1)
}

// 2. convert to tail call
// we doing multiplication, so in this instance acc should defualt to 1
func factorial1a(n, acc int) int {
	if n < 2 {
		return 1 * acc
	}
	return factorial1a(n-1, n*acc)
}

// 3. Introduce a one-shot loop around the function body.
func factorial1b(n, acc int) int {
	if acc == 0 {
		acc = 1
	}

	for {
		if n < 2 {
			return 1 * acc
		}
		return factorial1a(n-1, n*acc)
		break
	}
	return acc
}

// 4. Replace all recursive tail calls f(x=x1, y=y1, ...) with (x, y, ...) = (x1, y1, ...); continue. Be sure to update all arguments in the assignment
func factorial1c(n, acc int) int {
	if acc == 0 {
		acc = 1
	}

	for {
		if n < 2 {
			return 1 * acc
		}
		n, acc = n-1, acc*n
		continue
	}
}

func factorial1d(n, acc int) int {
	for n > 1 {
		n, acc = n-1, acc*n
	}
	return acc
}

/*
1. Find a recursive call that’s not a tail call.
2. Identify what work is being done between that call and its return statement.
3. Extend the function with a secret feature to do that work, as controlled by a new accumulator argument with a default value that causes it to do nothing.
4. Use the secret feature to eliminate the old work.
5. You’ve now got a tail call!
6. Repeat until all recursive calls are tail calls.
*/

func binomial(n, k int) int {
	if k == 0 {
		return 1
	}
	// x := binomial(n-1, k-1)
	// n * x / k
	return n * binomial(n-1, k-1) / k
}

func step(x, lmul, rdiv int) int {
	return lmul * x / rdiv
}
