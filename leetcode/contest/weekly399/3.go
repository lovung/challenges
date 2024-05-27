package weekly399

import (
	"slices"
	"sort"
)

// https://leetcode.com/problems/find-the-number-of-good-pairs-ii/
func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	res := int64(0)
	sort.Ints(nums1)
	for _, n2 := range nums2 {
		kn2 := n2 * k
		i, _ := slices.BinarySearch(nums1, kn2)
		for ; i < len(nums1); i++ {
			if nums1[i]%kn2 == 0 {
				res++
			}
		}
	}
	return res
}

func numberOfPairs2(nums1 []int, nums2 []int, k int) int64 {
	ans := int64(0)
	freq := make(map[int]int)
	for _, n := range nums2 {
		freq[n*k]++
	}
	for _, n := range nums1 {
		for i := 1; i*i <= n; i++ {
			if n%i != 0 {
				continue
			}
			ans += int64(freq[i])
			if i != n/i {
				ans += int64(freq[n/i])
			}
		}
	}
	return ans
}
