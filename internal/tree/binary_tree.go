package tree

// TreeNode is a node of a binary tree.
type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func (t *TreeNode[T]) ToListPreorder() []T {
	return Tree2Preorder(t)
}

func (t *TreeNode[T]) ToListInorder() []T {
	return Tree2Inorder(t)
}

func (t *TreeNode[T]) ToListPostorder() []T {
	return Tree2Postorder(t)
}

func NewTreeNode[T any](value T) *TreeNode[T] {
	return &TreeNode[T]{
		Val: value,
	}
}

// Ref: https://github.com/halfrost/LeetCode-Go
func Slice2TreeNode[T any](ints []*T) *TreeNode[T] {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode[T]{
		Val: *ints[0],
	}

	queue := make([]*TreeNode[T], 1, n*2)
	queue[0] = root

	i := 1
	for i < n && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != nil {
			node.Left = NewTreeNode(*ints[i])
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != nil {
			node.Right = NewTreeNode(*ints[i])
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Tree2Preorder travels with pre-order
// - root.data
// - root.left
// - root.right
func Tree2Preorder[T any](root *TreeNode[T]) []T {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []T{root.Val}
	}

	res := []T{root.Val}
	res = append(res, Tree2Preorder(root.Left)...)
	res = append(res, Tree2Preorder(root.Right)...)

	return res
}

// Tree2Inorder travels with in-order
// - root.left
// - root.data
// - root.right
func Tree2Inorder[T any](root *TreeNode[T]) []T {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []T{root.Val}
	}

	res := Tree2Inorder(root.Left)
	res = append(res, root.Val)
	res = append(res, Tree2Inorder(root.Right)...)

	return res
}

// Tree2Postorder traversal with post-order
// - root.left
// - root.right
// - root.data
func Tree2Postorder[T any](root *TreeNode[T]) []T {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []T{root.Val}
	}

	res := Tree2Postorder(root.Left)
	res = append(res, Tree2Postorder(root.Right)...)
	res = append(res, root.Val)

	return res
}
