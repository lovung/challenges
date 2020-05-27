package maychallenge

func findMaxLength(nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	counted := make(map[int]bool)
	sum := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}
		prefixSum[i+1] = sum
	}
	max := 0
	for i := 0; i <= n-max; i++ {
		if counted[prefixSum[i]] {
			continue
		}
		for j := n; j > i; j-- {
			if prefixSum[i] == prefixSum[j] {
				length := j - i
				if max < length {
					max = length
				}
				counted[prefixSum[i]] = true
				break
			}
		}
	}
	return max
}
