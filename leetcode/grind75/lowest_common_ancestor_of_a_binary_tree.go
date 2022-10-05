package grind75

import "github.com/lovung/ds/trees"

// Link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *trees.TreeNode[int]) *trees.TreeNode[int] {
	var pathToP, pathToQ []*trees.TreeNode[int]
	pathFromRoot(root, p.Val, make([]*trees.TreeNode[int], 0), &pathToP)
	pathFromRoot(root, q.Val, make([]*trees.TreeNode[int], 0), &pathToQ)
	return compareAndReturn(pathToP, pathToQ)
}

func compareAndReturn(pathToP, pathToQ []*trees.TreeNode[int]) *trees.TreeNode[int] {
	if len(pathToP) > len(pathToQ) {
		pathToP, pathToQ = pathToQ, pathToP
	}
	var prev *trees.TreeNode[int]
	for i := range pathToP {
		if pathToP[i] != pathToQ[i] {
			return prev
		}
		prev = pathToP[i]
	}
	return prev
}

func pathFromRoot(node *trees.TreeNode[int], val int, passedNodes []*trees.TreeNode[int], finalPath *[]*trees.TreeNode[int]) {
	if node == nil {
		return
	}
	passedNodes = append(passedNodes, node)
	if node.Val == val {
		*finalPath = passedNodes
		return
	}
	clonePassedNodes := make([]*trees.TreeNode[int], len(passedNodes))
	copy(clonePassedNodes, passedNodes)
	pathFromRoot(node.Left, val, passedNodes, finalPath)
	pathFromRoot(node.Right, val, passedNodes, finalPath)
}
