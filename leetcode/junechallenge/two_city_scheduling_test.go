package junechallenge

import (
	"testing"
)

func Test_twoCitySchedCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				costs: [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}},
			},
			want: 110,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoCitySchedCost(tt.args.costs); got != tt.want {
				t.Errorf("twoCitySchedCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoCitySchedCostV2(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				costs: [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}},
			},
			want: 110,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoCitySchedCostV2(tt.args.costs); got != tt.want {
				t.Errorf("twoCitySchedCostV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
