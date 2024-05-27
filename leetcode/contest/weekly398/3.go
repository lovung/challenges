package weekly398

func sumDigitDifferences(nums []int) int64 {
	res := int64(0)
	for nums[0] > 0 {
		digits := countLastDigit(nums)
		res += countDiff(digits)
	}
	return res
}

func countLastDigit(nums []int) []int {
	digits := [10]int{}
	for i := range nums {
		digits[nums[i]%10]++
		nums[i] /= 10
	}
	return digits[:]
}

func countDiff(digits []int) int64 {
	res := int64(0)
	for i := range digits {
		for j := i + 1; j < len(digits); j++ {
			res += int64(digits[i] * digits[j])
		}
	}
	return res
}
