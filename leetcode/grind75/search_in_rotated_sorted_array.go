package grind75

import "math"

// Link: https://leetcode.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) / 2
		var num int
		if (nums[mid] < nums[0]) == (target < nums[0]) {
			num = nums[mid]
		} else {
			if target < nums[0] {
				num = math.MinInt64
			} else {
				num = math.MaxInt64
			}
		}

		switch {
		case num < target:
			lo = mid + 1
		case num > target:
			hi = mid
		default:
			return mid
		}
	}
	return -1
}

// Link: https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
func search2(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (hi + lo) / 2
		if nums[mid] == target {
			return true
		}
		switch {
		case nums[lo] == nums[mid] && nums[hi] == nums[mid]:
			lo++
			hi--
		case nums[lo] <= nums[mid]:
			if nums[lo] <= target && nums[mid] > target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		default:
			if nums[mid] < target && nums[hi] >= target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return false
}
