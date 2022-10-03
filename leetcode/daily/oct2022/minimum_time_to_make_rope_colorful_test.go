package oct2022

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		colors     string
		neededTime []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				colors:     "aabaa",
				neededTime: []int{1, 2, 3, 4, 1},
			},
			want: 2,
		},
		{
			args: args{
				colors:     "aaabaa",
				neededTime: []int{6, 1, 2, 3, 4, 1},
			},
			want: 4,
		},
		{
			args: args{
				colors:     "abaac",
				neededTime: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.colors, tt.args.neededTime); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
