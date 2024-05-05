package apr2024

import "github.com/lovung/ds/trees"

func addOneRow(root *trees.TreeNode[int], val int, depth int) *trees.TreeNode[int] {
	if depth == 1 {
		return &trees.TreeNode[int]{
			Val:  val,
			Left: root,
		}
	}
	addOneRowRecursive(root, val, depth, 1)
	return root
}

func addOneRowRecursive(root *trees.TreeNode[int], val int, depth, nodeDepth int) {
	if root == nil {
		return
	}
	if nodeDepth == depth-1 {
		root.Left = &trees.TreeNode[int]{
			Val:  val,
			Left: root.Left,
		}
		root.Right = &trees.TreeNode[int]{
			Val:   val,
			Right: root.Right,
		}
		return
	}
	addOneRowRecursive(root.Left, val, depth, nodeDepth+1)
	addOneRowRecursive(root.Right, val, depth, nodeDepth+1)
}
