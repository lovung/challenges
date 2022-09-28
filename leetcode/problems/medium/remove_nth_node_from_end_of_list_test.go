package medium

import (
	"reflect"
	"testing"

	"github.com/lovung/ds/lists"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *lists.ListNode[int]
		n    int
	}
	tests := []struct {
		name string
		args args
		want *lists.ListNode[int]
	}{
		{
			args: args{
				head: lists.NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: lists.NewLinkedListFromValues([]int{1, 2, 3, 5}),
		},
		{
			args: args{
				head: lists.NewLinkedListFromValues([]int{1, 2}),
				n:    1,
			},
			want: lists.NewLinkedListFromValues([]int{1}),
		},
		{
			args: args{
				head: lists.NewLinkedListFromValues([]int{1}),
				n:    1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
