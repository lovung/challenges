package may2024

import (
	"sort"
)

func specialArray(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	prev := 0
	for i := range nums {
		if (n-i) <= nums[i] && prev < (n-i) {
			return n - i
		}
		prev = nums[i]
	}
	return -1
}
