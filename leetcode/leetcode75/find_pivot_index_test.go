package leetcode75

import "testing"

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{[]int{1, 7, 3, 6, 5, 6}},
			want: 3,
		},
		{
			args: args{[]int{1, 2, 3}},
			want: -1,
		},
		{
			args: args{[]int{2, 1, -1}},
			want: 0,
		},
		{
			args: args{[]int{-1, -1, -1, -1, -1, 0}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
