package hard

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
	return 0
}
