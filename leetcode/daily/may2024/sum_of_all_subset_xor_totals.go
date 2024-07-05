package may2024

import "slices"

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
	for _, n := range slices.Backward(nums) {
		if bitMask&bit != 0 {
			res ^= n
		}
		bit <<= 1
	}
	return res
}
