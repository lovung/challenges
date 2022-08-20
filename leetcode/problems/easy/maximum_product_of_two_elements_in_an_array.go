package easy

// Link: https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
func maxProduct(nums []int) int {
	var firstMax, secondMax int
	for i := range nums {
		if nums[i] > firstMax {
			firstMax, secondMax = nums[i], firstMax
		} else if nums[i] > secondMax {
			secondMax = nums[i]
		}
	}
	return (firstMax - 1) * (secondMax - 1)
}
