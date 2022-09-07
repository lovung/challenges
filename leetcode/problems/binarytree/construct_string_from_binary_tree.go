package binarytree

import (
	"strconv"

	"github.com/lovung/ds/trees"
)

// Link: https://leetcode.com/problems/construct-string-from-binary-tree/submissions/
func tree2str(root *trees.TreeNode[int]) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val)
	}
	str := strconv.Itoa(root.Val)
	if root.Left == nil {
		str += "()"
	} else {
		str += "(" + tree2str(root.Left) + ")"
	}
	if root.Right == nil {
		return str
	}
	return str + "(" + tree2str(root.Right) + ")"
}
