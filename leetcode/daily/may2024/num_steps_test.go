package may2024

import "testing"

func Test_numSteps(t *testing.T) {
	type args struct {
		org string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				org: "1101",
			},
			want: 6,
		},
		{
			args: args{
				org: "10",
			},
			want: 1,
		},
		{
			args: args{
				org: "1",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSteps(tt.args.org); got != tt.want {
				t.Errorf("numSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
