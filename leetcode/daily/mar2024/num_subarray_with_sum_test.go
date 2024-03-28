package mar2024

import "testing"

func Test_numSubarraysWithSum(t *testing.T) {
	type args struct {
		nums []int
		goal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, 0, 1, 0, 1},
				goal: 2,
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{0, 0, 1, 1, 0, 0, 0},
				goal: 2,
			},
			want: 12,
		},
		{
			args: args{
				nums: []int{0, 0, 0, 0, 0},
				goal: 0,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarraysWithSum(tt.args.nums, tt.args.goal); got != tt.want {
				t.Errorf("numSubarraysWithSum() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarraysWithSum1(tt.args.nums, tt.args.goal); got != tt.want {
				t.Errorf("numSubarraysWithSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}
