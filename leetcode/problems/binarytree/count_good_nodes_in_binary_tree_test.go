package binarytree

import (
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_goodNodes(t *testing.T) {
	type args struct {
		root *trees.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(3), pointer.Of(1), pointer.Of(4), pointer.Of(3), nil, nil, pointer.Of(1), pointer.Of(5),
				}),
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(3), pointer.Of(3), nil, pointer.Of(4), pointer.Of(2),
				}),
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(1),
				}),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodNodes(tt.args.root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
