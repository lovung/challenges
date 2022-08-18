package june2022

import (
	"testing"

	. "github.com/lovung/challenges/internal/tree"
)

func Test_findTarget(t *testing.T) {
	type args struct {
		root *TreeNode[int]
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				root: &TreeNode[int]{
					Val: 5,
					Left: &TreeNode[int]{
						Val: 3,
						Left: &TreeNode[int]{
							Val: 2,
						},
						Right: &TreeNode[int]{
							Val: 4,
						},
					},
					Right: &TreeNode[int]{
						Val: 6,
						Right: &TreeNode[int]{
							Val: 7,
						},
					},
				},
				k: 9,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
