package microsoft

// Link: https://leetcode.com/problems/sign-of-the-product-of-an-array/description/
func arraySign(nums []int) int {
	x := 1
	for i := range nums {
		x *= signFunc(nums[i])
	}
	return x
}

func signFunc(x int) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}
