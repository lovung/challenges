package contest392

func longestMonotonicSubarray(nums []int) int {
	return max(
		longestStrickyIncrease(nums),
		longestStrickyDecrease(nums),
	)
}

func longestStrickyIncrease(nums []int) int {
	res, length := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			length = 1
		} else {
			length++
			res = max(res, length)
		}
	}
	return res
}

func longestStrickyDecrease(nums []int) int {
	res, length := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			length = 1
		} else {
			length++
			res = max(res, length)
		}
	}
	return res
}
