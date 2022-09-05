package tree

// Link: https://leetcode.com/problems/n-ary-tree-level-order-traversal/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// Node of tree
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	levels := make([][]int, 0)
	traversalTopDown(root, 0, &levels)
	return levels
}

func traversalTopDown(node *Node, parentLevel int, levels *[][]int) {
	if node == nil {
		return
	}
	if len(*levels) < parentLevel+1 {
		*levels = append(*levels, make([]int, 0))
	}
	(*levels)[parentLevel] = append((*levels)[parentLevel], node.Val)
	for i := range node.Children {
		traversalTopDown(node.Children[i], parentLevel+1, levels)
	}
}
