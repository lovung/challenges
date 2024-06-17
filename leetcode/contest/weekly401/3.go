package weekly401

import "sort"

// https://leetcode.com/problems/maximum-total-reward-using-operations-i/description/
func maxTotalReward1(rewardValues []int) int {
	sort.Ints(rewardValues)
	maxVal := rewardValues[len(rewardValues)-1]
	mark := make([]bool, maxVal*2)
	mark[0] = true
	res := 0
	for _, v := range rewardValues {
		for i := range v {
			if mark[i] {
				mark[i+v] = true
				res = max(res, i+v)
			}
		}
	}
	return res
}
