package problems

import (
	"reflect"
	"testing"

	. "github.com/lovung/ds/lists"
)

func Test_swapNodes(t *testing.T) {
	type args struct {
		head *ListNode[int]
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode[int]
	}{
		{
			name: "test1",
			args: args{
				head: NewLinkedListFromValues([]int{1, 2, 3, 4}),
				k:    2,
			},
			want: NewLinkedListFromValues([]int{1, 3, 2, 4}),
		},
		{
			name: "test2",
			args: args{
				head: NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			want: NewLinkedListFromValues([]int{1, 4, 3, 2, 5}),
		},
		{
			name: "test3",
			args: args{
				head: NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
				k:    3,
			},
			want: NewLinkedListFromValues([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapNodes(tt.args.head, tt.args.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
