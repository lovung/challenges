package binarytree

import "github.com/lovung/ds/trees"

// Link: https://leetcode.com/problems/binary-tree-pruning/
func pruneTree(root *trees.TreeNode[int]) *trees.TreeNode[int] {
	if needPruneTree(root) {
		return nil
	}
	return root
}

func needPruneTree(node *trees.TreeNode[int]) bool {
	if node == nil {
		return false
	}
	if needPruneTree(node.Left) {
		node.Left = nil
	}
	if needPruneTree(node.Right) {
		node.Right = nil
	}
	if node.Left == nil && node.Right == nil && node.Val == 0 {
		return true
	}
	return false
}
