package apr2024

import (
	"math"

	"github.com/lovung/ds/maths"
)

type step struct {
	pos, cost int
}

const a = 'a'

// https://leetcode.com/problems/freedom-trail/description/
func findRotateSteps(ring string, key string) int {
	n := len(ring)
	pos := [26][]int{}
	for i := range ring {
		pos[ring[i]-a] = append(pos[ring[i]-a], i)
	}
	lastSteps := []step{{0, 0}}
	for _, k := range key {
		newPossitions := pos[k-a]
		costMap := make(map[int]int)
		for _, s := range lastSteps {
			for _, newPos := range newPossitions {
				abs := maths.ABS(newPos - s.pos)
				newCost := s.cost + min(abs, n-abs) + 1 // press button
				if costMap[newPos] == 0 {
					costMap[newPos] = newCost
				} else {
					costMap[newPos] = min(costMap[newPos], newCost)
				}
			}
		}
		lastSteps = summarySteps(costMap)
	}
	res := math.MaxInt
	for i := range lastSteps {
		res = min(res, lastSteps[i].cost)
	}
	return res
}

func summarySteps(costs map[int]int) []step {
	steps := make([]step, 0, 26)
	for idx, cost := range costs {
		steps = append(steps, step{idx, cost})
	}
	return steps
}
