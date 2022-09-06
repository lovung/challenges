package binarytree

import (
	"reflect"
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_pruneTree(t *testing.T) {
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
				root: trees.Slice2TreeNode(
					[]*int{
						pointer.Of(1), nil, pointer.Of(0), pointer.Of(0), pointer.Of(1),
					},
				),
			},
			want: trees.Slice2TreeNode(
				[]*int{
					pointer.Of(1), nil, pointer.Of(0), nil, pointer.Of(1),
				},
			),
		},
		{
			args: args{
				root: trees.Slice2TreeNode(
					[]*int{
						pointer.Of(1), pointer.Of(0), pointer.Of(1),
						pointer.Of(0), pointer.Of(0), pointer.Of(0), pointer.Of(1),
					},
				),
			},
			want: trees.Slice2TreeNode(
				[]*int{
					pointer.Of(1), nil, pointer.Of(1), nil, pointer.Of(1),
				},
			),
		},
		{
			args: args{
				root: trees.Slice2TreeNode(
					[]*int{
						pointer.Of(0),
					},
				),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pruneTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pruneTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
