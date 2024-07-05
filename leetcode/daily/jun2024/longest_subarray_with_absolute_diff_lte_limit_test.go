package jun2024

import "testing"

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		nums   []int
		limit  int
		output int
	}{
		{nums: []int{8, 2, 4, 7}, limit: 4, output: 2},
		{nums: []int{10, 1, 2, 4, 7, 2}, limit: 5, output: 4},
		{nums: []int{4, 2, 2, 2, 4, 4, 2, 2}, limit: 0, output: 3},
		{nums: []int{1, 5, 6, 7, 8, 10, 6, 5, 6}, limit: 4, output: 5},
		{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, limit: 0, output: 10},
		{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 10, 11, 12}, limit: 9, output: 10},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := longestSubarray1(tt.nums, tt.limit)
			if result != tt.output {
				t.Errorf("longestSubarray1(%v, %d) = %d; want %d", tt.nums, tt.limit, result, tt.output)
			}
		})
		t.Run("", func(t *testing.T) {
			result := longestSubarray2(tt.nums, tt.limit)
			if result != tt.output {
				t.Errorf("longestSubarray2(%v, %d) = %d; want %d", tt.nums, tt.limit, result, tt.output)
			}
		})
	}
}
