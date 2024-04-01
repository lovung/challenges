package mar2024

import "slices"

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array
func searchRange(nums []int, target int) []int {
	start, found := slices.BinarySearch(nums, target)
	if !found {
		return []int{-1, -1}
	}
	end := start
	for end+1 < len(nums) && nums[end+1] == target {
		end++
	}
	return []int{start, end}
}
