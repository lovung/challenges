package binarytree

import "github.com/lovung/ds/trees"

// Link: https://leetcode.com/problems/average-of-levels-in-binary-tree/
func averageOfLevels(root *trees.TreeNode[int]) []float64 {
	sum := make([]int, 0)
	cnt := make([]int, 0)
	traversal(root, 0, &sum, &cnt)
	res := make([]float64, len(sum))
	for i := range sum {
		res[i] = float64(sum[i]) / float64(cnt[i])
	}
	return res
}

func traversal(node *trees.TreeNode[int], level int, sum, cnt *[]int) {
	if node == nil {
		return
	}
	if len(*sum) < level+1 {
		*sum = append(*sum, 0)
	}
	if len(*cnt) < level+1 {
		*cnt = append(*cnt, 0)
	}
	(*sum)[level] += node.Val
	(*cnt)[level]++
	traversal(node.Left, level+1, sum, cnt)
	traversal(node.Right, level+1, sum, cnt)
}
