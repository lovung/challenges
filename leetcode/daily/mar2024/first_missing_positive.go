package mar2024

import (
	"slices"
	"sort"
)

// https://leetcode.com/problems/first-missing-positive
func firstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= 0 && nums[i]+1 < nums[i+1] {
			return nums[i] + 1
		}
	}
	return nums[len(nums)-1] + 1
}

func firstMissingPositive1(nums []int) int {
	nums = append(nums, 0)
	sort.Ints(nums)
	i, _ := slices.BinarySearch(nums, 0)
	for ; i < len(nums)-1; i++ {
		if nums[i]+1 < nums[i+1] {
			return nums[i] + 1
		}
	}
	return nums[len(nums)-1] + 1
}

func firstMissingPositive2(nums []int) int {
	n := len(nums)
	for i := range nums {
		idx := nums[i] - 1
		for 0 < nums[i] && nums[i] < n && nums[i] != nums[idx] {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx = nums[i] - 1
		}
	}
	for i := 1; i < n+1; i++ {
		if nums[i-1] != i {
			return i
		}
	}
	return n + 1
}
