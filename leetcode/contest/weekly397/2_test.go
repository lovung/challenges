package weekly397

import "testing"

func Test_maximumEnergy(t *testing.T) {
	type args struct {
		energy []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				energy: []int{5, 2, -10, -5, 1},
				k:      3,
			},
			want: 3,
		},
		{
			args: args{
				energy: []int{-2, -3, -1},
				k:      2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumEnergy(tt.args.energy, tt.args.k); got != tt.want {
				t.Errorf("maximumEnergy() = %v, want %v", got, tt.want)
			}
		})
	}
}
