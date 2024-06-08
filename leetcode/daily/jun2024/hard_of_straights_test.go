package jun2024

import "testing"

func Test_isNStraightHand(t *testing.T) {
	type args struct {
		hand      []int
		groupSize int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				hand:      []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
				groupSize: 3,
			},
			want: true,
		},
		{
			args: args{
				hand:      []int{1, 2, 3, 6, 2, 3, 4, 7, 9},
				groupSize: 3,
			},
			want: false,
		},
		{
			args: args{
				hand:      []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
				groupSize: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNStraightHand(tt.args.hand, tt.args.groupSize); got != tt.want {
				t.Errorf("isNStraightHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
