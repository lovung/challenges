package grind75

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			args: args{
				coins:  []int{2, 5},
				amount: 3,
			},
			want: -1,
		},
		{
			args: args{
				coins:  []int{2, 5, 10, 1},
				amount: 27,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChange2(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			args: args{
				coins:  []int{2, 5},
				amount: 3,
			},
			want: -1,
		},
		{
			args: args{
				coins:  []int{2, 5, 10, 1},
				amount: 27,
			},
			want: 4,
		},
		{
			args: args{
				coins:  []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422},
				amount: 9864,
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange2(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange2() = %v, want %v", got, tt.want)
			}
		})
	}
}
