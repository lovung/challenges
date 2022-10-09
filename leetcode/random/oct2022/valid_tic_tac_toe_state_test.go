package oct2022

import "testing"

func Test_validTicTacToe(t *testing.T) {
	type args struct {
		board []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				board: []string{"XXX", "OOO", "   "},
			},
			want: false,
		},
		{
			args: args{
				board: []string{"X X", "OOO", " X "},
			},
			want: true,
		},
		{
			args: args{
				board: []string{"XOX", " X ", "   "},
			},
			want: false,
		},
		{
			args: args{
				board: []string{"O  ", "   ", "   "},
			},
			want: false,
		},
		{
			args: args{
				board: []string{"   ", "   ", "   "},
			},
			want: true,
		},
		{
			args: args{
				board: []string{"XOX", "O O", "XOX"},
			},
			want: true,
		},
		{
			args: args{
				board: []string{"XXX", "OO ", "   "},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validTicTacToe(tt.args.board); got != tt.want {
				t.Errorf("validTicTacToe() = %v, want %v", got, tt.want)
			}
		})
	}
}
