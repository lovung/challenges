package weekly399

import "testing"

func Test_maximumSumSubsequence(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:    []int{3, 5, 9, 11, 53, 13},
				queries: [][]int{{1, -2}, {0, -3}, {4, 0}, {0, 10}},
			},
			want: 185,
		},
		{
			args: args{
				nums:    []int{3, 5, 9, 11, 53, 13},
				queries: [][]int{{1, -2}, {0, -3}, {5, 0}, {0, 10}},
			},
			want: 261,
		},
		{
			args: args{
				nums:    []int{3, 5, 9},
				queries: [][]int{{1, -2}, {0, -3}},
			},
			want: 21,
		},
		{
			args: args{
				nums:    []int{6},
				queries: [][]int{{0, -2}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSumSubsequence(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("maximumSumSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
