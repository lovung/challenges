package grind75

// Link: https://leetcode.com/problems/permutations/
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	recursivePermute([]int{}, nums, &res)
	return res
}

func recursivePermute(picked []int, remain []int, res *[][]int) {
	if len(remain) == 0 {
		*res = append(*res, picked)
		return
	}
	for i, r := range remain {
		newPicked := make([]int, len(picked)+1)
		copy(newPicked, picked)
		newPicked[len(newPicked)-1] = r
		newRemain := make([]int, len(remain)-1)
		copy(newRemain, remain[:i])
		copy(newRemain[i:], remain[i+1:])
		recursivePermute(newPicked, newRemain, res)
	}
}
