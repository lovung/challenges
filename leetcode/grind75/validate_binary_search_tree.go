package grind75

import (
	"github.com/lovung/ds/maths"
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

func fowardMax(max *int, val int) *int {
	if max == nil {
		return nil
	}
	newMax := maths.Max(*max, val)
	return &newMax
}

func fowardMin(min *int, val int) *int {
	if min == nil {
		return nil
	}
	newMin := maths.Min(*min, val)
	return &newMin
}
