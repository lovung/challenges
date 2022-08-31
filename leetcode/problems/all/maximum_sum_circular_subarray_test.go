package problems

import "testing"

func Test_maxSubarraySumCircular(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 0",
			args: args{
				A: []int{1, -2, 3, -2},
			},
			want: 3,
		},
		{
			name: "test 1",
			args: args{
				A: []int{5, -3, -10, 5},
			},
			want: 10,
		},
		{
			name: "test 2",
			args: args{
				A: []int{5, -3, 5},
			},
			want: 10,
		},
		{
			name: "test 3",
			args: args{
				A: []int{3, -1, 2, -1},
			},
			want: 4,
		},
		{
			name: "test 4",
			args: args{
				A: []int{3, -2, 2, -3},
			},
			want: 3,
		},
		{
			name: "test 5",
			args: args{
				A: []int{-2, -3, -1},
			},
			want: -1,
		},
		{
			name: "test 6",
			args: args{
				A: []int{3, 1, 3, 2, 6},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircular(tt.args.A); got != tt.want {
				t.Errorf("maxSubarraySumCircular() = %v, want %v", got, tt.want)
			}
		})
	}
}
