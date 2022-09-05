package binarytree

import (
	"sort"

	"github.com/lovung/ds/trees"
)

// Link: https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type nodeWithLevel struct {
	val    int
	vl, hl int
}

func verticalTraversal(root *trees.TreeNode[int]) [][]int {
	if root == nil {
		return [][]int{}
	}
	levels := make([]*nodeWithLevel, 0)
	recursiveTraversal(root, 0, 0, &levels)
	sort.Slice(levels, func(i, j int) bool {
		if levels[i].hl < levels[j].hl {
			return true
		}
		if levels[i].hl > levels[j].hl {
			return false
		}
		if levels[i].vl < levels[j].vl {
			return true
		}
		if levels[i].vl > levels[j].vl {
			return false
		}
		return levels[i].val < levels[j].val
	})

	curHL := levels[0].hl
	runHL := 0
	res := make([][]int, 0, levels[len(levels)-1].hl-curHL+1)
	res = append(res, make([]int, 0))
	for i := range levels {
		for curHL < levels[i].hl {
			res = append(res, make([]int, 0))
			runHL++
			curHL++
		}
		res[runHL] = append(res[runHL], levels[i].val)
	}
	return res
}

func recursiveTraversal(node *trees.TreeNode[int], vLevel, hLevel int, levels *[]*nodeWithLevel) {
	if node == nil {
		return
	}
	*levels = append(*levels, &nodeWithLevel{
		val: node.Val,
		vl:  vLevel,
		hl:  hLevel,
	})
	recursiveTraversal(node.Left, vLevel+1, hLevel-1, levels)
	recursiveTraversal(node.Right, vLevel+1, hLevel+1, levels)
}
