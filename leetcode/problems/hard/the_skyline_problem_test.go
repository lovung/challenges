package hard

import (
	"reflect"
	"testing"
)

func Test_getSkyline(t *testing.T) {
	type args struct {
		buildings [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				buildings: [][]int{{2, 9, 10}, {3, 9, 15}, {5, 9, 12}},
			},
			want: [][]int{{2, 10}, {3, 15}, {9, 0}},
		},
		{
			args: args{
				buildings: [][]int{{2, 9, 10}, {3, 8, 15}, {4, 7, 20}},
			},
			want: [][]int{{2, 10}, {3, 15}, {4, 20}, {7, 15}, {8, 10}, {9, 0}},
		},
		{
			args: args{
				buildings: [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
			},
			want: [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSkyline(tt.args.buildings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSkyline() = %v, want %v", got, tt.want)
			}
		})
	}
}
