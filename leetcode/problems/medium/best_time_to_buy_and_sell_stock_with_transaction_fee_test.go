package medium

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2-uptrend",
			args: args{
				prices: []int{1, 3, 2, 8, 4, 9},
				fee:    2,
			},
			want: 8,
		},
		{
			name: "1-uptrend",
			args: args{
				prices: []int{1, 3, 2, 8, 7, 9},
				fee:    2,
			},
			want: 6,
		},
		{
			name: "1-uptrend-after-downtrend",
			args: args{
				prices: []int{3, 4, 1, 8},
				fee:    2,
			},
			want: 5,
		},
		{
			name: "downtrend",
			args: args{
				prices: []int{8, 4, 5, 2},
				fee:    2,
			},
			want: 0,
		},
		{
			name: "normal",
			args: args{
				prices: []int{1, 3, 7, 5, 10, 3},
				fee:    3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
