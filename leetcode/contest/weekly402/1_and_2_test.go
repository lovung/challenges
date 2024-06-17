package weekly402

import "testing"

func Test_countCompleteDayPairs(t *testing.T) {
	type args struct {
		hours []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				hours: []int{11, 13, 30, 24, 24},
			},
			want: 2,
		},
		{
			args: args{
				hours: []int{12, 12, 30, 24, 24},
			},
			want: 2,
		},
		{
			args: args{
				hours: []int{72, 48, 24, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompleteDayPairs(tt.args.hours); got != tt.want {
				t.Errorf("countCompleteDayPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
