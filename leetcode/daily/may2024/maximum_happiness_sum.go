package may2024

import "sort"

// https://leetcode.com/problems/maximize-happiness-of-selected-children/
func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	res := int64(0)
	cnt := 0
	for i := len(happiness) - 1; cnt < k; i-- {
		point := int64(happiness[i] - cnt)
		if point <= 0 {
			break
		}
		res += point
		cnt++
	}
	return res
}
