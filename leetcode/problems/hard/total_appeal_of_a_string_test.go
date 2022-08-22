package hard

import "testing"

func Test_appealSum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{
				s: "abbca",
			},
			want: 28,
		},
		{
			name: "test2",
			args: args{
				s: "code",
			},
			want: 20,
		},
		{
			name: "test3",
			args: args{
				s: "vulong",
			},
			want: 56,
		},
		{
			name: "test4",
			args: args{
				s: "vulonggolang",
			},
			want: 279,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appealSum(tt.args.s); got != tt.want {
				t.Errorf("appealSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
