package june2022

import (
	. "github.com/lovung/challenges/internal/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode[int] struct {
 *     Val int
 *     Left *TreeNode[int]
 *     Right *TreeNode[int]
 * }
 */
// Link: https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
func findTarget(root *TreeNode[int], k int) bool {
	runner := root
	// implement DFS
	stack := make([]*TreeNode[int], 0)
	stack = append(stack, root)
	processedNode := make(map[*TreeNode[int]]struct{})
	for len(stack) > 0 {
		runner = stack[0]
		stack = stack[1:]
		if _, ok := processedNode[runner]; !ok {
			processedNode[runner] = struct{}{}
			if matchedNode := findNode(root, k-runner.Val); matchedNode != nil {
				if matchedNode != runner {
					return true
				}
			}
			if runner.Left != nil {
				stack = append(stack, runner.Left)
			}
			if runner.Right != nil {
				stack = append(stack, runner.Right)
			}
		}
	}
	return false
}

func findNode(root *TreeNode[int], k int) *TreeNode[int] {
	runner := root
	for runner != nil {
		if runner.Val == k {
			return runner
		}
		if runner.Val < k {
			runner = runner.Right
		} else {
			runner = runner.Left
		}
	}
	return nil
}

func findTarget2(root *TreeNode[int], k int) bool {
	markMap := make(map[int]struct{})
	return subFindTarget2(root, k, markMap)
}

func subFindTarget2(root *TreeNode[int], k int, markMap map[int]struct{}) bool {
	if root == nil {
		return false
	}
	if _, ok := markMap[k-root.Val]; ok {
		return true
	}
	markMap[root.Val] = struct{}{}
	return subFindTarget2(root.Left, k, markMap) || subFindTarget2(root.Right, k, markMap)
}
