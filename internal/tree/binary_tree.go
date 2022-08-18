package tree

// TreeNode is a node of a binary tree.
type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewTreeNode[T any](value T) *TreeNode[T] {
	return &TreeNode[T]{
		Val: value,
	}
}
