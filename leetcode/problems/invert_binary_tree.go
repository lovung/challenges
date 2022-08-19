package problems

import (
	. "github.com/lovung/challenges/internal/tree"
)

func invertTree(root *TreeNode[int]) *TreeNode[int] {
	swap(root)
	return root
}

func swap(n *TreeNode[int]) {
	if n == nil {
		return
	}
	if n.Left == nil && n.Right == nil {
		return
	}
	tmp := n.Left
	n.Left = n.Right
	n.Right = tmp
	swap(n.Left)
	swap(n.Right)
}
