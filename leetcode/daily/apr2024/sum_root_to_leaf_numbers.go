package apr2024

import "github.com/lovung/ds/trees"

// https://leetcode.com/problems/sum-root-to-leaf-numbers/
func sumNumbers(root *trees.TreeNode[int]) int {
	sum := 0
	dfsSumNumbers(root, &sum, 0)
	return sum
}

func dfsSumNumbers(root *trees.TreeNode[int], sum *int, num int) {
	if root == nil {
		return
	}
	if isLeaf(root) {
		*sum += num*10 + root.Val
	}
	dfsSumNumbers(root.Right, sum, num*10+root.Val)
	dfsSumNumbers(root.Left, sum, num*10+root.Val)
}

func isLeaf(node *trees.TreeNode[int]) bool {
	if node == nil {
		return false
	}
	return node.Right == nil && node.Left == nil
}
