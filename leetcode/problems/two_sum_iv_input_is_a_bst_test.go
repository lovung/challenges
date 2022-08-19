package problems

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
				k: 28,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget2(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findTarget1(b *testing.B) {
	tree := &TreeNode[int]{
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
	}
	for i := 0; i < b.N; i++ {
		findTarget(tree, 9)
		findTarget(tree, 28)
	}
}

func Benchmark_findTarget2(b *testing.B) {
	tree := &TreeNode[int]{
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
	}
	for i := 0; i < b.N; i++ {
		findTarget2(tree, 9)
		findTarget2(tree, 28)
	}
}
