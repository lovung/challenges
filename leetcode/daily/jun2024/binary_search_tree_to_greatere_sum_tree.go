package jun2024

import "github.com/emirpasic/gods/v2/stacks/linkedliststack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	travelTree(root, 0)
	return root
}

// Return most-left value
func travelTree(root *TreeNode, greater int) int {
	if root == nil {
		return 0
	}
	sum := greater
	if root.Right != nil {
		sum = travelTree(root.Right, greater)
	}
	root.Val += sum
	sum = root.Val
	if root.Left != nil {
		sum = travelTree(root.Left, sum)
	}
	return sum
}

func bstToGst_norecursive(root *TreeNode) *TreeNode {
	s := linkedliststack.New[*TreeNode]()
	sum := 0
	node := root
	for node != nil {
		s.Push(node)
		node = node.Right
	}
	for !s.Empty() {
		node, _ = s.Pop()
		node.Val += sum
		sum = node.Val
		if node.Left != nil {
			for node = node.Left; node != nil; node = node.Right {
				s.Push(node)
			}
		}
	}
	return root
}
