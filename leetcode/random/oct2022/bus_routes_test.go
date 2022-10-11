package oct2022

import "testing"

func Test_numBusesToDestination(t *testing.T) {
	type args struct {
		routes [][]int
		source int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				routes: [][]int{
					{1, 9, 12, 20, 23, 24, 35, 38},
					{10, 21, 24, 31, 32, 34, 37, 38, 43},
					{10, 19, 28, 37},
					{8},
					{14, 19},
					{11, 17, 23, 31, 41, 43, 44},
					{21, 26, 29, 33},
					{5, 11, 33, 41},
					{4, 5, 8, 9, 24, 44}},
				source: 37,
				target: 28,
			},
			want: 1,
		},
		{
			args: args{
				routes: [][]int{
					{2},
					{2, 8},
				},
				source: 8,
				target: 2,
			},
			want: 1,
		},
		{
			args: args{
				routes: [][]int{
					{1, 2, 7},
					{3, 6, 7},
				},
				source: 1,
				target: 6,
			},
			want: 2,
		},
		{
			args: args{
				routes: [][]int{
					{1, 2, 7},
					{3, 6, 7},
				},
				source: 1,
				target: 5,
			},
			want: -1,
		},
		{
			args: args{
				routes: [][]int{
					{1, 2, 7},
					{3, 6, 7},
				},
				source: 1,
				target: 1,
			},
			want: 0,
		},
		{
			args: args{
				routes: [][]int{
					{1, 2, 7},
					{4, 5, 6},
					{3, 6, 7},
				},
				source: 1,
				target: 5,
			},
			want: 3,
		},
		{
			args: args{
				routes: [][]int{
					{7, 12},
					{4, 5, 15},
					{6},
					{15, 19},
					{9, 12, 13},
				},
				source: 15,
				target: 12,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numBusesToDestination(tt.args.routes, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("numBusesToDestination() = %v, want %v", got, tt.want)
			}
			if got := numBusesToDestination2(tt.args.routes, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("numBusesToDestination2() = %v, want %v", got, tt.want)
			}
		})
	}
}
