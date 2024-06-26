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

	t.Run("4", func(t *testing.T) {
		p := trees.Slice2TreeNode([]*int{
			pointer.Of(2), pointer.Of(0), pointer.Of(4),
			nil, nil, pointer.Of(3), pointer.Of(5),
		})
		q := trees.Slice2TreeNode([]*int{
			pointer.Of(8), pointer.Of(7), pointer.Of(9),
		})
		root := &trees.TreeNode[int]{
			Val:   6,
			Left:  p,
			Right: q,
		}
		assert.Equal(t, 4, lowestCommonAncestor(root, &trees.TreeNode[int]{
			Val: 3,
		}, &trees.TreeNode[int]{
			Val: 5,
		}).Val)
	})
}

func Test_lowestCommonAncestor2(t *testing.T) {
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
		assert.Equal(t, root, lowestCommonAncestor2(root, p, q))
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
		assert.Equal(t, p, lowestCommonAncestor2(root, &trees.TreeNode[int]{
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
		assert.Equal(t, p, lowestCommonAncestor2(root, &trees.TreeNode[int]{
			Val: 4,
		}, &trees.TreeNode[int]{
			Val: 5,
		}))
	})

	t.Run("4", func(t *testing.T) {
		p := trees.Slice2TreeNode([]*int{
			pointer.Of(2), pointer.Of(0), pointer.Of(4),
			nil, nil, pointer.Of(3), pointer.Of(5),
		})
		q := trees.Slice2TreeNode([]*int{
			pointer.Of(8), pointer.Of(7), pointer.Of(9),
		})
		root := &trees.TreeNode[int]{
			Val:   6,
			Left:  p,
			Right: q,
		}
		assert.Equal(t, 4, lowestCommonAncestor2(root, &trees.TreeNode[int]{
			Val: 3,
		}, &trees.TreeNode[int]{
			Val: 5,
		}).Val)
	})
}
