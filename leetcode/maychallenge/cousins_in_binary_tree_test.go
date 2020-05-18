package maychallenge

import (
	"testing"
)

func Test_isCousins(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	type args struct {
		root *TreeNode
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
			want: true,
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

// func Test_findNode(t *testing.T) {
// 	type args struct {
// 		n     *TreeNode
// 		p     *TreeNode
// 		level int
// 		ref   *int
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  int
// 		want1 *TreeNode
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1 := findNode(tt.args.n, tt.args.p, tt.args.level, tt.args.ref)
// 			if got != tt.want {
// 				t.Errorf("findNode() got = %v, want %v", got, tt.want)
// 			}
// 			if !reflect.DeepEqual(got1, tt.want1) {
// 				t.Errorf("findNode() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }
