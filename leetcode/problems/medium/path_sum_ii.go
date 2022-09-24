package medium

import "github.com/lovung/ds/trees"

// Link: https://leetcode.com/problems/path-sum-ii/
func pathSum(root *trees.TreeNode[int], targetSum int) [][]int {
	res := make([][]int, 0)
	passedValues := make([]int, 0)
	dfsPathSum(&res, passedValues, root, 0, targetSum)
	return res
}

func dfsPathSum(out *[][]int, passedValues []int,
	node *trees.TreeNode[int], curSum, targetSum int) {
	if node == nil {
		return
	}
	// leaf
	if node.Left == nil && node.Right == nil {
		if node.Val+curSum == targetSum {
			*out = append(*out, append(passedValues, node.Val))
		}
	}
	passedValues = append(passedValues, node.Val)
	passedValues2 := make([]int, len(passedValues))
	copy(passedValues2, passedValues)
	dfsPathSum(out, passedValues, node.Left, curSum+node.Val, targetSum)
	dfsPathSum(out, passedValues2, node.Right, curSum+node.Val, targetSum)
}
