package medium

import (
	"github.com/lovung/ds/queue"
)

// Link: https://leetcode.com/problems/avoid-flood-in-the-city/

// More cases:
// [1, 2, 0, 2, 0, 1] -> [-1, -1, 2, -1, 1, -1]
// [1, 2, 0, 2, 0, 0, 1] -> [-1, -1, 2, -1, 1, 1/2, -1]

func avoidFlood(rains []int) []int {
	// We have the map to save the full of water lakes
	// Value is the index of rain day
	fullOfWaterLakes := make(map[int]int)
	// Queue to save the index of the day which we can dry
	// to get it when we got the rain into the full of water lake
	dryDays := queue.NewSimpleQueue[int]()

	res := make([]int, len(rains))

	for i := range rains {
		if rains[i] > 0 {
			if lastRainDay, full := fullOfWaterLakes[rains[i]]; full {
				if dryDays.Len() == 0 {
					return []int{}
				}
				// try to find the suitable dryDay in the queue
				found := false
				length := dryDays.Len()
				for j := 0; j < length; j++ {
					canDryDay, _ := dryDays.DeQueue()
					// canDryDay should be after the lastRainDay
					if canDryDay > lastRainDay {
						found = true
						res[canDryDay] = rains[i]
						delete(fullOfWaterLakes, rains[i])
						break
					} else {
						dryDays.EnQueue(canDryDay)
					}
				}
				if !found {
					return []int{}
				}
			} else {
				fullOfWaterLakes[rains[i]] = i
			}
			res[i] = -1
		} else {
			if len(fullOfWaterLakes) == 1 {
				for k := range fullOfWaterLakes {
					res[i] = k
					delete(fullOfWaterLakes, k)
				}
				continue
			}
			dryDays.EnQueue(i)
		}
	}
	for i := range res {
		if res[i] == 0 {
			res[i] = 1
		}
	}
	return res
}
