package weekly403

import "testing"

func Test_minimumArea(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				grid: [][]int{
					{1, 1, 0},
					{1, 0, 1},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumArea(tt.args.grid); got != tt.want {
				t.Errorf("minimumArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
