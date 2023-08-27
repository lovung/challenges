package aug2023

import (
	"sort"
)

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r := 0, nums[n-1]-nums[0]
	for l < r {
		mid := (l + r) / 2
		if countValidPairs(nums, mid) >= p {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func countValidPairs(nums []int, threshold int) int {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] <= threshold {
			cnt++
			i++
		}
	}
	return cnt
}
