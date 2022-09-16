package medium

import "sort"

// Link: https://leetcode.com/problems/find-original-array-from-doubled-array/
func findOriginalArray(changed []int) []int {
	sort.Ints(changed)
	res := make([]int, 0, len(changed)/2)
	tracked := make(map[int]int)
	for i := range changed {
		if changed[i]%2 == 0 && tracked[changed[i]>>1] > 0 {
			tracked[changed[i]>>1]--
			res = append(res, changed[i]>>1)
			continue
		}
		tracked[changed[i]]++
	}
	for _, v := range tracked {
		if v > 0 {
			return []int{}
		}
	}
	return res
}
