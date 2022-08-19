package problems

/*
 * Link: https://leetcode.com/problems/cousins-in-binary-tree/
 */

func isCousins(root *TreeNode, x int, y int) bool {
	level1, p1 := isCousinsFindNode(root, nil, 0, &x)
	level2, p2 := isCousinsFindNode(root, nil, 0, &y)
	return (level1 == level2) && (p1 != p2)
}

func isCousinsFindNode(n, p *TreeNode, level int, ref *int) (int, *TreeNode) {
	if n == nil {
		return 0, nil
	}
	if n.Val == *ref {
		return level, p
	}
	if val, p := isCousinsFindNode(n.Left, n, level+1, ref); p != nil {
		return val, p
	}
	if val, p := isCousinsFindNode(n.Right, n, level+1, ref); p != nil {
		return val, p
	}
	return 0, nil
}
