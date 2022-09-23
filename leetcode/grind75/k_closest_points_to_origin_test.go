package grind75

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				points: [][]int{{1, 3}, {-2, 2}},
				k:      1,
			},
			want: [][]int{{-2, 2}},
		},
		{
			args: args{
				points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
				k:      2,
			},
			want: [][]int{{3, 3}, {-2, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
