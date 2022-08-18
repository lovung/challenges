package algorithm

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: []int{1, 2},
		},
		{
			name: "test2",
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
		{
			name: "test3",
			args: args{
				numbers: []int{-1, 0},
				target:  -1,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum2(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{2, 7, 11, 15}, 9)
	}
}

func Benchmark_twoSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum2([]int{2, 7, 11, 15}, 9)
	}
}
