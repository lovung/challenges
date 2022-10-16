package oct2022

import (
	"reflect"
	"testing"

	"github.com/lovung/ds/lists"
)

func Test_deleteMiddle(t *testing.T) {
	type args struct {
		head *lists.ListNode[int]
	}
	tests := []struct {
		name string
		args args
		want *lists.ListNode[int]
	}{
		{
			args: args{
				head: lists.NewLinkedListFromValues(
					[]int{1, 2, 3, 4, 5, 6, 7},
				),
			},
			want: lists.NewLinkedListFromValues(
				[]int{1, 2, 3, 5, 6, 7},
			),
		},
		{
			args: args{
				head: lists.NewLinkedListFromValues(
					[]int{1, 2, 3, 4},
				),
			},
			want: lists.NewLinkedListFromValues(
				[]int{1, 2, 4},
			),
		},
		{
			args: args{
				head: lists.NewLinkedListFromValues(
					[]int{1, 2},
				),
			},
			want: lists.NewLinkedListFromValues(
				[]int{1},
			),
		},
		{
			args: args{
				head: lists.NewLinkedListFromValues(
					[]int{1},
				),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
