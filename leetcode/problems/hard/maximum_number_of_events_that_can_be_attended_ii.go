package hard

import (
	"sort"

	"github.com/lovung/ds/maths"
)

// Link: https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/
/*
	You are given an array of events where events[i] = [startDay_i, endDay_i, value_i].
	The ith event starts at startDay_i and ends at endDay_i, and if you attend this event,
	you will receive a value of value_i.

	! You are also given an integer k which represents the maximum number of events you can attend.
	You can only attend one event at a time.
	! If you choose to attend an event, you must attend the entire event.
	Note that the end day is inclusive: that is, you cannot attend two events
	where one of them starts and the other ends on the same day.

	* Return the maximum sum of values that you can receive by attending events.

	Constraints:

	* 1 <= k <= events.length
	* 1 <= k * events.length <= 10e6
	* 1 <= startDay_i <= endDay_i <= 10e9
	* 1 <= value_i <= 10e6
*/
func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] < events[j][0] {
			return true
		} else if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return false
	})
	n := len(events)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	return maxValueDpSolve(dp, events, n, 0, k)
}

func maxValueDpSolve(dp [][]int, events [][]int, n, pos, k int) int {
	if pos >= n || k == 0 {
		return 0
	}
	if dp[pos][k] != 0 {
		return dp[pos][k]
	}
	i := 0
	for i = pos + 1; i < n; i++ {
		if events[i][0] > events[pos][1] {
			break
		}
	}
	dp[pos][k] = maths.Max(maxValueDpSolve(dp, events, n, pos+1, k),
		events[pos][2]+maxValueDpSolve(dp, events, n, i, k-1))
	return dp[pos][k]
}
