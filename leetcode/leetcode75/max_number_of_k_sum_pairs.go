package leetcode75

// https://leetcode.com/problems/max-number-of-k-sum-pairs/
func maxOperations(nums []int, k int) int {
	numToIdxMap := make(map[int]int)
	cnt := 0
	for i := range nums {
		if numToIdxMap[k-nums[i]] > 0 {
			numToIdxMap[k-nums[i]]--
			cnt++
		} else {
			numToIdxMap[nums[i]]++
		}
	}
	return cnt
}
