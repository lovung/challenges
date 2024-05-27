package biweekly130

import "testing"

func Test_satisfiesConditions(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want bool
	}{
		// TODO: Add test cases.
		{
			grid: [][]int{{1, 0, 2}, {1, 0, 2}},
			want: true,
		},
		{
			grid: [][]int{{1, 1, 1}, {0, 0, 0}},
			want: false,
		},
		{
			grid: [][]int{{1}, {2}, {3}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := satisfiesConditions(tt.grid); got != tt.want {
				t.Errorf("satisfiesConditions() = %v, want %v", got, tt.want)
			}
		})
	}
}
