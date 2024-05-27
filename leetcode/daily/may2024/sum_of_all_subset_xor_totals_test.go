package may2024

import "testing"

func Test_subsetXORSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{},
			want: 0,
		},
		{
			nums: []int{1},
			want: 1,
		},
		{
			nums: []int{1, 3},
			want: 6,
		},
		{
			nums: []int{5, 1, 6},
			want: 28,
		},
		{
			nums: []int{3, 4, 5, 6, 7, 8},
			want: 480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetXORSum(tt.nums); got != tt.want {
				t.Errorf("subsetXORSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
