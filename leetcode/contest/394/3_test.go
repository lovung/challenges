package contest394

import (
	"testing"
)

func Test_minimumOperations(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{1, 1, 1},
				{0, 0, 0},
			},
			want: 3,
		},
		{
			grid: [][]int{
				{1, 0, 2},
				{1, 0, 2},
			},
			want: 0,
		},
		{
			grid: [][]int{
				{1, 0, 2},
				{1, 1, 3},
				{0, 1, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.grid); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
