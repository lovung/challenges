package medium

import "github.com/lovung/ds/trees"

// Link: https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
func pseudoPalindromicPaths(root *trees.TreeNode[int]) int {
	valCnt := make([]int, 10)
	return recursivePseudoPalindromicPaths(valCnt, root)
}

func recursivePseudoPalindromicPaths(valCnt []int, node *trees.TreeNode[int]) int {
	valCnt[node.Val]++
	switch {
	case node.Left == nil && node.Right == nil:
		return checkIfValid(valCnt)
	case node.Left == nil && node.Right != nil:
		return recursivePseudoPalindromicPaths(valCnt, node.Right)
	case node.Left != nil && node.Right == nil:
		return recursivePseudoPalindromicPaths(valCnt, node.Left)
	default:
		newValCnt := make([]int, len(valCnt))
		copy(newValCnt, valCnt)
		return recursivePseudoPalindromicPaths(valCnt, node.Right) +
			recursivePseudoPalindromicPaths(newValCnt, node.Left)
	}
}

func checkIfValid(valCnt []int) int {
	haveOdd := false
	for i := range valCnt {
		if valCnt[i]%2 != 0 {
			if haveOdd {
				return 0
			}
			haveOdd = true
		}
	}
	return 1
}
