package contest392

import "testing"

func Test_longestMonotonicSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{1, 4, 3, 3, 2, 1},
			want: 3,
		},
		{
			nums: []int{1, 4, 3, 3, 2},
			want: 2,
		},
		{
			nums: []int{3, 3, 3, 3},
			want: 1,
		},
		{
			nums: []int{3, 2, 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMonotonicSubarray(tt.nums); got != tt.want {
				t.Errorf("longestMonotonicSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
