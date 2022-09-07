package binarytree

import (
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_tree2str(t *testing.T) {
	type args struct {
		root *trees.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				root: nil,
			},
			want: "",
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(1), pointer.Of(2), pointer.Of(3), pointer.Of(4),
				}),
			},
			want: "1(2(4))(3)",
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(1), pointer.Of(2), pointer.Of(3), nil, pointer.Of(4),
				}),
			},
			want: "1(2()(4))(3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.root); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
