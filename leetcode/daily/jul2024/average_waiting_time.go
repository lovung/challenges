package jul2024

// https://leetcode.com/problems/average-waiting-time/
func averageWaitingTime(customers [][]int) float64 {
	nextEndOfCurOrder, waitTimeSum := 0, 0
	for _, c := range customers {
		nextEndOfCurOrder = max(nextEndOfCurOrder, c[0])
		nextEndOfCurOrder += c[1]
		waitTimeSum += nextEndOfCurOrder - c[0]
	}
	return float64(waitTimeSum) / float64(len(customers))
}
