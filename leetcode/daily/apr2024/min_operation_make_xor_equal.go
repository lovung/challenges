package apr2024

func minOperations(nums []int, k int) int {
	xor := 0
	for i := range nums {
		xor ^= nums[i]
	}
	mask, cnt := 1, 0
	for xor > 0 || k > 0 {
		if k&mask != xor&mask {
			cnt++
		}
		k >>= 1
		xor >>= 1
	}
	return cnt
}
