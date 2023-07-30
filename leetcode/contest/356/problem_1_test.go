package contest356

import "testing"

func Test_numberOfEmployeesWhoMetTarget(t *testing.T) {
	type args struct {
		hours  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				hours:  []int{0, 1, 2, 3, 4},
				target: 2,
			},
			want: 3,
		},
		{
			args: args{
				hours:  []int{5, 1, 4, 2, 2},
				target: 6,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfEmployeesWhoMetTarget(tt.args.hours, tt.args.target); got != tt.want {
				t.Errorf("numberOfEmployeesWhoMetTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
