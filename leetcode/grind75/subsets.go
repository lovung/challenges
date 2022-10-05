package grind75

// Link: https://leetcode.com/problems/subsets
// Cascading
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	for i := range nums {
		for j := range res {
			newSubset := make([]int, len(res[j])+1)
			copy(newSubset, res[j])
			newSubset[len(newSubset)-1] = nums[i]
			res = append(res, newSubset)
		}
	}
	return res
}
