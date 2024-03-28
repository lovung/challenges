package mar2024

// https://leetcode.com/problems/subarray-product-less-than-k
func numSubarrayProductLessThanK(nums []int, k int) int {
	cnt, product := 0, 1
	l, r := 0, 0
	for r < len(nums) {
		if product*nums[r] < k {
			product *= nums[r]
			r++
			cnt += r - l
			continue
		}
		if l == r {
			r++
		} else {
			product /= nums[l]
		}
		l++
	}
	return cnt
}

func numSubarrayProductLessThanK2(nums []int, k int) int {
	cnt := 0
	for i := range nums {
		product := 1
		for j := i; j < len(nums); j++ {
			product *= nums[j]
			if product >= k {
				break
			}
			cnt++
		}
	}
	return cnt
}
