package jul2024

import "slices"

// https://leetcode.com/problems/sort-an-array/description/
func sortArray(nums []int) []int {
	cnt := make([]int, 1e5+1)
	const low = -5 * 1e4
	for _, num := range nums {
		cnt[num-low]++
	}
	res := make([]int, 0, len(nums))
	for i := range cnt {
		for range cnt[i] {
			res = append(res, i+low)
		}
	}
	return res
}

func sortArray2(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	slices.Sort(res)
	return res
}
