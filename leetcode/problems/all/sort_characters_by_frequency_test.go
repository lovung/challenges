package problems

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				s: "tree",
			},
			want: "eetr",
		},
		{
			name: "test 2",
			args: args{
				s: "cccaaa",
			},
			want: "cccaaa",
		},
		{
			name: "test 3",
			args: args{
				s: "Aabb",
			},
			want: "bbaA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); got != tt.want {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
