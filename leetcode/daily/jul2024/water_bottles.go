package jul2024

// https://leetcode.com/problems/water-bottles/
func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	for numBottles >= numExchange {
		res += numBottles / numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	return res
}
