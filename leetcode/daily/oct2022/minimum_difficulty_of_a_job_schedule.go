package oct2022

import (
	"math"

	"github.com/lovung/ds/trees"
)

// Link: https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/
func minDifficulty_segmentTree(jobDifficulty []int, d int) int {
	if len(jobDifficulty) < d {
		return -1
	}
	st := trees.NewSegmentTreeWithArray(len(jobDifficulty), func(a, b int) int {
		return max(a, b)
	})
	for i, d := range jobDifficulty {
		st.Update(i, d)
	}
	return minDifficultyRecursive(st, jobDifficulty, 0, len(jobDifficulty), d)
}

func minDifficultyRecursive(st trees.SegmentTree[int], jobDifficulty []int, s, e, d int) int {
	if d == 1 {
		return st.Query(s, e)
	}
	res := math.MaxInt
	for i := s + 1; i <= e-d+1; i++ {
		res = min(res, st.Query(s, i)+minDifficultyRecursive(st, jobDifficulty, i, e, d-1))
	}
	return res
}

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}
	st := trees.NewSegmentTreeWithArray(n, func(a, b int) int {
		return max(a, b)
	})
	for i, d := range jobDifficulty {
		st.Update(i, d)
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = st.Query(0, i+1)
	}
	for i := 2; i <= d; i++ {
		newDp := make([]int, n)
		for j := i - 1; j < n; j++ {
			newDp[j] = math.MaxInt
			for k := i - 2; k < j; k++ {
				newDp[j] = min(newDp[j], dp[k]+st.Query(k+1, j+1))
			}
		}
		copy(dp, newDp)
	}
	return dp[n-1]
}
