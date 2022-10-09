package problems

import (
	"testing"

	"github.com/lovung/ds/pointer"
	"github.com/lovung/ds/trees"
	. "github.com/lovung/ds/trees"
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
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(5), pointer.Of(3), pointer.Of(6),
					pointer.Of(2), pointer.Of(4), nil, pointer.Of(7),
				}),
				k: 9,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(5), pointer.Of(3), pointer.Of(6),
					pointer.Of(2), pointer.Of(4), nil, pointer.Of(7),
				}),
				k: 28,
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				root: trees.Slice2TreeNode([]*int{
					pointer.Of(5), pointer.Of(3), pointer.Of(6),
					pointer.Of(2), pointer.Of(4), nil, pointer.Of(7),
				}),
				k: 8,
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
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget2(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget3(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/problems/all
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_findTarget1-12    	 2914141	       429.0 ns/op	      96 B/op	       7 allocs/op
func Benchmark_findTarget1(b *testing.B) {
	tree := trees.Slice2TreeNode([]*int{
		pointer.Of(5), pointer.Of(3), pointer.Of(6),
		pointer.Of(2), pointer.Of(4), nil, pointer.Of(7),
	})
	for i := 0; i < b.N; i++ {
		findTarget(tree, 9)
		findTarget(tree, 28)
		findTarget(tree, 8)
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/problems/all
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_findTarget2-12    	 5638112	       212.7 ns/op	       0 B/op	       0 allocs/op
func Benchmark_findTarget2(b *testing.B) {
	tree := trees.Slice2TreeNode([]*int{
		pointer.Of(5), pointer.Of(3), pointer.Of(6),
		pointer.Of(2), pointer.Of(4), nil, pointer.Of(7),
	})
	for i := 0; i < b.N; i++ {
		findTarget2(tree, 9)
		findTarget2(tree, 28)
		findTarget2(tree, 8)
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/lovung/challenges/leetcode/problems/all
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_findTarget3-12    	 4200997	       285.6 ns/op	     240 B/op	       8 allocs/op
func Benchmark_findTarget3(b *testing.B) {
	tree := trees.Slice2TreeNode([]*int{
		pointer.Of(5), pointer.Of(3), pointer.Of(6),
		pointer.Of(2), pointer.Of(4), nil, pointer.Of(7),
	})
	for i := 0; i < b.N; i++ {
		findTarget3(tree, 9)
		findTarget3(tree, 28)
		findTarget2(tree, 8)
	}
}
