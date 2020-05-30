package maychallenge

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		K      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test 1",
			args: args{
				points: [][]int{{1, 3}, {-2, 2}},
				K:      1,
			},
			want: [][]int{{-2, 2}},
		},
		{
			name: "test 2",
			args: args{
				points: [][]int{{1, 3}, {-2, 2}, {0, 1}},
				K:      2,
			},
			want: [][]int{{-2, 2}, {0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
