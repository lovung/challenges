package weekly401

import "testing"

func Test_maxTotalReward(t *testing.T) {
	type args struct {
		rewardValues []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				rewardValues: []int{5, 8, 9, 12},
			},
			want: 21,
		},
		{
			args: args{
				rewardValues: []int{1, 1, 3, 3},
			},
			want: 4,
		},
		{
			args: args{
				rewardValues: []int{1, 6, 4, 3, 2},
			},
			want: 11,
		},
		{
			args: args{
				rewardValues: []int{1, 5, 7, 14, 15},
			},
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTotalReward(tt.args.rewardValues); got != tt.want {
				t.Errorf("maxTotalReward() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTotalReward2(tt.args.rewardValues); got != tt.want {
				t.Errorf("maxTotalReward2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTotalReward1(tt.args.rewardValues); got != tt.want {
				t.Errorf("maxTotalReward1() = %v, want %v", got, tt.want)
			}
		})
	}
}
