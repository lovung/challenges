package grind75

import "github.com/lovung/ds/trees"

func invertTree(root *trees.TreeNode[int]) *trees.TreeNode[int] {
	swapTree(root)
	return root
}

func swapTree(node *trees.TreeNode[int]) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	swapTree(node.Left)
	swapTree(node.Right)
}
