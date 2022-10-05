package grind75

import (
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		p := trees.Slice2TreeNode([]*int{
			pointer.Of(5), pointer.Of(6), pointer.Of(2),
			nil, nil, pointer.Of(7), pointer.Of(4),
		})
		q := trees.Slice2TreeNode([]*int{
			pointer.Of(1), pointer.Of(0), pointer.Of(8),
		})
		root := &trees.TreeNode[int]{
			Val:   3,
			Left:  p,
			Right: q,
		}
		assert.Equal(t, root, lowestCommonAncestor(root, p, q))
	})
	t.Run("2", func(t *testing.T) {
		p := trees.Slice2TreeNode([]*int{
			pointer.Of(5), pointer.Of(6), pointer.Of(2),
			nil, nil, pointer.Of(7), pointer.Of(4),
		})
		q := trees.Slice2TreeNode([]*int{
			pointer.Of(1), pointer.Of(0), pointer.Of(8),
		})
		root := &trees.TreeNode[int]{
			Val:   3,
			Left:  p,
			Right: q,
		}
		assert.Equal(t, p, lowestCommonAncestor(root, &trees.TreeNode[int]{
			Val: 5,
		}, &trees.TreeNode[int]{
			Val: 4,
		}))
	})
	t.Run("3", func(t *testing.T) {
		p := trees.Slice2TreeNode([]*int{
			pointer.Of(5), pointer.Of(6), pointer.Of(2),
			nil, nil, pointer.Of(7), pointer.Of(4),
		})
		q := trees.Slice2TreeNode([]*int{
			pointer.Of(1), pointer.Of(0), pointer.Of(8),
		})
		root := &trees.TreeNode[int]{
			Val:   3,
			Left:  p,
			Right: q,
		}
		assert.Equal(t, p, lowestCommonAncestor(root, &trees.TreeNode[int]{
			Val: 4,
		}, &trees.TreeNode[int]{
			Val: 5,
		}))
	})
}
