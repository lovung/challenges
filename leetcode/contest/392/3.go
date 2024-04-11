package contest392

import "sort"

func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)
	mid := len(nums) / 2
	if nums[mid] < k {
		return cntToRight(nums, mid, k)
	}
	return cntToLeft(nums, mid, k)
}

func cntToRight(nums []int, idx, k int) int64 {
	res := int64(0)
	for i := idx; i < len(nums); i++ {
		if nums[i] < k {
			res += int64(k - nums[i])
		} else {
			break
		}
	}
	return res
}

func cntToLeft(nums []int, idx, k int) int64 {
	res := int64(0)
	for i := idx; i >= 0; i-- {
		if nums[i] > k {
			res += int64(nums[i] - k)
		} else {
			break
		}
	}
	return res
}
