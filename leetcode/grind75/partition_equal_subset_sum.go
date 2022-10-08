package grind75

// Link: https://leetcode.com/problems/partition-equal-subset-sum/
func canPartition(nums []int) bool {
	return canPartitionKSubsets(nums, 2)
}
