package oct2022

import (
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *trees.TreeNode[int]
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(5), pointer.Of(4), pointer.Of(8),
					pointer.Of(11), nil, pointer.Of(13), pointer.Of(4),
					pointer.Of(7), pointer.Of(2), nil, nil, nil, pointer.Of(1),
				}),
				targetSum: 22,
			},
			want: true,
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(1), pointer.Of(2), pointer.Of(3),
				}),
				targetSum: 5,
			},
			want: false,
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(1), pointer.Of(2),
				}),
				targetSum: 1,
			},
			want: false,
		},
		{
			args: args{
				root:      nil,
				targetSum: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
