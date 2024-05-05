package weekly395

import "testing"

func Test_minEnd(t *testing.T) {
	type args struct {
		n int
		x int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				n: 3,
				x: 4,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEnd(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("minEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
