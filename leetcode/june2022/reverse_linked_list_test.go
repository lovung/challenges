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
			name: "normal",
			args: args{
				head: NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
			},
			want: NewLinkedListFromValues([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "nil",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "1-item",
			args: args{
				head: NewLinkedListFromValues([]int{1}),
			},
			want: NewLinkedListFromValues([]int{1}),
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

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode[int]
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode[int]
	}{
		{
			name: "normal",
			args: args{
				head:  NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
				left:  2,
				right: 4,
			},
			want: NewLinkedListFromValues([]int{1, 4, 3, 2, 5}),
		},
		{
			name: "left == right",
			args: args{
				head:  NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
				left:  3,
				right: 3,
			},
			want: NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "right-at-last",
			args: args{
				head:  NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
				left:  3,
				right: 5,
			},
			want: NewLinkedListFromValues([]int{1, 2, 5, 4, 3}),
		},
		{
			name: "left-at-first",
			args: args{
				head:  NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
				left:  1,
				right: 5,
			},
			want: NewLinkedListFromValues([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
