package oct2022

import (
	"sort"

	"github.com/lovung/ds/trees"
)

// Link: https://leetcode.com/problems/russian-doll-envelopes/
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})
	dp := make([]int, len(envelopes))
	maxVal := 1
	segTree := trees.NewSegmentTreeWithArray(len(envelopes), maxFn)

	for i := 0; i < len(envelopes); i++ {
		maxSmaller := 0
		for j := i - 1; j >= 0; j-- {
			if envelopes[i][0] > envelopes[j][0] &&
				envelopes[i][1] > envelopes[j][1] {
				maxSmaller = max(maxSmaller, dp[j])
				if maxSmaller > segTree.Query(0, j) {
					break
				}
			}
		}
		if maxSmaller > 0 {
			dp[i] = maxSmaller + 1
			maxVal = max(maxVal, dp[i])
		} else {
			dp[i] = 1
		}
		segTree.Update(i, dp[i])
	}
	return maxVal
}

func maxFn(a, b int) int {
	return max(a, b)
}

func maxEnvelopes2(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})
	dp := make([]int, len(envelopes))
	size := 0

	for i := range envelopes {
		l, r := 0, size
		for l < r {
			m := (l + r) / 2
			if dp[m] < envelopes[i][1] {
				l = m + 1
			} else {
				r = m
			}
		}

		dp[l] = envelopes[i][1]
		if l == size {
			size++
		}
	}
	return size
}
