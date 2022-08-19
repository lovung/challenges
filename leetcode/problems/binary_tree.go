package problems

// import (
// 	"fmt"
// )

// TreeNode in binary tree
/* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newBinaryTree(nums []*int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	var (
		stack = make([]*TreeNode, length)
		head  = &TreeNode{
			Val:   *nums[0],
			Left:  nil,
			Right: nil,
		}
		parrent *TreeNode
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
		newNode := &TreeNode{
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
