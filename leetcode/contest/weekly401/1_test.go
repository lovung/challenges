package weekly401

import "testing"

func Test_numberOfChild(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 3,
				k: 5,
			},
			want: 1,
		},
		{
			args: args{
				n: 5,
				k: 6,
			},
			want: 2,
		},
		{
			args: args{
				n: 4,
				k: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfChild(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("numberOfChild() = %v, want %v", got, tt.want)
			}
		})
	}
}
