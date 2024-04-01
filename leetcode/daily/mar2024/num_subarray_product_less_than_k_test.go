package mar2024

import "testing"

func Test_numSubarrayProductLessThanK(t *testing.T) {
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
				nums: []int{57, 44, 92, 28, 66, 60, 37, 33, 52, 38, 29, 76, 8, 75, 22},
				k:    18,
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{10, 5, 2, 6},
				k:    100,
			},
			want: 8,
		},
		{
			args: args{
				nums: []int{1, 2, 3},
				k:    0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK2() = %v, want %v", got, tt.want)
			}
		})
	}
}
