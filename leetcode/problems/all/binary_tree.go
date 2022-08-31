package problems

import (
	. "github.com/lovung/ds/trees"
)

func newBinaryTree(nums []*int) *TreeNode[int] {
	length := len(nums)
	if length == 0 {
		return nil
	}
	var (
		stack = make([]*TreeNode[int], length)
		head  = &TreeNode[int]{
			Val:   *nums[0],
			Left:  nil,
			Right: nil,
		}
		parrent *TreeNode[int]
	)
	stack[0] = head
	// fmt.Printf("head: %v, nums: %v\n", head.Val, *nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] == nil {
			continue
		}
		parrent = stack[(i+1)>>1-1]
		if parrent == nil {
			return nil
		}
		newNode := &TreeNode[int]{
			Val:   *nums[i],
			Left:  nil,
			Right: nil,
		}
		stack[i] = newNode

		if i%2 == 1 {
			parrent.Left = newNode
		} else {
			parrent.Right = newNode
		}
		// fmt.Printf("parrent: %v, i: %v, nums: %v\n", parrent.Val, i, *nums[i])
	}
	return head
}
