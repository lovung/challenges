package algorithm

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{-4, -1, 0, 3, 10},
			},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "test2",
			args: args{
				nums: []int{-7, -3, 2, 3, 11},
			},
			want: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_sortedSquares(t *testing.B) {
	for i := 0; i < t.N; i++ {
		sortedSquares([]int{-4, -1, 0, 3, 10})
	}
}


func Benchmark_sortedSquares2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		sortedSquares2([]int{-4, -1, 0, 3, 10})
	}
}
