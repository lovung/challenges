package hard

import "strconv"

// Link: https://leetcode.com/problems/two-best-non-overlapping-events
func maxTwoEvents(events [][]int) int {
	eventMap := make(map[string]struct{})
	newEvents := make([][]int, 0)
	for i := range events {
		eventKey := buildEventKey(events[i])
		if _, ok := eventMap[eventKey]; !ok {
			newEvents = append(newEvents, events[i])
			eventMap[eventKey] = struct{}{}
		}
	}

	return maxValue(newEvents, 2)
}

func buildEventKey(event []int) string {
	s := ""
	for i := range event {
		s += strconv.Itoa(event[i]) + "_"
	}
	return s
}
