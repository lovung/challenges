package grind75

import (
	"reflect"
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_levelOrder(t *testing.T) {
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
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(3), pointer.Of(9), pointer.Of(20),
					nil, nil, pointer.Of(15), pointer.Of(7),
				}),
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(3), pointer.Of(9), pointer.Of(20),
					pointer.Of(5), pointer.Of(8), pointer.Of(15), pointer.Of(7),
				}),
			},
			want: [][]int{{3}, {9, 20}, {5, 8, 15, 7}},
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(1),
				}),
			},
			want: [][]int{{1}},
		},
		{
			args: args{
				root: nil,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
