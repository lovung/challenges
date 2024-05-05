package weekly395

import (
	"math"
	"sort"
)

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := math.MaxInt
	for i := 0; i <= 2; i++ {
		d := nums2[len(nums2)-1] - nums1[len(nums1)-1-i]
		if check(nums1, nums2, d, 2-i) {
			res = min(res, d)
		}
	}
	return res
}

func check(nums1, nums2 []int, delta, limit int) bool {
	wrongCnt := 0
	j := 0
	for i := range nums1 {
		if nums1[i]+delta != nums2[j] {
			wrongCnt++
			if wrongCnt > limit {
				return false
			}
		} else {
			j++
			if j >= len(nums2) {
				return true
			}
		}
	}
	return true
}
