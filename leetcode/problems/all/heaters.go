package problems

import (
	"sort"

	"github.com/lovung/ds/maths"
)

func findRadius(houses []int, heaters []int) int {
	nearestDis := make([]int, len(houses))
	sort.SliceStable(heaters, func(i, j int) bool {
		return heaters[i] < heaters[j]
	})
	for i := 0; i < len(houses); i++ {
		nearestDis[i] = binarySearchNearestDistance(houses[i], heaters)
	}
	return maths.Max(nearestDis...)
}

func binarySearchNearestDistance(location int, sortedPoints []int) int {
	l, r := 0, len(sortedPoints)
	for {
		mid := (l + r) / 2
		dis := location - sortedPoints[mid]
		switch {
		case dis == 0:
			return 0
		case dis < 0:
			if mid == 0 {
				return -dis
			}
			if location > sortedPoints[mid-1] {
				return min(location-sortedPoints[mid-1], -dis)
			}
			r = mid - 1
		case dis > 0:
			if mid == len(sortedPoints)-1 {
				return dis
			}
			if location < sortedPoints[mid+1] {
				return maths.Min(sortedPoints[mid+1]-location, dis)
			}
			l = mid + 1
		}
	}
}
