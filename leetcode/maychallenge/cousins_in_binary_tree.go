package maychallenge

/*
 * Link: https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3322/
 */

func isCousins(root *TreeNode, x int, y int) bool {
	level1, p1 := findNode(root, nil, 0, &x)
	level2, p2 := findNode(root, nil, 0, &y)
	return (level1 == level2) && (p1 != p2)
}

func findNode(n, p *TreeNode, level int, ref *int) (int, *TreeNode) {
	if n == nil {
		return 0, nil
	}
	if n.Val == *ref {
		return level, p
	}
	if val, p := findNode(n.Left, n, level+1, ref); p != nil {
		return val, p
	}
	if val, p := findNode(n.Right, n, level+1, ref); p != nil {
		return val, p
	}
	return 0, nil
}
