package may2024

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right = removeLeafNodes(root.Right, target)
	root.Left = removeLeafNodes(root.Left, target)
	if isLeaf(root) && target == root.Val {
		return nil
	}
	return root
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
