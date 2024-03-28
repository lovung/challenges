package easy

import (
	"github.com/lovung/ds/queue"
	"github.com/lovung/ds/trees"
)

// Link: https://leetcode.com/problems/same-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *trees.TreeNode[int], q *trees.TreeNode[int]) bool {
	queueP := queue.NewSimpleQueue[*trees.TreeNode[int]]()
	queueQ := queue.NewSimpleQueue[*trees.TreeNode[int]]()
	queueP.EnQueue(p)
	queueQ.EnQueue(q)
	for queueP.Len() > 0 && queueQ.Len() > 0 {
		rp, _ := queueP.DeQueue()
		rq, _ := queueQ.DeQueue()
		if rq == nil && rp == nil {
			continue
		}
		if (rq == nil) != (rp == nil) {
			return false
		}
		if rq.Val != rp.Val {
			return false
		}
		if (rq.Left != nil) != (rp.Left != nil) {
			return false
		}
		if (rq.Right != nil) != (rp.Right != nil) {
			return false
		}
		if rq.Left != nil {
			queueP.EnQueue(rp.Left)
			queueQ.EnQueue(rq.Left)
		}
		if rq.Right != nil {
			queueP.EnQueue(rp.Right)
			queueQ.EnQueue(rq.Right)
		}
	}
	return true
}

func isSameTree2(p *trees.TreeNode[int], q *trees.TreeNode[int]) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
