package problems

import (
	"testing"

	. "github.com/lovung/ds/trees"
)

func Test_isCousins(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	type args struct {
		root *TreeNode[int]
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				root: newBinaryTree([]*int{&nums[1], &nums[2], &nums[3], &nums[4]}),
				x:    4,
				y:    3,
			},
			want: false,
		},
		{
			name: "test 2",
			args: args{
				root: newBinaryTree([]*int{&nums[1], &nums[2], &nums[3], nil, &nums[4], nil, &nums[5]}),
				x:    5,
				y:    4,
			},
			want: true,
		},
		{
			name: "test 3",
			args: args{
				root: newBinaryTree([]*int{}),
				x:    0,
				y:    0,
			},
			want: false,
		},
		{
			name: "test 4",
			args: args{
				root: newBinaryTree([]*int{&nums[1], nil, &nums[3], nil, &nums[4], nil, &nums[5]}),
				x:    5,
				y:    4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCousins(tt.args.root, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}
