package grind75

import (
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
	"github.com/stretchr/testify/assert"
)

func Test_isBalanced(t *testing.T) {
	assert.True(t, isBalanced(trees.Slice2TreeNode([]*int{
		pointer.Of(5), pointer.Of(6), pointer.Of(2),
		nil, nil, pointer.Of(7), pointer.Of(4),
	})))
	assert.False(t, isBalanced(trees.Slice2TreeNode([]*int{
		pointer.Of(1), pointer.Of(2), pointer.Of(2),
		pointer.Of(3), pointer.Of(3), nil, nil,
		pointer.Of(4), pointer.Of(4),
	})))
	assert.True(t, isBalanced(trees.Slice2TreeNode([]*int{
		pointer.Of(1), pointer.Of(0), pointer.Of(8),
	})))
	assert.True(t, isBalanced(nil))
	assert.True(t, isBalanced(trees.Slice2TreeNode([]*int{
		pointer.Of(1), pointer.Of(0),
	})))
}
