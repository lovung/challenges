package problems

// Link: https://leetcode.com/problems/teemo-attacking
// Input: [1, 4] dur: 2 -> Output: 4
// Input: [1, 2] dur: 2 -> Output: 3
func findPoisonedDuration(timeSeries []int, duration int) int {
	totalPoison := 0
	n := len(timeSeries)
	for i := 1; i < n; i++ {
		if timeSeries[i]-timeSeries[i-1] > duration {
			totalPoison += duration
		} else {
			totalPoison += timeSeries[i] - timeSeries[i-1]
		}
	}
	return totalPoison + duration
}
