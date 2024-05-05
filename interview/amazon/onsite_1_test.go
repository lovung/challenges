package amazon

import "testing"

func TestCountGroupOfHouseType(t *testing.T) {
	type args struct {
		grid   [][]string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				grid: [][]string{
					{"o", "x", "x", "o", "x"},
					{"x", "x", "o", "x", "o"},
					{"o", "x", "x", "o", "x"},
					{"o", "x", "x", "o", "x"},
				},
				target: "x",
			},
			want: 4,
		},
		{
			args: args{
				grid: [][]string{
					{"o", "o", "o", "o", "x"},
					{"x", "x", "o", "x", "x"},
					{"o", "o", "x", "o", "x"},
					{"o", "x", "x", "o", "x"},
				},
				target: "x",
			},
			want: 3,
		},
		{
			args: args{
				grid: [][]string{
					{"o", "o", "o", "o", "x"},
					{"x", "x", "o", "x", "x"},
					{"x", "x", "x", "o", "x"},
					{"o", "x", "x", "o", "x"},
				},
				target: "x",
			},
			want: 2,
		},
		{
			args: args{
				grid: [][]string{
					{"x", "x", "x", "x", "x"},
					{"x", "x", "x", "x", "x"},
					{"x", "x", "x", "x", "x"},
					{"x", "x", "x", "x", "x"},
				},
				target: "x",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountGroupOfHouseType(tt.args.grid, tt.args.target); got != tt.want {
				t.Errorf("CountGroupOfHouseType() = %v, want %v", got, tt.want)
			}
		})
	}
}
