package biweekly133

// https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/description/
func minOperations1(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 0 {
			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			res++
		}
	}
	if nums[len(nums)-1] == 1 && nums[len(nums)-2] == 1 {
		return res
	}
	return -1
}
