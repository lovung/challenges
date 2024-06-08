package jun2024

import "testing"

func Test_checkSubarraySum(t *testing.T) {
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
				nums: []int{23, 2, 4, 6, 7},
				k:    6,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{23, 2, 6, 4, 7},
				k:    6,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{23, 2, 6, 4, 7},
				k:    13,
			},
			want: false,
		},
		{
			args: args{
				nums: []int{5, 0, 0, 0},
				k:    3,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{5, 0, 0, 0},
				k:    3,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{0},
				k:    1,
			},
			want: false,
		},
		{
			args: args{
				nums: []int{1, 0},
				k:    2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
