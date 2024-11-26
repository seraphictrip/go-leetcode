package dp

/*
You are given an integer array coins representing coins of different
denominations and an integer amount representing a total amount of money.
ex: [1,2,5] (penny, 2cent, 5cent)
- len(coins) small (<300)
- DOMAIN: [1..5000]
- amount can be zero? what does that mean, seems like can make 0 an inf number of ways by jsut skipping reppeating

Return the number of combinations that make up that amount.
If that amount of money cannot be made up by any combination of the coins,
return 0.
- count


you may assume that you have an infinite number of each kind of coin.

The answer is guaranteed to fit into a signed 32-bit integer.


SUMMARY:
Using an array of coins, come up with all combinations that sum to amount


Input: amount = 5, coins = [1,2,5]
Output: 4
take: count(amount-n, coins)
skip: count(amount, coins[1:])



recurrence realtion
effected parameters
amount: [0, amount]
i: [0, len(coins)]

n * amounts is number of unique states

formula for every dp problem is just the number of unique sates * cached complexity (O(1) usually, i.e. assume)
*/

func Change(amount int, coins []int) int {
	// I actually change coins, so this isn't good enough, need something more...
	return change(amount, coins)
}

// to do a memoized version I need a memo of [len(coins)+1][amount], we build a pretty big table
// TODO: come back to this

func change(amount int, coins []int) int {
	if amount < 0 {
		// this is not a valid way
		return 0
	}
	if amount == 0 {
		// this counts as a valid way
		return 1
	}
	if len(coins) == 0 {
		// we ran out of coins, this is also terminal if we haven't
		// hit a different base case
		return 0
	}
	// at each step I can with take or skip
	// take implies I remove amount from amount
	return change(amount-coins[0], coins) + change(amount, coins[1:])
}

/*
Laws of recursion
1. Recursive call is correct
2. All recursive calls must lead to base case
3. All base cases must be correct
*/
