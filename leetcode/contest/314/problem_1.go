package contest

import "github.com/lovung/ds/maths"

// Link: https://leetcode.com/problems/the-employee-that-worked-on-the-longest-task
func hardestWorker(n int, logs [][]int) int {
	resID := n
	longestTaskDur := 0
	prevStart := 0
	for i := range logs {
		dur := logs[i][1] - prevStart
		if dur > longestTaskDur {
			resID = logs[i][0]
			longestTaskDur = dur
		} else if dur == longestTaskDur {
			resID = maths.Min(resID, logs[i][0])
		}
		prevStart = logs[i][1]
	}
	return resID
}
