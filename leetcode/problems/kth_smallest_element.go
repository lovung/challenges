package problems

// import (
// 	"fmt"
// )

// Link: https://leetcode.com/submissions/detail/342120127/?from=/explore/challenge/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3335/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	stack := NewStack(k)
	worker := root
	popedStack := false
	for k > 0 && worker != nil {
		// fmt.Printf("worker value: %v. Pop?: %v\n", worker.Val, popedStack)
		if popedStack {
			k--
			if k == 0 {
				break
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
				break
			}
			worker = (stack.pop()).(*TreeNode)
			popedStack = true
			continue
		}
		k--
		if k == 0 {
			break
		}
		worker = worker.Right
		popedStack = false
	}
	return worker.Val
}
