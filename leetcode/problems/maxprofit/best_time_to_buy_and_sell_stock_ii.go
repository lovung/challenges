package maxprofit

import "github.com/lovung/ds/maths"

// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfitII(prices []int) int {
	cash, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		cash = maths.Max(cash, hold+prices[i])
		hold = maths.Max(hold, cash-prices[i])
	}
	return cash
}

func maxProfitII2(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}
