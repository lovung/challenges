package oct2022

import "github.com/lovung/ds/trees"

// Link: https://leetcode.com/problems/add-one-row-to-tree/
func addOneRow(root *trees.TreeNode[int], val int, depth int) *trees.TreeNode[int] {
	if depth == 1 {
		return &trees.TreeNode[int]{
			Val:  val,
			Left: root,
		}
	}
	addOneRowDeeper(root, val, depth)
	return root
}

func addOneRowDeeper(root *trees.TreeNode[int], val int, depth int) {
	if root == nil {
		return
	}
	if depth == 2 {
		newLeft := &trees.TreeNode[int]{
			Val:  val,
			Left: root.Left,
		}
		newRight := &trees.TreeNode[int]{
			Val:   val,
			Right: root.Right,
		}
		root.Left = newLeft
		root.Right = newRight
		return
	}
	addOneRowDeeper(root.Left, val, depth-1)
	addOneRowDeeper(root.Right, val, depth-1)
}
