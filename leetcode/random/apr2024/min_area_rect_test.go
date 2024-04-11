package apr2024

import "testing"

func Test_minAreaRect(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				points: [][]int{{3, 3}, {3, 1}, {1, 3}, {1, 1}, {2, 2}},
			},
			want: 4,
		},
		{
			args: args{
				points: [][]int{
					{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2},
				},
			},
			want: 4,
		},
		{
			args: args{
				points: [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}},
			},
			want: 2,
		},
		{
			args: args{
				points: [][]int{{3, 2}, {3, 1}, {4, 4}, {1, 1}, {4, 3}, {0, 3}, {0, 2}, {4, 0}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAreaRect(tt.args.points); got != tt.want {
				t.Errorf("minAreaRect() = %v, want %v", got, tt.want)
			}
		})
	}
}
