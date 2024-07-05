package jun2024

func inorderTraversal(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left, values)
	*values = append(*values, root.Val)
	inorderTraversal(root.Right, values)
}

// Function to build a balanced BST from a sorted slice of values
func buildBalancedBST(values []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	root := &TreeNode{Val: values[mid]}
	root.Left = buildBalancedBST(values, start, mid-1)
	root.Right = buildBalancedBST(values, mid+1, end)
	return root
}

// Function to balance a BST
// https://leetcode.com/problems/balance-a-binary-search-tree/description
func balanceBST(root *TreeNode) *TreeNode {
	var values []int
	inorderTraversal(root, &values)
	return buildBalancedBST(values, 0, len(values)-1)
}
