package oct2022

import "testing"

func Test_minDifficulty(t *testing.T) {
	type args struct {
		jobDifficulty []int
		d             int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				jobDifficulty: []int{6, 5, 4, 3, 2, 1},
				d:             2,
			},
			want: 7,
		},
		{
			args: args{
				jobDifficulty: []int{1, 1, 1},
				d:             3,
			},
			want: 3,
		},
		{
			args: args{
				jobDifficulty: []int{6},
				d:             2,
			},
			want: -1,
		},
		{
			args: args{
				jobDifficulty: []int{6, 5, 4, 7, 2, 9},
				d:             2,
			},
			want: 15,
		},
		{
			args: args{
				jobDifficulty: []int{6, 5, 4, 7, 2, 9},
				d:             3,
			},
			want: 18,
		},
		{
			args: args{
				jobDifficulty: []int{6, 5, 4, 7, 2, 9},
				d:             4,
			},
			want: 24,
		},
		{
			args: args{
				jobDifficulty: []int{380, 302, 102, 681, 863, 676, 243, 671, 651, 612, 162, 561, 394, 856, 601, 30, 6, 257, 921, 405, 716, 126, 158, 476, 889, 699, 668, 930, 139, 164, 641, 801, 480, 756, 797, 915, 275, 709, 161, 358, 461, 938, 914, 557, 121, 964, 315},
				d:             10,
			},
			want: 3807,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifficulty(tt.args.jobDifficulty, tt.args.d); got != tt.want {
				t.Errorf("minDifficulty() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.d < 6 {
				if got := minDifficulty_segmentTree(tt.args.jobDifficulty, tt.args.d); got != tt.want {
					t.Errorf("minDifficulty_segmentTree() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}
