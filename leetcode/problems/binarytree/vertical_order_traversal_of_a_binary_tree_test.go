package binarytree

import (
	"reflect"
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_verticalTraversal(t *testing.T) {
	type args struct {
		root *trees.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				root: trees.Slice2TreeNode(
					[]*int{
						pointer.Of(3), pointer.Of(9), pointer.Of(20),
						nil, nil, pointer.Of(15), pointer.Of(7),
					},
				),
			},
			want: [][]int{
				{9}, {3, 15}, {20}, {7},
			},
		},
		{
			args: args{
				root: nil,
			},
			want: [][]int{},
		},
		{
			args: args{
				root: trees.Slice2TreeNode(
					[]*int{
						pointer.Of(1), pointer.Of(2), pointer.Of(3),
						pointer.Of(4), pointer.Of(5), pointer.Of(6), pointer.Of(7),
					},
				),
			},
			want: [][]int{
				{4}, {2}, {1, 5, 6}, {3}, {7},
			},
		},
		{
			args: args{
				root: trees.Slice2TreeNode(
					[]*int{
						pointer.Of(1), pointer.Of(2), pointer.Of(3),
						pointer.Of(4), pointer.Of(6), pointer.Of(5), pointer.Of(7),
					},
				),
			},
			want: [][]int{
				{4}, {2}, {1, 5, 6}, {3}, {7},
			},
		},
		{
			args: args{
				root: trees.Slice2TreeNode(
					[]*int{
						pointer.Of(0), pointer.Of(10), pointer.Of(1), nil, nil,
						pointer.Of(2), pointer.Of(4), pointer.Of(3), pointer.Of(5), nil, nil,
						pointer.Of(6), nil, pointer.Of(7), pointer.Of(9), pointer.Of(8), nil, nil, nil, nil,
						pointer.Of(11), nil, nil, pointer.Of(12),
					},
				),
			},
			want: [][]int{
				{8}, {6}, {10, 3}, {0, 2, 7}, {1, 5}, {4, 9, 12}, {11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verticalTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
