package contest

import "testing"

func Test_hardestWorker(t *testing.T) {
	type args struct {
		n    int
		logs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 10,
				logs: [][]int{
					{0, 3}, {2, 5}, {0, 9}, {1, 15},
				},
			},
			want: 1,
		},
		{
			args: args{
				n: 26,
				logs: [][]int{
					{1, 1}, {3, 7}, {2, 12}, {7, 17},
				},
			},
			want: 3,
		},
		{
			args: args{
				n: 2,
				logs: [][]int{
					{0, 10}, {1, 20},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hardestWorker(tt.args.n, tt.args.logs); got != tt.want {
				t.Errorf("hardestWorker() = %v, want %v", got, tt.want)
			}
		})
	}
}
