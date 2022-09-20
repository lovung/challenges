package medium

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums1: []int{1, 2, 3, 2, 1},
				nums2: []int{3, 2, 1, 4, 7},
			},
			want: 3,
		},
		{
			args: args{
				nums1: []int{1, 2, 3, 2, 1},
				nums2: []int{1, 2, 3, 2, 1, 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
