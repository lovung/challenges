package grind75

import (
	"github.com/lovung/ds/maths"
	"github.com/lovung/ds/trees"
)

// https://leetcode.com/problems/balanced-binary-tree/description/
func isBalanced(root *trees.TreeNode[int]) bool {
	b, _ := isBalancedAndTreeHeight(root)
	return b
}

func isBalancedAndTreeHeight(root *trees.TreeNode[int]) (bool, int) {
	if root == nil {
		return true, 0
	}
	if root.Left == nil && root.Right == nil {
		return true, 1
	}
	b1, h1 := isBalancedAndTreeHeight(root.Left)
	b2, h2 := isBalancedAndTreeHeight(root.Right)
	return b1 && b2 && (h1 == h2 || h1 == h2-1 || h1 == h2+1), 1 + maths.Max(h1, h2)
}
