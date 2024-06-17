package jun2024

import "testing"

func Test_minMovesToSeat(t *testing.T) {
	type args struct {
		seats    []int
		students []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				seats:    []int{3, 1, 5},
				students: []int{2, 7, 4},
			},
			want: 4,
		},
		{
			args: args{
				seats:    []int{4, 1, 5, 9},
				students: []int{1, 3, 2, 6},
			},
			want: 7,
		},
		{
			args: args{
				seats:    []int{2, 2, 6, 6},
				students: []int{1, 3, 2, 6},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMovesToSeat(tt.args.seats, tt.args.students); got != tt.want {
				t.Errorf("minMovesToSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
