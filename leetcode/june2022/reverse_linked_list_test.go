package june2022

import (
	"reflect"
	"testing"

	. "github.com/lovung/challenges/internal/linkedlist"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode[int]
	}
	tests := []struct {
		name string
		args args
		want *ListNode[int]
	}{
		{
			name: "test1",
			args: args{
				head: NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
			},
			want: NewLinkedListFromValues([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
