package binarytree

import (
	"reflect"
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *trees.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: trees.Slice2TreeNode(
					[]*int{
						pointer.Of(1), nil, pointer.Of(2), pointer.Of(3),
					},
				),
			},
			want: []int{1, 3, 2},
		},
		{
			args: args{
				root: trees.Slice2TreeNode(
					[]*int{
						pointer.Of(1),
					},
				),
			},
			want: []int{1},
		},
		{
			args: args{
				root: trees.Slice2TreeNode(
					[]*int{},
				),
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
