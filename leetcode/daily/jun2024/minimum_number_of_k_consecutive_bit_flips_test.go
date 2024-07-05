package jun2024

import "testing"

func Test_minKBitFlips(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				k:    1,
				nums: []int{0, 1, 0},
			},
			want: 2,
		},
		{
			args: args{
				k:    2,
				nums: []int{1, 1, 0},
			},
			want: -1,
		},
		{
			args: args{
				k:    3,
				nums: []int{0, 0, 0, 1, 0, 1, 1, 0},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minKBitFlips(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minKBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
