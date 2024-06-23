package biweekly133

import "testing"

func TestMinOperations2(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{nums: []int{0, 1, 1, 0, 1}, output: 4},
		{nums: []int{1, 0, 0, 0}, output: 1},
		{nums: []int{0, 1, 1, 1, 0, 0}, output: 3},
		{nums: []int{1, 1, 1, 0, 0}, output: 1},
		{nums: []int{0, 0, 0}, output: 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := minOperations2(tt.nums)
			if result != tt.output {
				t.Errorf("minOperations2(%v) = %d; want %d", tt.nums, result, tt.output)
			}
		})
	}
}
