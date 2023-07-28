package jul2023

import "testing"

func Test_maxRunTime(t *testing.T) {
	type args struct {
		n         int
		batteries []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{
				n:         2,
				batteries: []int{3, 3, 3},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				n:         2,
				batteries: []int{1, 1, 1, 1},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				n:         4,
				batteries: []int{3, 1, 6, 7, 10, 9},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRunTime(tt.args.n, tt.args.batteries); got != tt.want {
				t.Errorf("maxRunTime() = %v, want %v", got, tt.want)
			}
			if got := maxRunTime2(tt.args.n, tt.args.batteries); got != tt.want {
				t.Errorf("maxRunTime2() = %v, want %v", got, tt.want)
			}
		})
	}
}
