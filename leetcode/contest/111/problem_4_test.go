package contest

import "testing"

func Test_numberOfBeautifulIntegers(t *testing.T) {
	type args struct {
		low  int
		high int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{10, 20, 3},
			want: 2,
		},
		{
			args: args{5, 8, 5},
			want: 0,
		},
		{
			args: args{1, 10, 1},
			want: 1,
		},
		{
			args: args{8, 18, 4},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBeautifulIntegers(tt.args.low, tt.args.high, tt.args.k); got != tt.want {
				t.Errorf("numberOfBeautifulIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
