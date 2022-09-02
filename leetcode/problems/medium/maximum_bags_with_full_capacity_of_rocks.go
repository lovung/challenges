package medium

import "sort"

// Link: https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/
type bag struct {
	cap  int
	rock int
	miss int
}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	n := len(capacity)
	bags := make([]*bag, 0, n)
	for i := range capacity {
		bags = append(bags, &bag{
			cap:  capacity[i],
			rock: rocks[i],
			miss: capacity[i] - rocks[i],
		})
	}
	sort.Slice(bags, func(i, j int) bool {
		return bags[i].miss < bags[j].miss
	})
	count := 0
	for i := range bags {
		if bags[i].miss <= additionalRocks {
			additionalRocks -= bags[i].miss
			count++
		}
	}
	return count
}
