package grind75

import (
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *trees.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(2), pointer.Of(1), pointer.Of(3),
				}),
			},
			want: true,
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(2), pointer.Of(2), pointer.Of(2),
				}),
			},
			want: false,
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(5), pointer.Of(1), pointer.Of(4),
					nil, nil, pointer.Of(3), pointer.Of(6),
				}),
			},
			want: false,
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(5), pointer.Of(4), pointer.Of(6),
					nil, nil, pointer.Of(3), pointer.Of(7),
				}),
			},
			want: false,
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(32), pointer.Of(26), pointer.Of(47),
					pointer.Of(19), nil, nil, pointer.Of(56), nil, pointer.Of(27),
				}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
