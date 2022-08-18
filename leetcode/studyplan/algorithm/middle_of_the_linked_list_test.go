package algorithm

import (
	"reflect"
	"testing"

	. "github.com/lovung/challenges/internal/linkedlist"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode[int]
	}
	tests := []struct {
		name string
		args args
		want *ListNode[int]
	}{
		{
			name: "odd",
			args: args{
				head: NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
			},
			want: NewLinkedListFromValues([]int{3, 4, 5}),
		},
		{
			name: "even",
			args: args{
				head: NewLinkedListFromValues([]int{1, 2, 3, 4, 5, 6}),
			},
			want: NewLinkedListFromValues([]int{4, 5, 6}),
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
