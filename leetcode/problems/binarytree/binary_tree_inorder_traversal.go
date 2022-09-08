package binarytree

import "github.com/lovung/ds/trees"

func inorderTraversal(root *trees.TreeNode[int]) []int {
	if root == nil {
		return []int{}
	}
	res := append(inorderTraversal(root.Left), root.Val)
	return append(res, inorderTraversal(root.Right)...)
}
