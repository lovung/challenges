package grind75

import "sort"

// Link: https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	sum /= k
	n := len(nums)
	visited := make([]bool, n)
	return recursiveCanPartitionKSubsets(nums, visited, 0, k, 0, sum)
}

func recursiveCanPartitionKSubsets(nums []int, visited []bool, startIdx, k, curSum, target int) bool {
	if k == 1 {
		return true
	}
	if curSum == target {
		return recursiveCanPartitionKSubsets(nums, visited, 0, k-1, 0, target)
	}
	for i := startIdx; i < len(nums); i++ {
		if !visited[i] {
			visited[i] = true
			if recursiveCanPartitionKSubsets(nums, visited, i+1, k, curSum+nums[i], target) {
				return true
			}
			visited[i] = false
		}
	}
	return false
}

func canPartitionKSubsets2(nums []int, k int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	sum /= k
	n := len(nums)
	dp := make([]bool, 1<<n)
	total := make([]int, 1<<n)
	dp[0] = true
	sort.Ints(nums)
	if nums[n-1] > sum {
		return false
	}
	for i := 0; i < (1 << n); i++ {
		if !dp[i] {
			continue
		}
		for j := 0; j < n; j++ {
			temp := i | 1<<j
			if temp != i {
				if nums[j] <= (sum - total[i]%sum) {
					dp[temp] = true
					total[temp] = nums[j] + total[i]
				} else {
					break
				}
			}
		}
	}
	return dp[(1<<n)-1]
}
