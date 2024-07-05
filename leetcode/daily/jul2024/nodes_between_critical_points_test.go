package jul2024

import (
	"reflect"
	"testing"

	"github.com/lovung/ds/lists"
)

func Test_nodesBetweenCriticalPoints(t *testing.T) {
	type args struct {
		head *lists.ListNode[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				head: lists.NewLinkedListFromValues([]int{0, 3, 1, 0, 4, 5, 2, 0}),
			},
			want: []int{2, 4},
		},
		{
			args: args{
				head: lists.NewLinkedListFromValues([]int{3, 1}),
			},
			want: []int{-1, -1},
		},
		{
			args: args{
				head: lists.NewLinkedListFromValues([]int{5, 3, 1, 2, 5, 1, 2}),
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nodesBetweenCriticalPoints(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodesBetweenCriticalPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
