package contest392

import (
	"testing"
)

func Test_distance(t *testing.T) {
	type args struct {
		a byte
		b byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: 'a',
				b: 'z',
			},
			want: 1,
		},
		{
			args: args{
				a: 'z',
				b: 'a',
			},
			want: 1,
		},
		{
			args: args{
				a: 'a',
				b: 'c',
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSmallestString(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "zbbz",
				k: 3,
			},
			want: "aaaz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSmallestString(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("getSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
