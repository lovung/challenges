package grind75

import "testing"

func Test_rob2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{2},
			want: 2,
		},
		{
			nums: []int{2, 3},
			want: 3,
		},
		{
			nums: []int{2, 3, 3},
			want: 3,
		},
		{
			nums: []int{2, 1, 5, 3},
			want: 7,
		},
		{
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob2_1(tt.nums); got != tt.want {
				t.Errorf("rob2_1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := rob2_2(tt.nums); got != tt.want {
				t.Errorf("rob2_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
