package jun2024

// https://leetcode.com/problems/patching-array/
func minPatches(nums []int, n int) int {
	miss := int64(1)
	result, i := 0, 0
	for miss <= int64(n) {
		if i < len(nums) && int64(nums[i]) <= miss {
			miss += int64(nums[i])
			i++
		} else {
			miss += miss
			result++
		}
	}
	return result
}
