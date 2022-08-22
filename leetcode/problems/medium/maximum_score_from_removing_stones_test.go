package medium

import "testing"

func Test_maximumScore(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{2, 4, 6},
			want: 6,
		},
		{
			name: "test2",
			args: args{1, 3, 2},
			want: 3,
		},
		{
			name: "test3",
			args: args{4, 4, 6},
			want: 7,
		},
		{
			name: "test4",
			args: args{6, 2, 1},
			want: 3,
		},
		{
			name: "test5",
			args: args{6, 1, 2},
			want: 3,
		},
		{
			name: "test6",
			args: args{2, 3, 1},
			want: 3,
		},
		{
			name: "test7",
			args: args{2, 1, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
