package may2024

func subsetXORSum(nums []int) int {
	res := 0
	for i := range 1 << (len(nums)) {
		res += xorWithBitmask(nums, i)
	}
	return res
}

func xorWithBitmask(nums []int, bitMask int) int {
	res := 0
	bit := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if bitMask&bit != 0 {
			res ^= nums[i]
		}
		bit <<= 1
	}
	return res
}
