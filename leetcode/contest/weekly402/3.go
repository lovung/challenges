package weekly402

import (
	"slices"
)

func maximumTotalDamage(power []int) int64 {
	pm := make(map[int]int)
	exist := make([]int, 0, len(power))
	for _, p := range power {
		if pm[p] == 0 {
			exist = append(exist, p)
		}
		pm[p]++
	}
	slices.SortFunc(exist, func(a, b int) int { return b - a })
	res := int64(0)
	take := make(map[int]int64)
	dont := make(map[int]int64)
	for _, e := range exist {
		// calculate take
		if pm[e+2] > 0 {
			take[e] = dont[e+2] + int64(e*pm[e])
			if pm[e+1] > 0 {
				take[e] = min(take[e], dont[e+1]+int64(e*pm[e]))
			}
		} else {
			if pm[e+1] > 0 {
				take[e] = dont[e+1] + int64(e*pm[e])
			} else {
				take[e] = res + int64(e*pm[e])
			}
		}
		// calculate dont
		dont[e] = res
		if pm[e+2] > 0 {
			dont[e] = max(res, take[e+2], dont[e+2])
		}
		if pm[e+1] > 0 {
			dont[e] = max(res, take[e+1], dont[e+1])
		}
		// calculate res
		res = max(take[e], dont[e])
	}
	return res
}
