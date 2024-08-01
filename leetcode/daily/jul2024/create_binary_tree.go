package jul2024

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/create-binary-tree-from-descriptions/description/
// 14:42
func createBinaryTree(descriptions [][]int) *TreeNode {
	noParentNodes := make(map[int]bool)
	parentNodes := make(map[int]*TreeNode)
	valueToNode := make(map[int]*TreeNode)
	for _, d := range descriptions {
		parentNode := &TreeNode{Val: d[1]}
		if valueToNode[d[0]] != nil {
			parentNode = valueToNode[d[0]]
		}
		childNode := &TreeNode{Val: d[1]}
		if valueToNode[d[1]] != nil {
			childNode = valueToNode[d[1]]
			delete(noParentNodes, d[1])
		}
		if d[2] == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
		if parentNodes[d[0]] == nil {
			noParentNodes[d[0]] = true
		}
		parentNodes[d[1]] = parentNode
		valueToNode[d[0]] = parentNode
		valueToNode[d[1]] = childNode
	}
	var root *TreeNode
	for k := range noParentNodes {
		root = valueToNode[k]
	}
	return root
}
