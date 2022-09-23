package grind75

import "sort"

// Link: https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			o := nums[l] + nums[i] + nums[r]
			switch {
			case o == 0:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
			case o < 0:
				l++
			default:
				r--
			}
		}
	}
	return res
}
