package grind75

// Link: https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	totalProduct := 1
	cntZero := 0
	zeroIdx := -1
	for i := range nums {
		if nums[i] == 0 {
			cntZero++
			zeroIdx = i
		} else {
			totalProduct *= nums[i]
		}
	}
	res := make([]int, len(nums))
	if cntZero > 1 {
		return res
	} else if cntZero == 1 {
		res[zeroIdx] = totalProduct
		return res
	}
	for i := range nums {
		res[i] = totalProduct / nums[i]
	}
	return res
}

func productExceptSelf2(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	return res
}
