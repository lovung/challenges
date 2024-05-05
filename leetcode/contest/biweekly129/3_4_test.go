package biweekly129

import (
	"testing"
)

func Test_numberOfStableArrays(t *testing.T) {
	type args struct {
		zero  int
		one   int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				zero:  3,
				one:   5,
				limit: 2,
			},
			want: 16,
		},
		{
			args: args{
				zero:  3,
				one:   4,
				limit: 2,
			},
			want: 18,
		},
		// {
		// 	args: args{
		// 		zero:  71,
		// 		one:   12,
		// 		limit: 26,
		// 	},
		// 	want: 115201918,
		// },
		{
			args: args{
				zero:  1,
				one:   2,
				limit: 2,
			},
			want: 3,
		},
		{
			args: args{
				zero:  3,
				one:   3,
				limit: 2,
			},
			want: 14,
		},
		{
			args: args{
				zero:  13,
				one:   5,
				limit: 3,
			},
			want: 216,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfStableArrays1(tt.args.zero, tt.args.one, tt.args.limit); got != tt.want {
				t.Errorf("numberOfStableArrays() = %v, want %v", got, tt.want)
			}
			if got := numberOfStableArrays2(tt.args.zero, tt.args.one, tt.args.limit); got != tt.want {
				t.Errorf("numberOfStableArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
