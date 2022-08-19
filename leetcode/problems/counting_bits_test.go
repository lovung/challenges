package problems

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				num: 2,
			},
			want: []int{0, 1, 1},
		},
		{
			name: "test 1",
			args: args{
				num: 5,
			},
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
