package problems

import (
	"reflect"
	"testing"

	. "github.com/lovung/ds/lists"
)

func Test_oddEvenList(t *testing.T) {

	type args struct {
		head *ListNode[int]
	}
	tests := []struct {
		name string
		args args
		want *ListNode[int]
	}{
		{
			name: "test 1",
			args: args{
				head: NewLinkedListFromValues([]int{}),
			},
			want: NewLinkedListFromValues([]int{}),
		},
		{
			name: "test 2",
			args: args{
				head: NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
			},
			want: NewLinkedListFromValues([]int{1, 3, 5, 2, 4}),
		},
		{
			name: "test 3",
			args: args{
				head: NewLinkedListFromValues([]int{1, 2, 3, 4, 5, 6}),
			},
			want: NewLinkedListFromValues([]int{1, 3, 5, 2, 4, 6}),
		},
		{
			name: "test 4",
			args: args{
				head: NewLinkedListFromValues([]int{1}),
			},
			want: NewLinkedListFromValues([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
