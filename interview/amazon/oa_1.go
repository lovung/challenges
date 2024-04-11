package amazon

import "sort"

func reduceGifts(prices []int32, k int32, threshold int32) int32 {
	// Write your code here
	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})
	n := len(prices)
	sum := int32(0)
	// Calculate sum of the k-highest price items
	l := n - 1
	for l >= n-int(k) {
		sum += prices[l]
		l--
	}
	r := n - 1
	// Shift the siding window to the left
	for l >= 0 && sum > threshold {
		// Calculate new sum of k items between l and r-1
		sum += prices[l] - prices[r]
		l--
		r--
	}
	if l < 0 {
		return int32(n - r)
	}
	return int32(n - r - 1)
}
