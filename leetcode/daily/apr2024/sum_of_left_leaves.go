package apr2024

import "github.com/lovung/ds/trees"

// https://leetcode.com/problems/sum-of-left-leaves

/**
 * Definition for a binary tree node.
 */

func sumOfLeftLeaves(root *trees.TreeNode[int]) int {
	if root == nil {
		return 0
	}
	sum := 0
	if isLeaf(root.Left) {
		sum += root.Left.Val
	}
	sum += sumOfLeftLeaves(root.Left)
	sum += sumOfLeftLeaves(root.Right)
	return sum
}
