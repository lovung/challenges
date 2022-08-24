package hard

import (
	"testing"

	"github.com/lovung/challenges/internal/linkedlist"
	"github.com/stretchr/testify/assert"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*linkedlist.ListNode[int]
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.ListNode[int]
	}{
		{
			name: "test1",
			args: args{
				lists: []*linkedlist.ListNode[int]{
					linkedlist.NewLinkedListFromValues([]int{1, 4, 5}),
					linkedlist.NewLinkedListFromValues([]int{1, 3, 4}),
					linkedlist.NewLinkedListFromValues([]int{2, 6}),
				},
			},
			want: linkedlist.NewLinkedListFromValues([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, mergeKLists(tt.args.lists), tt.want)
		})
	}
}
