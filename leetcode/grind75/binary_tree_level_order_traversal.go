package grind75

import "github.com/lovung/ds/trees"

// Link: https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *trees.TreeNode[int]) [][]int {
	res := make([][]int, 0)
	levelOrderTraversal(&res, root, 0)
	return res
}

func levelOrderTraversal(res *[][]int, node *trees.TreeNode[int], parentLevel int) {
	if node == nil {
		return
	}
	if len(*res) < parentLevel+1 {
		*res = append(*res, make([]int, 0, 1))
	}
	(*res)[parentLevel] = append((*res)[parentLevel], node.Val)
	levelOrderTraversal(res, node.Left, parentLevel+1)
	levelOrderTraversal(res, node.Right, parentLevel+1)
}
