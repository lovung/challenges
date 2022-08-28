package medium

import "github.com/lovung/ds/maths"

// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
func maxProfit(prices []int, fee int) int {
	cash, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		cash = maths.Max(cash, hold+prices[i]-fee)
		hold = maths.Max(hold, cash-prices[i])
	}
	return cash
}
