package sep2022

import "testing"

func Test_getMinMoves(t *testing.T) {
	type args struct {
		plates []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{
				plates: []int32{2, 4, 3, 1, 6},
			},
			want: 3,
		},
		{
			args: args{
				plates: []int32{4, 11, 9, 10, 12},
			},
			want: 0,
		},
		{
			args: args{
				plates: []int32{3, 2, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinMoves(tt.args.plates); got != tt.want {
				t.Errorf("getMinMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
