package problems

import (
	"reflect"
	"testing"
)

func Test_newBinaryTree(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	type args struct {
		nums []*int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test 1",
			args: args{
				nums: []*int{},
			},
			want: nil,
		},
		{
			name: "test 2",
			args: args{
				nums: []*int{&nums[1], nil, &nums[3], &nums[4]},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
