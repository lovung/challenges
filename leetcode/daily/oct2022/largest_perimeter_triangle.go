package oct2022

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for k := n - 1; k >= 2; k-- {
		if nums[k-2]+nums[k-1] > nums[k] {
			return nums[k-2] + nums[k-1] + nums[k]
		}
		continue
	}
	return 0
}
