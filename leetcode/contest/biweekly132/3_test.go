package biweekly132

import "testing"

func Test_maximumLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{48, 49, 49, 48, 48},
				k:    1,
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{2, 15},
				k:    2,
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{13, 9, 13},
				k:    1,
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{1, 2, 1, 1, 3},
				k:    2,
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 2, 1, 1, 3},
				k:    0,
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{1, 2, 2, 2, 3, 4, 5, 1},
				k:    0,
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{1, 2, 2, 2, 3, 4, 5, 1},
				k:    2,
			},
			want: 5,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 4, 5, 1},
				k:    0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = maximumLength_wrong(tt.args.nums, tt.args.k)
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
