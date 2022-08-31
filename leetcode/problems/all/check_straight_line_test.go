package problems

import "testing"

func Test_checkStraightLine(t *testing.T) {
	type args struct {
		coordinates [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				coordinates: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}},
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				coordinates: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 7}},
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				coordinates: [][]int{{1, 2}, {2, 2}, {3, 2}, {4, 2}},
			},
			want: true,
		},
		{
			name: "test 4",
			args: args{
				coordinates: [][]int{{1, 2}, {2, 2}, {3, 2}, {4, 3}},
			},
			want: false,
		},
		{
			name: "test 5",
			args: args{
				coordinates: [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStraightLine(tt.args.coordinates); got != tt.want {
				t.Errorf("checkStraightLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
