package jun2024

// https://leetcode.com/problems/count-number-of-nice-subarrays
func numberOfSubarrays(nums []int, k int) int {
	l, r := 0, 0
	cntOdd, res := 0, 0
	lastROddIdx := 0
	for ; r < len(nums); r++ {
		if nums[r]%2 != 0 {
			cntOdd++
		}
		if cntOdd > k {
			for ; l < r && cntOdd > k; l++ {
				if nums[l]%2 != 0 {
					cntOdd--
				}
				res += r - lastROddIdx
			}
		}
		if nums[r]%2 != 0 {
			lastROddIdx = r
		}
	}
	for ; l < r && cntOdd == k; l++ {
		if nums[l]%2 != 0 {
			cntOdd--
		}
		res += r - lastROddIdx
	}
	return res
}
