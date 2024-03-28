package coupang

import (
	"math"
	"slices"
)

func findLIS3(nums []int) int {
	sub := make([]int, 0, len(nums))
	for _, x := range nums {
		if len(sub) == 0 || sub[len(sub)-1] < x {
			sub = append(sub, x)
		} else {
			idx, found := slices.BinarySearch(sub, x)
			if found {
				continue
			}
			sub[idx] = x
		}
	}
	return len(sub)
}

func findLIS2(nums []int) int {
	dp := make([]int, len(nums))
	size := 0
	for i := range nums {
		l, r := 0, size
		for l < r {
			m := (l + r) / 2
			if dp[m] < nums[i] {
				l = m + 1
			} else {
				r = m
			}
		}

		dp[l] = nums[i]
		if l == size {
			size++
		}
	}
	return size
}

func findLIS(input []int) int {
	m := make(map[int]int)
	m[math.MinInt] = 0
	for i := range input {
		for maxVal, length := range m {
			if input[i] > maxVal {
				m[input[i]] = max(m[input[i]], length+1)
			}
		}
	}
	ret := 0
	for _, length := range m {
		ret = max(ret, length)
	}
	return ret
}
