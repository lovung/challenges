package weekly397

import "testing"

func Test_findPermutationDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "abc",
				t: "bac",
			},
			want: 2,
		},
		{
			args: args{
				s: "abcde",
				t: "edbac",
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPermutationDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findPermutationDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
