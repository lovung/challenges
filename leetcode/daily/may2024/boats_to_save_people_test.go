package may2024

import "testing"

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				people: []int{3, 2, 2, 1},
				limit:  3,
			},
			want: 3,
		},
		{
			args: args{
				people: []int{2, 1},
				limit:  3,
			},
			want: 1,
		},
		{
			args: args{
				people: []int{3, 5, 3, 4},
				limit:  5,
			},
			want: 4,
		},
		{
			args: args{
				people: []int{3, 5, 4, 2},
				limit:  7,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.args.people, tt.args.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}
