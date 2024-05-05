package grind75

import (
	"github.com/lovung/ds/trees"
)

// Link: https://leetcode.com/problems/validate-binary-search-tree/
func isValidBST(root *trees.TreeNode[int]) bool {
	return isValidBSTWithCondition(root, nil, nil)
}

func isValidBSTWithCondition(node *trees.TreeNode[int], max, min *int) bool {
	if node == nil {
		return true
	}
	if max != nil && node.Val >= *max {
		return false
	}
	if min != nil && node.Val <= *min {
		return false
	}
	return isValidBSTWithCondition(
		node.Left, &node.Val, fowardMin(min, node.Val),
	) && isValidBSTWithCondition(
		node.Right, fowardMax(max, node.Val), &node.Val,
	)
}

func fowardMax(maxPtr *int, val int) *int {
	if maxPtr == nil {
		return nil
	}
	newMax := max(*maxPtr, val)
	return &newMax
}

func fowardMin(minPtr *int, val int) *int {
	if minPtr == nil {
		return nil
	}
	newMin := min(*minPtr, val)
	return &newMin
}
