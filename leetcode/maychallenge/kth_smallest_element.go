package maychallenge

// import (
// 	"fmt"
// )

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	stack := NewStack(100)
	worker := root
	popedStack := false
	for k > 0 && worker != nil {
		// fmt.Printf("worker value: %v. Pop?: %v\n", worker.Val, popedStack)
		if popedStack {
			k--
			if k == 0 {
				return worker.Val
			}
			if worker.Right == nil {
				worker = (stack.pop()).(*TreeNode)
				popedStack = true
			} else {
				worker = worker.Right
				popedStack = false
			}
			continue
		}
		if worker.Left != nil {
			stack.push(worker)
			worker = worker.Left
			popedStack = false
			continue
		}
		if worker.Right == nil {
			k--
			if k == 0 {
				return worker.Val
			}
			worker = (stack.pop()).(*TreeNode)
			popedStack = true
			continue
		}
		k--
		if k == 0 {
			return worker.Val
		}
		worker = worker.Right
		popedStack = false
	}
	return worker.Val
}
