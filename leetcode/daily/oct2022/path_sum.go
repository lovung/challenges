package oct2022

import "github.com/lovung/ds/trees"

// Link: https://leetcode.com/problems/path-sum/
func hasPathSum(root *trees.TreeNode[int], targetSum int) bool {
	if root == nil {
		return false
	}
	return recursiveHasPathSum(root, targetSum)
}

func recursiveHasPathSum(root *trees.TreeNode[int], targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	return recursiveHasPathSum(
		root.Left, targetSum-root.Val,
	) || recursiveHasPathSum(
		root.Right, targetSum-root.Val,
	)
}
