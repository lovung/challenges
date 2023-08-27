package contest

import (
	"math"
)

func uniq(nums []int) []int {
	cntMap := make(map[int]int)
	for i := range nums {
		cntMap[nums[i]]++

	}
	values := make([]int, 0)
	for k := range cntMap {
		values = append(values, k)
	}
	return values
}

func findNearest2Way(nums []int, from int, ref int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[(from+i)%n] == ref {
			return i
		}
		if nums[(from-i+n)%n] == ref {
			return i
		}
	}
	return 0
}

// https://leetcode.com/problems/minimum-seconds-to-equalize-a-circular-array/
// O(N^3) solution
func minimumSeconds(nums []int) int {
	values := uniq(nums)
	minResult := math.MaxInt
	for _, v := range values {
		maxLen := 0
		for i := 0; i < len(nums); i++ {
			far := findNearest2Way(nums, i, v)
			if maxLen < far {
				maxLen = far
			}
		}
		if minResult > maxLen {
			minResult = maxLen
		}
	}
	return minResult
}
