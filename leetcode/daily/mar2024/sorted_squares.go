package mar2024

import "sort"

// https://leetcode.com/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	for i := range nums {
		ret[i] = nums[i] * nums[i]
	}
	sort.Ints(ret)
	return ret
}
