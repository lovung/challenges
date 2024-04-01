package mar2024

func numSubarraysWithSum1(nums []int, goal int) int {
	prefixSum := make([]int, len(nums)+1)
	sum := 0
	for i := range nums {
		sum += nums[i]
		prefixSum[i+1] = sum
	}
	cnt := 0
	for i := range nums {
		for j := i; j < len(nums); j++ {
			if prefixSum[j+1]-prefixSum[i] == goal {
				cnt++
			}
		}
	}
	return cnt
}

func numSubarraysWithSum(nums []int, goal int) int {
	var count, sum int
	m := make(map[int]int)
	m[0] = 1
	for i := range nums {
		sum += nums[i]
		count += m[sum-goal]
		m[sum]++
	}
	return count
}
