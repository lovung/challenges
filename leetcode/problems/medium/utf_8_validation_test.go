package medium

import "testing"

func Test_validUtf8(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				data: []int{197, 130, 1},
			},
			want: true,
		},
		{
			args: args{
				data: []int{235, 140, 4},
			},
			want: false,
		},
		{
			args: args{
				data: []int{237},
			},
			want: false,
		},
		{
			args: args{
				data: []int{0b11110000, 0b10111111, 0b10111111, 0b10111111},
			},
			want: true,
		},
		{
			args: args{
				data: []int{0b11111000, 0b10111111, 0b10111111, 0b10111111},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validUtf8(tt.args.data); got != tt.want {
				t.Errorf("validUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
