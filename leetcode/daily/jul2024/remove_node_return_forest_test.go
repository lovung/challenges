package jul2024

import (
	"reflect"
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_delNodes(t *testing.T) {
	type args struct {
		root     *trees.TreeNode[int]
		toDelete []int
	}
	tests := []struct {
		name string
		args args
		want []*trees.TreeNode[int]
	}{
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(1),
					pointer.Of(2),
					pointer.Of(3),
					pointer.Of(4),
					pointer.Of(5),
					pointer.Of(6),
					pointer.Of(7),
				}),
				toDelete: []int{3, 5},
			},
			want: []*trees.TreeNode[int]{
				trees.Slice2TreeNode([]*int{
					pointer.Of(1),
					pointer.Of(2),
					nil,
					pointer.Of(4)}),
				trees.Slice2TreeNode([]*int{pointer.Of(6)}),
				trees.Slice2TreeNode([]*int{pointer.Of(7)}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delNodes(tt.args.root, tt.args.toDelete); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
