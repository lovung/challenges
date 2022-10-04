package grind75

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 8,
			},
			want: -1,
		},
		{
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 0,
			},
			want: true,
		},
		{
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 1,
			},
			want: true,
		},
		{
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 3,
			},
			want: false,
		},
		{
			args: args{
				nums:   []int{1, 0, 1, 1, 1},
				target: 0,
			},
			want: true,
		},
		{
			args: args{
				nums:   []int{1, 1, 1, 0, 1},
				target: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search2() = %v, want %v", got, tt.want)
			}
		})
	}
}
