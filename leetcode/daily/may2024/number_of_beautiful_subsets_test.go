package may2024

import "testing"

func Test_beautifulSubsets(t *testing.T) {
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
				nums: []int{2, 4, 6},
				k:    2,
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 1, 2, 3},
				k:    1,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautifulSubsets(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("beautifulSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
