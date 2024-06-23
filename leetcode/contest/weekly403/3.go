package weekly403

// https://leetcode.com/problems/maximize-total-cost-of-alternating-subarrays/description/
func maximumTotalCost(nums []int) int64 {
	var first, second = int64(nums[0]), int64(nums[0])
	for i := 1; i < len(nums); i++ {
		first, second = max(first, second)+int64(nums[i]), first-int64(nums[i])
	}
	return max(first, second)
}
