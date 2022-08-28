package hard

import (
	"sort"
)

// Link: https://leetcode.com/problems/minimum-initial-energy-to-finish-tasks/
func minimumEffort(tasks [][]int) int {
	// the important thing is that the order of tasks is desc of difference
	// of minimum_i and actual_i
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] > tasks[j][1]-tasks[j][0]
	})
	sum := 0
	for i := len(tasks) - 1; i >= 0; i-- {
		// the sum is the energy before starting the tasks[i]
		// we go from the bottom, so, need to get the max of
		// (last energy + actual) vs minimum
		sum += tasks[i][0]
		if sum < tasks[i][1] {
			sum = tasks[i][1]
		}
	}
	return sum
}
