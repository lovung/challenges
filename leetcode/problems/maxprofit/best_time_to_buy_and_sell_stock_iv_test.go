package maxprofit

import "testing"

func Test_maxProfitIV(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				k:      1,
				prices: []int{1},
			},
			want: 0,
		},
		{
			args: args{
				k:      1,
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9},
			},
			want: 9,
		},
		{
			args: args{
				k:      2,
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9},
			},
			want: 17,
		},
		{
			args: args{
				k:      3,
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9},
			},
			want: 22,
		},
		{
			args: args{
				k:      2,
				prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			},
			want: 6,
		},
		{
			args: args{
				k:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			args: args{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
		{
			args: args{
				k:      1,
				prices: []int{2, 4, 7, 5, 9, 6, 4, 10, 12, 1},
			},
			want: 10,
		},
		{
			args: args{
				k:      2,
				prices: []int{2, 4, 7, 5, 9, 6, 4, 10, 12, 1},
			},
			want: 15,
		},
		{
			args: args{
				k:      3,
				prices: []int{2, 4, 7, 5, 9, 6, 4, 10, 12, 1},
			},
			want: 17,
		},
		{
			args: args{
				k:      1,
				prices: []int{6, 1, 6, 4, 3, 0, 2},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitIV(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfitIV() = %v, want %v", got, tt.want)
			}
		})
	}
}
