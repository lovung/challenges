package weekly404

import "testing"

func Test_maxHeightOfTriangle(t *testing.T) {
	type args struct {
		red  int
		blue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{2, 4}, want: 3},
		{args: args{2, 1}, want: 2},
		{args: args{1, 1}, want: 1},
		{args: args{10, 1}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxHeightOfTriangle(tt.args.red, tt.args.blue); got != tt.want {
				t.Errorf("maxHeightOfTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
