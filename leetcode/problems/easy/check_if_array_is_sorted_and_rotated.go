package easy

// Link: https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/
func check(nums []int) bool {
	dropCount := 0
	for i := 1; i < len(nums) && dropCount <= 2; i++ {
		if nums[i] < nums[i-1] {
			dropCount++
		}
	}
	switch dropCount {
	case 0:
		return true
	case 1:
		return nums[len(nums)-1] <= nums[0]
	default:
		return false
	}
}
