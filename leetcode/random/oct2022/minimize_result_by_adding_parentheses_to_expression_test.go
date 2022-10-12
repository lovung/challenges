package oct2022

import "testing"

func Test_minimizeResult(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				expression: "a+38",
			},
			want: "",
		},
		{
			args: args{
				expression: "12+b",
			},
			want: "",
		},
		{
			args: args{
				expression: "247+38",
			},
			want: "2(47+38)",
		},
		{
			args: args{
				expression: "12+34",
			},
			want: "1(2+3)4",
		},
		{
			args: args{
				expression: "999+999",
			},
			want: "(999+999)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizeResult(tt.args.expression); got != tt.want {
				t.Errorf("minimizeResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
