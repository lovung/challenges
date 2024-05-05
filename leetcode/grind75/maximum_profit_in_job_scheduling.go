package grind75

import (
	"sort"
)

// Link: https://leetcode.com/problems/maximum-profit-in-job-scheduling
type job struct {
	start, end int
	profit     int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)

	jobs := make([]*job, 0, n)
	for i := range startTime {
		jobs = append(jobs, &job{
			start:  startTime[i],
			end:    endTime[i],
			profit: profit[i],
		})
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})
	dp := make([]int, n)
	dp[0] = jobs[0].profit
	for i := 1; i < n; i++ {
		prof := jobs[i].profit
		l := binarySearchJobs(jobs, i)
		if l != -1 {
			prof += dp[l]
		}
		dp[i] = max(prof, dp[i-1])
	}
	return dp[n-1]
}

func binarySearchJobs(jobs []*job, index int) int {
	l, r := 0, index-1
	for l <= r {
		mid := (l + r) / 2
		if jobs[mid].end <= jobs[index].start {
			if jobs[mid+1].end <= jobs[index].start {
				l = mid + 1
			} else {
				return mid
			}
		} else {
			r = mid - 1
		}
	}
	return -1
}
