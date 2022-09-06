package problems

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{3, 1, 3, 2, 4, 3},
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{1, 2, 2, 2, 2},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{2, 2},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{2, 1, 3, 1, 1, 3, 1, 3},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 2, 1, 3, 3, 1, 3, 1},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 1, 2, 1, 3, 1, 1},
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 2, 1, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
