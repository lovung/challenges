package biweekly133

import "testing"

func TestMinimumOperations(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{nums: []int{1, 2, 3, 4}, output: 3},
		{nums: []int{3, 6, 9}, output: 0},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := minimumOperations(tt.nums)
			if result != tt.output {
				t.Errorf("minimumOperations(%v) = %d; want %d", tt.nums, result, tt.output)
			}
		})
	}
}
