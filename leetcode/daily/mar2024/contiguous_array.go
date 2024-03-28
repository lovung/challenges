package mar2024

// https://leetcode.com/problems/contiguous-array/description/
func findMaxLength1(nums []int) int {
	prefixSum := make([]int, len(nums)+1)
	sum := 0
	for i := range nums {
		sum += nums[i]
		prefixSum[i+1] = sum
	}
	maxLen := 0
	for r := range prefixSum {
		for l := r % 2; l < r-maxLen; l = l + 2 {
			if r-l == 2*(prefixSum[r]-prefixSum[l]) {
				maxLen = max(maxLen, r-l)
				break
			}
		}
	}
	return maxLen
}

func findMaxLength2(nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	sum := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}
		prefixSum[i+1] = sum
	}
	maxLen := 0
	counted := make(map[int]bool)
	for i := 0; i <= n-maxLen; i++ {
		if counted[prefixSum[i]] {
			continue
		}
		for j := n; j > i; j-- {
			if prefixSum[i] == prefixSum[j] {
				maxLen = max(maxLen, j-i)
				counted[prefixSum[i]] = true
				break
			}
		}
	}
	return maxLen
}

func findMaxLength(nums []int) int {
	firstIndexSum := make(map[int]int)
	firstIndexSum[0] = -1
	sum, res := 0, 0
	for r := range nums {
		if nums[r] == 1 {
			sum++
		} else {
			sum--
		}
		if l, ok := firstIndexSum[sum]; ok {
			res = max(res, r-l)
		} else {
			firstIndexSum[sum] = r
		}
	}
	return res
}

//   1 0 0 0 1 1 0 0 0
// 0 1 1 1 1 2 3 3 3 3

// res = max(j-i, where((j - i) == 2* (p[j]-p[i]))
