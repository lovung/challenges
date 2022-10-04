package grind75

// Link: https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	recursiveCombinationSum(candidates, target, 0, 0, []int{}, &res)
	return res
}

func recursiveCombinationSum(candidates []int, target, curSum, lastCanIdx int, picked []int, res *[][]int) {
	if curSum > target {
		return
	}
	if curSum == target {
		*res = append(*res, picked)
		return
	}

	for i := lastCanIdx; i < len(candidates); i++ {
		newPicked := make([]int, len(picked)+1)
		copy(newPicked, picked)
		newPicked[len(newPicked)-1] = candidates[i]
		recursiveCombinationSum(candidates, target, curSum+candidates[i], i, newPicked, res)
	}
}
