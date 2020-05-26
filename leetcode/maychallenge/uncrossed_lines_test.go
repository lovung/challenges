package maychallenge

import "testing"

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				A: []int{1, 4, 2},
				B: []int{1, 2, 4},
			},
			want: 2,
		},
		{
			name: "test 2",
			args: args{
				A: []int{2, 5, 1, 2, 5},
				B: []int{10, 5, 2, 1, 5, 2},
			},
			want: 3,
		},
		{
			name: "test 3",
			args: args{
				A: []int{1, 3, 7, 1, 7, 5},
				B: []int{1, 9, 2, 5, 1},
			},
			want: 2,
		},
		{
			name: "test 4",
			args: args{
				A: []int{4, 1, 2, 5, 1, 5, 3, 4, 1, 5},
				B: []int{3, 1, 1, 3, 2, 5, 2, 4, 1, 3, 2, 2, 5, 5, 3, 5, 5, 1, 2, 1},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("maxUncrossedLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
