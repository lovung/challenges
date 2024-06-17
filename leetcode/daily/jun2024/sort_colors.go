package jun2024

// https://leetcode.com/problems/sort-colors/
func sortColors(nums []int) {
	cnt := [3]int{}
	for _, n := range nums {
		cnt[n]++
	}
	i := 0
	for k, n := range cnt {
		for range n {
			nums[i] = k
			i++
		}
	}
}
