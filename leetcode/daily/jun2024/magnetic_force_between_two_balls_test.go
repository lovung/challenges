package jun2024

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		position []int
		m        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				position: []int{1, 2, 3, 4, 7},
				m:        3,
			},
			want: 3,
		},
		{
			args: args{
				position: []int{1, 2, 3, 4, 9},
				m:        3,
			},
			want: 3,
		},
		{
			args: args{
				position: []int{5, 4, 3, 2, 1, 1000000000},
				m:        2,
			},
			want: 999999999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.position, tt.args.m); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
