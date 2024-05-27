package may2024

import "github.com/lovung/ds/maths"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/distribute-coins-in-binary-tree
func distributeCoins(root *TreeNode) int {
	res := 0
	distributeCoinsRecursive(root, &res)
	return res
}

func distributeCoinsRecursive(root *TreeNode, cnt *int) {
	if root == nil {
		return
	}
	lc, ln := 0, 0
	countCoinsAndNodes(root.Left, &lc, &ln)
	*cnt += maths.ABS(lc - ln)
	rc, rn := 0, 0
	countCoinsAndNodes(root.Right, &rc, &rn)
	*cnt += maths.ABS(rc - rn)
	distributeCoinsRecursive(root.Left, cnt)
	distributeCoinsRecursive(root.Right, cnt)
}

func countCoinsAndNodes(root *TreeNode, coins, nodes *int) {
	if root == nil {
		return
	}
	*nodes++
	*coins += root.Val
	countCoinsAndNodes(root.Left, coins, nodes)
	countCoinsAndNodes(root.Right, coins, nodes)
}
