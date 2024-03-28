package mar2024

import (
	"fmt"
	"testing"
)

func Test_firstMissingPositive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			nums: []int{3, 4, -1, 1, 1, 0, 3, 4, 1},
			want: 2,
		},
		{
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
		{
			nums: []int{7, 8, 9, 9, 11, 12},
			want: 1,
		},
		{
			nums: []int{1, 2},
			want: 3,
		},
		{
			nums: []int{-1, 4, 2, 1, 9, 10},
			want: 3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("sol_0_%d", i), func(t *testing.T) {
			if got := firstMissingPositive(tt.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("sol_1_%d", i), func(t *testing.T) {
			if got := firstMissingPositive1(tt.nums); got != tt.want {
				t.Errorf("firstMissingPositive1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("sol_2_%d", i), func(t *testing.T) {
			if got := firstMissingPositive2(tt.nums); got != tt.want {
				t.Errorf("firstMissingPositive2() = %v, want %v", got, tt.want)
			}
		})
	}
}
