package grind75

import (
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
	"github.com/stretchr/testify/assert"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *trees.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want *trees.TreeNode[int]
	}{
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(4), pointer.Of(2), pointer.Of(7),
					pointer.Of(1), pointer.Of(3), pointer.Of(6), pointer.Of(9),
				}),
			},
			want: trees.Slice2TreeNode([]*int{
				pointer.Of(4), pointer.Of(7), pointer.Of(2),
				pointer.Of(9), pointer.Of(6), pointer.Of(3), pointer.Of(1),
			}),
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(2), pointer.Of(1), pointer.Of(3),
				}),
			},
			want: trees.Slice2TreeNode([]*int{
				pointer.Of(2), pointer.Of(3), pointer.Of(1),
			}),
		},
		{
			args: args{
				root: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, invertTree(tt.args.root))
		})
	}
}
