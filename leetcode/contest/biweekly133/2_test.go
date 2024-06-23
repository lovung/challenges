package biweekly133

import "testing"

func TestMinOperations1(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{nums: []int{0, 1, 1, 1, 0, 0}, output: 3},
		{nums: []int{0, 1, 1, 1}, output: -1},
		{nums: []int{1, 0, 1, 0, 1, 0}, output: 3},
		{nums: []int{0, 0, 0, 0}, output: -1},
		{nums: []int{1, 1, 1, 1}, output: 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := minOperations1(tt.nums)
			if result != tt.output {
				t.Errorf("minOperations1(%v) = %d; want %d", tt.nums, result, tt.output)
			}
		})
	}
}
