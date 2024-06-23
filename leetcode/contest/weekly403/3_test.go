package weekly403

import "testing"

func Test_maximumTotalCost(t *testing.T) {
	tests := []struct {
		nums []int
		want int64
	}{
		{nums: []int{1, -2, 3, 4}, want: 10},
		{nums: []int{1, -1, 1, -1}, want: 4},
		{nums: []int{0}, want: 0},
		{nums: []int{1, -1}, want: 2},
		{nums: []int{1, -2, -3, -5, 4}, want: 9},
		{nums: []int{-14, -13, -20}, want: -7},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumTotalCost(tt.nums); got != tt.want {
				t.Errorf("maximumTotalCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
