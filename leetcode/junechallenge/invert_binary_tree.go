package junechallenge

func invertTree(root *TreeNode) *TreeNode {
	swap(root)
	return root
}

func swap(n *TreeNode) {
	if n == nil {
		return
	}
	if n.Left == nil && n.Right == nil {
		return
	}
	tmp := n.Left
	n.Left = n.Right
	n.Right = tmp
	swap(n.Left)
	swap(n.Right)
}
