package jun2024

import (
	"sort"

	"github.com/lovung/ds/maths"
)

// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone
func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	res := 0
	for i := range seats {
		res += maths.ABS(seats[i] - students[i])
	}
	return res
}
