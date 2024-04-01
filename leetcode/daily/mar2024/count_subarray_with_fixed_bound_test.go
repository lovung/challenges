package mar2024

import "testing"

func Test_countSubarraysWithFixedBounds(t *testing.T) {
	type args struct {
		nums []int
		minK int
		maxK int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				nums: []int{1, 3, 5, 2, 7, 5},
				minK: 1,
				maxK: 5,
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 1},
				minK: 1,
				maxK: 1,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarraysWithFixedBounds(tt.args.nums, tt.args.minK, tt.args.maxK); got != tt.want {
				t.Errorf("countSubarraysWithFixedBounds() = %v, want %v", got, tt.want)
			}
		})
	}
}
