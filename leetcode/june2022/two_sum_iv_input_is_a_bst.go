package june2022

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	runner := root
	// implement DFS
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	processedNode := make(map[*TreeNode]struct{})
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

func findNode(root *TreeNode, k int) *TreeNode {
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

// Test cases:
// [5,3,6,2,4,null,7]
// 9
// [5,3,6,2,4,null,7]
// 28
// [1]
// 2
// [2,1,3,-4,null,null,null,null,0]
// 2
// [2,1,3]
// 4

func findTarget2(root *TreeNode, k int) bool {
	markMap := make(map[int]struct{})
	return subFindTarget2(root, k, markMap)
}

func subFindTarget2(root *TreeNode, k int, markMap map[int]struct{}) bool {
	if root == nil {
		return false
	}
	if _, ok := markMap[root.Val-k]; ok {
		return true
	}
	markMap[root.Val] = struct{}{}
	return subFindTarget2(root.Left, k, markMap) || subFindTarget2(root.Right, k, markMap)
}
