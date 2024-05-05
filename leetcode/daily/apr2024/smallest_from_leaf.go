package apr2024

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return ""
	}
	smallest := ""
	recursiveSmallestFromLeafCheck(root, "", &smallest)
	return smallest
}

func recursiveSmallestFromLeafCheck(
	node *TreeNode,
	curString string,
	smallest *string,
) {
	if node == nil {
		return
	}
	if isLeafNode(node) {
		if len(*smallest) == 0 {
			*smallest = string('a'+byte(node.Val)) + curString
			return
		}
		*smallest = min(*smallest, string('a'+byte(node.Val))+curString)
		return
	}
	recursiveSmallestFromLeafCheck(node.Left, string('a'+byte(node.Val))+curString, smallest)
	recursiveSmallestFromLeafCheck(node.Right, string('a'+byte(node.Val))+curString, smallest)
}

func isLeafNode(node *TreeNode) bool {
	if node == nil {
		return false
	}
	return node.Right == nil && node.Left == nil
}
