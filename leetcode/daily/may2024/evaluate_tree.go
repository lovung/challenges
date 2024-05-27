package may2024

import "github.com/lovung/ds/trees"

func evaluateTree(root *trees.TreeNode[int]) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}
	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	} else {
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}
}
