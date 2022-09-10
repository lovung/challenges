package maxprofit

import "testing"

func Test_maxProfitIII(t *testing.T) {
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
				prices: []int{1},
			},
			want: 0,
		},
		{
			args: args{
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9},
			},
			want: 17,
		},
		{
			args: args{
				prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			},
			want: 6,
		},
		{
			args: args{
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			args: args{
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
		{
			args: args{
				prices: []int{2, 4, 7, 5, 9, 6, 4, 10, 12, 1},
			},
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitIII(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
