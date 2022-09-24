package medium

import (
	"reflect"
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *trees.TreeNode[int]
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(5), pointer.Of(4), pointer.Of(8),
					pointer.Of(11), nil, pointer.Of(13), pointer.Of(4),
					pointer.Of(7), pointer.Of(2), nil, nil, pointer.Of(5), pointer.Of(1),
				}),
				targetSum: 22,
			},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(1), pointer.Of(2), pointer.Of(3),
				}),
				targetSum: 5,
			},
			want: [][]int{},
		},
		{
			args: args{
				root:      nil,
				targetSum: 5,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
