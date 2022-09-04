package contest

import "testing"

func Test_checkDistances(t *testing.T) {
	type args struct {
		s        string
		distance []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s:        "aa",
				distance: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: false,
		},
		{
			args: args{
				s:        "abaccb",
				distance: []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDistances(tt.args.s, tt.args.distance); got != tt.want {
				t.Errorf("checkDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
