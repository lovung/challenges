package contest392

import "testing"

func Test_minOperationsToMakeMedianK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				nums: []int{2, 5, 6, 7, 5},
				k:    4,
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{2, 5, 6, 7, 5},
				k:    7,
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperationsToMakeMedianK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperationsToMakeMedianK() = %v, want %v", got, tt.want)
			}
		})
	}
}
