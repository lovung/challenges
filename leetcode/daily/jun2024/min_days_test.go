package jun2024

import "testing"

func Test_minDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				bloomDay: []int{1, 10, 3, 10, 1},
				m:        3,
				k:        1,
			},
			want: 3,
		},
		{
			args: args{
				bloomDay: []int{1, 10, 3, 10, 1},
				m:        3,
				k:        2,
			},
			want: -1,
		},
		{
			args: args{
				bloomDay: []int{7, 7, 7, 7, 12, 7, 7},
				m:        2,
				k:        3,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args.bloomDay, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
