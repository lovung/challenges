package jul2024

import "github.com/lovung/ds/trees"

// https://leetcode.com/problems/delete-nodes-and-return-forest/
func delNodes(root *trees.TreeNode[int], toDelete []int) []*trees.TreeNode[int] {
	deleteMap := make(map[int]bool)
	for _, d := range toDelete {
		deleteMap[d] = true
	}
	var forest []*trees.TreeNode[int]
	if !deleteMap[root.Val] {
		forest = append(forest, root)
	}
	recursiveDelNodes(root, deleteMap, &forest)
	return forest
}

func recursiveDelNodes(root *trees.TreeNode[int], deleteMap map[int]bool, forest *[]*trees.TreeNode[int]) {
	if root == nil {
		return
	}
	recursiveDelNodes(root.Left, deleteMap, forest)
	recursiveDelNodes(root.Right, deleteMap, forest)
	if deleteMap[root.Val] {
		if root.Left != nil && !deleteMap[root.Left.Val] {
			*forest = append(*forest, root.Left)
		}
		if root.Right != nil && !deleteMap[root.Right.Val] {
			*forest = append(*forest, root.Right)
		}
	} else {
		if root.Left != nil && deleteMap[root.Left.Val] {
			root.Left = nil
		}
		if root.Right != nil && deleteMap[root.Right.Val] {
			root.Right = nil
		}
	}
}
