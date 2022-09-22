package containsduplicatie

// Link: https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	mark := make(map[int]bool)
	for i := range nums {
		if mark[nums[i]] {
			return true
		}
		mark[nums[i]] = true
	}
	return false
}
