package grind75

import (
	"math"
	"sort"

	"github.com/lovung/ds/maths"
)

// Link: https://leetcode.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	// timeout
	return dfsCoinChange(coins, amount, 0)
}

func dfsCoinChange(remainCoins []int, remainAmount int, cnt int) int {
	maxDeno := remainCoins[len(remainCoins)-1]
	if remainAmount%maxDeno == 0 {
		return cnt + remainAmount/maxDeno
	}
	if len(remainCoins) <= 1 {
		return -1
	}
	min := -1
	for i := 0; i <= remainAmount/maxDeno; i++ {
		val := dfsCoinChange(
			remainCoins[:len(remainCoins)-1],
			remainAmount-i*maxDeno,
			cnt+i,
		)
		if val > 0 {
			if min == -1 || min > val {
				min = val
			}
		}
	}
	return min
}

func coinChange2(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt64
		for _, c := range coins {
			if i < c {
				break
			}
			if dp[i-c] != math.MaxInt64 {
				dp[i] = maths.Min(dp[i], 1+dp[i-c])
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}
