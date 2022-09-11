package contest

// https://leetcode.com/problems/most-frequent-even-element/
func mostFrequentEven(nums []int) int {
	cnt := make([]int, 1e5>>1+1)
	maxCnt := 0
	for i := range nums {
		if nums[i]%2 == 0 {
			cnt[nums[i]>>1]++
			if maxCnt < cnt[nums[i]>>1] {
				maxCnt = cnt[nums[i]>>1]
			}
		}
	}
	if maxCnt > 0 {
		for i := range cnt {
			if cnt[i] == maxCnt {
				return i << 1
			}
		}
	}
	return -1
}
