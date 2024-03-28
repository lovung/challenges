package microsoft

// Least Common Ancestor
//    1
//   / \
//  2   3
// / \   \
//4   5   6
//
// input: 4 & 5 -> 2
// input: 4 & 2 -> 1
// input: 4 & 6 -> 1

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Have a variable to save the last common ancestor
// Check 2 children of this?
//    -> If both 2 children is not common ancestor -> return the current node
//    -> If 1 of l, r is common ancestor -> recursive to check it

// for checkin if it's ancestor or not: DFS
func LeastCommonAncestor(root, a, b *TreeNode) *TreeNode {
	node := root
	for node != nil {
		checkLeftA := checkIfAncestor(node.Left, a)
		checkLeftB := checkIfAncestor(node.Left, b)
		if checkLeftA && checkLeftB {
			// Common ancestor
			node = node.Left
			continue
		}
		checkRightA := checkIfAncestor(node.Right, a)
		checkRightB := checkIfAncestor(node.Right, b)
		if checkRightA && checkRightB {
			// Common ancestor
			node = node.Right
			continue
		}
		// Least common ancestor
		return node
	}
	return nil
}

func checkIfAncestor(candidate, child *TreeNode) bool {
	if candidate == nil {
		return false
	}
	if candidate.Left != nil && candidate.Left.Val == child.Val {
		return true
	}
	if candidate.Right != nil && candidate.Right.Val == child.Val {
		return true
	}
	return checkIfAncestor(candidate.Left, child) || checkIfAncestor(candidate.Right, child)
}

func LeastCommonAncestor2(root, a, b *TreeNode) *TreeNode {
	return leastCommonAncestor(nil, root, a, b)
}

func leastCommonAncestor(parent, node, a, b *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node == a || node == b {
		return parent
	}
	left := leastCommonAncestor(node, node.Left, a, b)
	right := leastCommonAncestor(node, node.Right, a, b)
	if left == nil {
		return right
	} else if right == nil {
		return left
	}
	return node
}
