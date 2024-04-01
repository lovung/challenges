package mar2024

import (
	"testing"
)

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     2,
			},
			want: 8,
		},
		{
			args: args{
				tasks: []byte{'A', 'C', 'A', 'B', 'D', 'B'},
				n:     1,
			},
			want: 6,
		},
		{
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     3,
			},
			want: 10,
		},
		{
			args: args{
				tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
				n:     1,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval2(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval2() = %v, want %v", got, tt.want)
			}
		})
	}
}
