package oct2022

import (
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
	"github.com/stretchr/testify/assert"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
		root  *trees.TreeNode[int]
		val   int
		depth int
	}
	tests := []struct {
		name string
		args args
		want *trees.TreeNode[int]
	}{
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(4), pointer.Of(2), pointer.Of(6),
					pointer.Of(3), pointer.Of(1), pointer.Of(5),
				}),
				val:   1,
				depth: 1,
			},
			want: trees.Slice2TreeNode([]*int{
				pointer.Of(1), pointer.Of(4), nil, pointer.Of(2), pointer.Of(6),
				pointer.Of(3), pointer.Of(1), pointer.Of(5),
			}),
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(4), pointer.Of(2), pointer.Of(6),
					pointer.Of(3), pointer.Of(1), pointer.Of(5),
				}),
				val:   1,
				depth: 2,
			},
			want: trees.Slice2TreeNode([]*int{
				pointer.Of(4), pointer.Of(1), pointer.Of(1), pointer.Of(2), nil, nil,
				pointer.Of(6), pointer.Of(3), pointer.Of(1), pointer.Of(5),
			}),
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(4), pointer.Of(2), nil,
					pointer.Of(3), pointer.Of(1),
				}),
				val:   1,
				depth: 3,
			},
			want: trees.Slice2TreeNode([]*int{
				pointer.Of(4), pointer.Of(2), nil, pointer.Of(1), pointer.Of(1), pointer.Of(3),
				nil, nil, pointer.Of(1),
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addOneRow(tt.args.root, tt.args.val, tt.args.depth)
			assert.Equal(t, tt.want, got)
		})
	}
}
