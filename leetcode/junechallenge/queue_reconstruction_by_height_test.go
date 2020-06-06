package junechallenge

import (
	"reflect"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	type args struct {
		people [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test 1",
			args: args{
				people: [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			},
			want: [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
		},
		{
			name: "test 2",
			args: args{
				people: [][]int{{2, 4}, {3, 4}, {9, 0}, {0, 6}, {7, 1}, {6, 0}, {7, 3}, {2, 5}, {1, 1}, {8, 0}},
			},
			want: [][]int{{6, 0}, {1, 1}, {8, 0}, {7, 1}, {9, 0}, {2, 4}, {0, 6}, {2, 5}, {3, 4}, {7, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructQueue(tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
