package jun2024

// https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips
func minKBitFlips(nums []int, k int) int {
	n, flipped, res := len(nums), 0, 0
	isFlipped := make([]int, n)
	for i := range n {
		if i >= k {
			flipped ^= isFlipped[i-k]
		}
		if flipped == nums[i] {
			if i+k > n {
				return -1
			}
			isFlipped[i] = 1
			flipped ^= 1
			res++
		}
	}
	return res
}
