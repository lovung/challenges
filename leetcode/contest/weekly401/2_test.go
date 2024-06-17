package weekly401

import "testing"

func Test_valueAfterKSeconds(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 4,
				k: 5,
			},
			want: 56,
		},
		{
			args: args{
				n: 5,
				k: 3,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valueAfterKSeconds(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("valueAfterKSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
