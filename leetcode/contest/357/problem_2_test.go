package contest

import "testing"

func Test_canSplitArray(t *testing.T) {
	type args struct {
		nums []int
		m    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{2},
				m:    5,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{1, 2},
				m:    5,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{4, 1, 3, 2, 4},
				m:    6,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{2, 2, 1},
				m:    4,
			},
			want: true,
		},
		{
			args: args{
				nums: []int{2, 1, 3},
				m:    5,
			},
			want: false,
		},
		{
			args: args{
				nums: []int{2, 3, 3, 2, 3},
				m:    6,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canSplitArray(tt.args.nums, tt.args.m); got != tt.want {
				t.Errorf("canSplitArray() = %v, want %v", got, tt.want)
			}
			if got := canSplitArray2(tt.args.nums, tt.args.m); got != tt.want {
				t.Errorf("canSplitArray2() = %v, want %v", got, tt.want)
			}
			if got := canSplitArray3(tt.args.nums, tt.args.m); got != tt.want {
				t.Errorf("canSplitArray3() = %v, want %v", got, tt.want)
			}
		})
	}
}
