package algorithm

import (
	"reflect"
	"testing"

	"github.com/lovung/ds/lists"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *lists.ListNode[int]
	}
	tests := []struct {
		name string
		args args
		want *lists.ListNode[int]
	}{
		{
			name: "odd",
			args: args{
				head: lists.NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
			},
			want: lists.NewLinkedListFromValues([]int{3, 4, 5}),
		},
		{
			name: "even",
			args: args{
				head: lists.NewLinkedListFromValues([]int{1, 2, 3, 4, 5, 6}),
			},
			want: lists.NewLinkedListFromValues([]int{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
