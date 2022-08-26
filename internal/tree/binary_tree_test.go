package tree

import (
	"testing"

	"github.com/lovung/challenges/internal/pointer"
	"github.com/stretchr/testify/assert"
)

func TestSlice2TreeNode(t *testing.T) {
	type args struct {
		ints []*int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode[int]
	}{
		{
			name: "empty",
			args: args{
				ints: []*int{},
			},
			want: nil,
		},
		{
			name: "normal",
			args: args{
				ints: []*int{pointer.Of(1), nil, pointer.Of(3), pointer.Of(5), pointer.Of(7)},
			},
			want: &TreeNode[int]{
				Val: 1,
				Right: &TreeNode[int]{
					Val: 3,
					Left: &TreeNode[int]{
						Val: 5,
					},
					Right: &TreeNode[int]{
						Val: 7,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Slice2TreeNode(tt.args.ints))
		})
	}
}

func TestTraversalOrders(t *testing.T) {
	tree := &TreeNode[int]{
		Val: 1,
		Right: &TreeNode[int]{
			Val: 3,
			Left: &TreeNode[int]{
				Val: 5,
			},
			Right: &TreeNode[int]{
				Val: 7,
			},
		},
	}
	assert.Equal(t, []int{1, 5, 3, 7}, tree.ToListInorder())
	assert.Equal(t, []int{1, 3, 5, 7}, tree.ToListPreorder())
	assert.Equal(t, []int{5, 7, 3, 1}, tree.ToListPostorder())
}
