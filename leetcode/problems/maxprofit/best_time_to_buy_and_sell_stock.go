package maxprofit

// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	maxAfter := make([]int, len(prices))
	maxPrice := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if maxPrice < prices[i] {
			maxPrice = prices[i]
		}
		maxAfter[i] = maxPrice
	}
	maxProfit := 0
	for i := range prices {
		maxProfit = max(maxProfit, maxAfter[i]-prices[i])
	}
	return maxProfit
}
