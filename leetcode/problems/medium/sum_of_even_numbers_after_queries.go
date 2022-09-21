package medium

// Link: https://leetcode.com/problems/sum-of-even-numbers-after-queries/
func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	sumOfEvens := 0
	for _, num := range nums {
		if num%2 == 0 {
			sumOfEvens += num
		}
	}
	res := make([]int, 0, len(queries))
	for _, q := range queries {
		if nums[q[1]]%2 == 0 {
			sumOfEvens -= nums[q[1]]
		}
		nums[q[1]] += q[0]
		if nums[q[1]]%2 == 0 {
			sumOfEvens += nums[q[1]]
		}
		res = append(res, sumOfEvens)
	}
	return res
}
