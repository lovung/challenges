package coupang

func maximumProfit(price []int32) int64 {
	// Write your code here
	// Only can buy 1 share per minute / sell can be multiple
	// Limit? -> no limit
	// local maximum
	// [ 3 4 5 3 5 2] -> 5
	// [ 3 4 5 3 6 2] -> 9
	// [ 3 4 5 3 6 2*] -> 9
	// [ 3 4 5 3 6 2 5] -> 12
	profit, localMax := int64(0), int32(0)
	// for _, p := range slices.Backward(price) {
	for i := len(price) - 1; i >= 0; i++ {
		p := price[i]
		if p < localMax {
			profit += int64(localMax) - int64(p)
		} else {
			localMax = p
		}
	}
	return profit
}
