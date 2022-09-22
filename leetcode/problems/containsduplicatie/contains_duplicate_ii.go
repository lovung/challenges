package containsduplicatie

// Link: https://leetcode.com/problems/contains-duplicate-ii/
func containsNearbyDuplicate(nums []int, k int) bool {
	lastIdx := make(map[int]int)
	for i := range nums {
		if idx, ok := lastIdx[nums[i]]; ok && idx >= i-k {
			return true
		}
		lastIdx[nums[i]] = i
	}
	return false
}
