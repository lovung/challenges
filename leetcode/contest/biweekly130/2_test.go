package biweekly130

import "testing"

func Test_maxPointsInsideSquare(t *testing.T) {
	type args struct {
		points [][]int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				points: [][]int{{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3}},
				s:      "abdce",
			},
			want: 5,
		},
		{
			args: args{
				points: [][]int{{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3}},
				s:      "abdca",
			},
			want: 2,
		},
		{
			args: args{
				points: [][]int{{1, 1}, {-2, -2}, {-2, 2}},
				s:      "abb",
			},
			want: 1,
		},
		{
			args: args{
				points: [][]int{{1, 1}, {-1, -1}, {2, -2}},
				s:      "ccd",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPointsInsideSquare(tt.args.points, tt.args.s); got != tt.want {
				t.Errorf("maxPointsInsideSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
