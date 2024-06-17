package biweekly132

import "testing"

func Test_findWinningPlayer(t *testing.T) {
	type args struct {
		skills []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				skills: []int{4, 2, 6, 3, 9},
				k:      2,
			},
			want: 2,
		},
		{
			args: args{
				skills: []int{2, 5, 4},
				k:      3,
			},
			want: 1,
		},
		{
			args: args{
				skills: []int{2, 5, 4},
				k:      123123,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWinningPlayer1(tt.args.skills, tt.args.k); got != tt.want {
				t.Errorf("findWinningPlayer1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := findWinningPlayer(tt.args.skills, tt.args.k); got != tt.want {
				t.Errorf("findWinningPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
