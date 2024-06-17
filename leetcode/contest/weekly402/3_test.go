package weekly402

import "testing"

func Test_maximumTotalDamage(t *testing.T) {
	type args struct {
		power []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				power: []int{3, 4, 8, 10, 8, 8, 3},
			},
			want: 30,
		},
		{
			args: args{
				power: []int{7, 1, 6, 6},
			},
			want: 13,
		},
		{
			args: args{
				power: []int{5, 9, 2, 10, 2, 7, 10, 9, 3, 8},
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTotalDamage(tt.args.power); got != tt.want {
				t.Errorf("maximumTotalDamage() = %v, want %v", got, tt.want)
			}
		})
	}
}
