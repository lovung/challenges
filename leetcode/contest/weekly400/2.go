package weekly400

import "sort"

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	curMeeting := meetings[0]
	days -= (curMeeting[1] - curMeeting[0] + 1)
	for i := 1; i < len(meetings); i++ {
		// overlap
		if meetings[i][0] <= curMeeting[1] {
			// included
			if meetings[i][1] <= curMeeting[1] {
				continue
			}
			// gap
			days -= (meetings[i][1] - curMeeting[1])
			curMeeting[1] = meetings[i][1]
		} else {
			curMeeting = meetings[i]
			days -= (curMeeting[1] - curMeeting[0] + 1)
		}
	}
	return days
}
