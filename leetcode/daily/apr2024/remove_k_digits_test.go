package apr2024

import (
	"testing"
)

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				num: "1432219",
				k:   3,
			},
			want: "1219",
		},
		{
			args: args{
				num: "10200",
				k:   1,
			},
			want: "200",
		},
		{
			args: args{
				num: "10",
				k:   2,
			},
			want: "0",
		},
		{
			args: args{
				num: "1230",
				k:   3,
			},
			want: "0",
		},
		{
			args: args{
				num: "1230",
				k:   2,
			},
			want: "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits1(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits2(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits2() = %v, want %v", got, tt.want)
			}
		})
	}
}
