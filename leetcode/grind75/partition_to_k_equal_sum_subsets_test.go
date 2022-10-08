package grind75

import "testing"

func Test_canPartitionKSubsets(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{3, 2, 1, 2, 1},
				k:    3,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{4, 3, 2, 3, 5, 2, 1},
				k:    4,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{2, 2, 2, 2, 3, 4, 5},
				k:    4,
			},
			want: false,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    3,
			},
			want: false,
		},
		{
			args: args{
				nums: []int{6, 5, 9, 6, 3, 5, 1, 10, 4, 1, 4, 3, 9, 9, 3, 3},
				k:    9,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionKSubsets(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("canPartitionKSubsets() = %v, want %v", got, tt.want)
			}
			if got := canPartitionKSubsets2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("canPartitionKSubsets2() = %v, want %v", got, tt.want)
			}
		})
	}
}
