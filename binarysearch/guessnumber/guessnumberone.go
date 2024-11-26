package guessnumber

import "fmt"

/*
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API int guess(int num), which returns three possible results:

-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.



Example 1:

Input: n = 10, pick = 6
Output: 6
Example 2:

Input: n = 1, pick = 1
Output: 1
Example 3:

Input: n = 2, pick = 1
Output: 1


Constraints:

1 <= n <= 231 - 1
1 <= pick <= n
*/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * Domain ~Z+
 * Codomain: [-1, 0, 1]
 * Range -1 to 1
 * func guess(num int) int;
 */

// Observations: guess is a cmp function (think sort or parition)
//
//		[< Pick][Pick][>Pick] // this style is sort/inorder as opposed to say
//	  preorder [PICK][<PICK][>PICK]
//		postorder [<PICK][>PICK][PICK]
//
// Domain: ~Z+ (positive ints)
// Codomain: ~Z+
// Domain == Codomain
func GuessNumber(n int, guess func(g int) int) int {
	l := 1
	r := n
	count := 1
	for l <= n {
		mid := (l + r) / 2
		fmt.Printf("%v. %v of %v\n", count, mid, n)
		count++
		g := guess(mid)
		if g == 0 {
			return mid
		}
		if g < 0 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func MakeGuess(pick int) func(g int) int {
	return func(g int) int {
		if pick < g {
			return -1
		}
		if pick > g {
			return 1
		}
		return 0
	}
}
