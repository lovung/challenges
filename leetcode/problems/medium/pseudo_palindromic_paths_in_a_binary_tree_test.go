package medium

import (
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
)

func Test_pseudoPalindromicPaths(t *testing.T) {
	type args struct {
		root *trees.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(2), pointer.Of(1), pointer.Of(1), pointer.Of(1),
					pointer.Of(3),
				}),
			},
			want: 1,
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(2), pointer.Of(3), pointer.Of(1), pointer.Of(3),
					pointer.Of(1), nil, pointer.Of(1),
				}),
			},
			want: 2,
		},
		{
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(2), pointer.Of(3), pointer.Of(1), pointer.Of(3),
					pointer.Of(1), pointer.Of(1),
				}),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoPalindromicPaths(tt.args.root); got != tt.want {
				t.Errorf("pseudoPalindromicPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
