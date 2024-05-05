package weekly395

import "testing"

func Test_minimumAddedInteger(t *testing.T) {
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
				nums1: []int{9, 4, 3, 9, 4},
				nums2: []int{7, 8, 8},
			},
			want: 4,
		},
		{
			args: args{
				nums1: []int{4, 20, 16, 12, 8},
				nums2: []int{14, 18, 10},
			},
			want: -2,
		},
		{
			args: args{
				nums1: []int{3, 5, 5, 3},
				nums2: []int{7, 7},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAddedInteger(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minimumAddedInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
