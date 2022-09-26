package medium

import "testing"

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				equations: []string{"a==b", "b!=a"},
			},
			want: false,
		},
		{
			args: args{
				equations: []string{"a==b", "b==a"},
			},
			want: true,
		},
		{
			args: args{
				equations: []string{"a==b", "b==c", "d==e", "e==f", "f!=a"},
			},
			want: true,
		},
		{
			args: args{
				equations: []string{"a==b", "b==c", "d==e", "e==f", "f!=a", "e==b"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
