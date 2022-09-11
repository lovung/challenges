package contest

import "testing"

func Test_minGroups(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				intervals: [][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}},
			},
			want: 3,
		},
		{
			args: args{
				intervals: [][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minGroups(tt.args.intervals); got != tt.want {
				t.Errorf("minGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
