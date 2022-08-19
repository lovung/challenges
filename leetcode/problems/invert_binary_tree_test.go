package problems

import (
	"reflect"
	"testing"

	. "github.com/lovung/challenges/internal/tree"
)

func Test_invertTree(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	type args struct {
		root *TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want *TreeNode[int]
	}{
		{
			name: "test 1",
			args: args{
				root: newBinaryTree([]*int{&nums[1], &nums[2], &nums[3], &nums[4]}),
			},
			want: newBinaryTree([]*int{&nums[1], &nums[3], &nums[2], nil, nil, nil, &nums[4]}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
