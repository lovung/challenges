package mar2024

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				nums: []int{1, 3, 2, 3, 3, 1},
				k:    2,
			},
			want: 10,
		},
		{
			args: args{
				nums: []int{1, 3, 2, 3, 3},
				k:    2,
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{1, 4, 2, 1},
				k:    2,
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{
					61, 23, 38, 23, 56, 40, 82, 56, 82, 82,
					82, 70, 8, 69, 8, 7, 19, 14, 58, 42, 82,
					10, 82, 78, 15, 82,
				},
				k: 2,
			},
			want: 224,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countSubarrays2() = %v, want %v", got, tt.want)
			}
		})
	}
}
