package guessnumber

import "math"

/*
We are playing the Guessing Game. The game will work as follows:

I pick a number between 1 and n.
You guess a number.
If you guess the right number, you win the game.
If you guess the wrong number, then I will tell you whether the number I picked is higher or lower, and you will continue guessing.
Every time you guess a wrong number x, you will pay x dollars. If you run out of money, you lose the game.
Given a particular n, return the minimum amount of money you need to guarantee a win regardless of what number I pick.



Example 1:
//						(7)
//			(3)					(9)
//		(1)		    (5)		(8)		(10)
//			(2)	  (4) (6)

[
	[]
	[]
	[]
	[]
	[]
	[]
	[]
	[]
	[]
	[]
]

Input: n = 10
Output: 16
Explanation: The winning strategy is as follows:
- The range is [1,10]. Guess 7.
    - If this is my number, your total is $0. Otherwise, you pay $7.
    - If my number is higher, the range is [8,10]. Guess 9.
        - If this is my number, your total is $7. Otherwise, you pay $9.
        - If my number is higher, it must be 10. Guess 10. Your total is $7 + $9 = $16.
        - If my number is lower, it must be 8. Guess 8. Your total is $7 + $9 = $16.
    - If my number is lower, the range is [1,6]. Guess 3.
        - If this is my number, your total is $7. Otherwise, you pay $3.
        - If my number is higher, the range is [4,6]. Guess 5.
            - If this is my number, your total is $7 + $3 = $10. Otherwise, you pay $5.
            - If my number is higher, it must be 6. Guess 6. Your total is $7 + $3 + $5 = $15.
            - If my number is lower, it must be 4. Guess 4. Your total is $7 + $3 + $5 = $15.
        - If my number is lower, the range is [1,2]. Guess 1.
            - If this is my number, your total is $7 + $3 = $10. Otherwise, you pay $1.
            - If my number is higher, it must be 2. Guess 2. Your total is $7 + $3 + $1 = $11.
The worst case in all these scenarios is that you pay $16. Hence, you only need $16 to guarantee a win.
Example 2:

Input: n = 1
// INITIAL TABLE:
[
[0, 0]
[0, 0]
]
Output: 0
Explanation: There is only one possible number, so you can guess 1 and not have to pay anything.
Example 3:

Input: n = 2
// INITIAL MEMO:
[
 [0, 0, 0]
 [0, 0, 0]
 [0, 0, 0]
]
1. CalcCost(1, 2, memo) mid := 3/2 = 1
Output: 1
Explanation: There are two possible numbers, 1 and 2.
- Guess 1.
    - If this is my number, your total is $0. Otherwise, you pay $1.
    - If my number is higher, it must be 2. Guess 2. Your total is $1.
The worst case is that you pay $1.

*/

// Optimization game (dp?)
// playing the game will match binary search, but need to track ~minCost~ no, maxCost?, I need to stay in the game
// cost func F(x) = x; x member of [1...n]
// assuming I use Binary Search I would make log2(n)+1 guesses in worst case, can I use this?
//
//	The tree they provide guess 3 when I would have guessed 4, so maybe not?
func GetMoneyAmount(n int) int {
	memo := make([][]int, n+1)

	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	// we need to calc between 1 and n
	return calcCost(1, n, memo)
}

func calcCost(l, r int, memo [][]int) int {
	if l >= r {
		// if left equal right, we only have one guess
		// so cost is still 0
		return 0
	}

	// If we have already cached the answer, wew don't have to do it again
	// after I get this working, lets pull back out and do brute force, should be mechanical
	// but if I can do it in reverse hopefully it will help me do it forward
	if memo[l][r] != 0 {
		// we've already calculated this
		return memo[l][r]
	}
	minCost := math.MaxInt
	for guess := (l + r) / 2; guess <= r; guess++ {
		cost := guess + max(calcCost(l, guess-1, memo), calcCost(guess+1, r, memo))
		minCost = min(minCost, cost)
	}
	memo[l][r] = minCost
	return minCost
}
