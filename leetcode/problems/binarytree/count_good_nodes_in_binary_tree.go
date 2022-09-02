package binarytree

import "github.com/lovung/ds/trees"

// Link: https://leetcode.com/problems/count-good-nodes-in-binary-tree/
func goodNodes(root *trees.TreeNode[int]) int {
	return countGoodNodes(root, -10e4)
}

func countGoodNodes(current *trees.TreeNode[int], maxUtilParrent int) int {
	if current == nil {
		return 0
	}
	count := 0
	if maxUtilParrent <= current.Val {
		count++
		maxUtilParrent = current.Val
	}
	return count + countGoodNodes(current.Left, maxUtilParrent) + countGoodNodes(current.Right, maxUtilParrent)
}
