package jun2024

// https://leetcode.com/problems/continuous-subarray-sum/
func checkSubarraySum(nums []int, k int) bool {
	modMap := make(map[int]int)
	modMap[0] = -1
	prefixSum := 0
	for i := range nums {
		prefixSum += nums[i]
		if val, ok := modMap[prefixSum%k]; ok {
			if val < i-1 {
				return true
			}
		} else {
			modMap[prefixSum%k] = i
		}
	}
	return false
}
