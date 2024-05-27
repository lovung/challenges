package biweekly131

import (
	"reflect"
	"testing"
)

func Test_occurrencesOfElement(t *testing.T) {
	type args struct {
		nums    []int
		queries []int
		x       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums:    []int{1, 3, 1, 7},
				queries: []int{1, 3, 2, 4},
				x:       1,
			},
			want: []int{0, -1, 2, -1},
		},
		{
			args: args{
				nums:    []int{1, 2, 3},
				queries: []int{10},
				x:       5,
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := occurrencesOfElement(tt.args.nums, tt.args.queries, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("occurrencesOfElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
