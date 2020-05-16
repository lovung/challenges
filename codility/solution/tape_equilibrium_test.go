package solution

import "testing"

func TestTapeEquilibrium(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				A: []int{3, 1, 2, 4, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TapeEquilibrium(tt.args.A); got != tt.want {
				t.Errorf("TapeEquilibrium() = %v, want %v", got, tt.want)
			}
		})
	}
}
