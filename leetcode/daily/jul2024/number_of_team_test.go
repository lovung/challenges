package jul2024

import "testing"

func Test_numTeams(t *testing.T) {
	type args struct {
		rating []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				rating: []int{2, 5, 3, 4, 1},
			},
			want: 3,
		},
		{
			args: args{
				rating: []int{6, 2, 5, 3, 4, 1},
			},
			want: 9,
		},
		{
			args: args{
				rating: []int{2, 1, 3},
			},
			want: 0,
		},
		{
			args: args{
				rating: []int{1, 2, 3, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTeams(tt.args.rating); got != tt.want {
				t.Errorf("numTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
