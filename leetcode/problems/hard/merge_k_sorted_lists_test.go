package hard

import (
	"testing"

	. "github.com/lovung/ds/lists"
	"github.com/stretchr/testify/assert"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode[int]
	}
	tests := []struct {
		name string
		args args
		want *ListNode[int]
	}{
		{
			name: "test1",
			args: args{
				lists: []*ListNode[int]{
					NewLinkedListFromValues([]int{1, 4, 5}),
					NewLinkedListFromValues([]int{1, 3, 4}),
					NewLinkedListFromValues([]int{2, 6}),
					nil,
				},
			},
			want: NewLinkedListFromValues([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name: "test2",
			args: args{
				lists: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, mergeKLists(tt.args.lists), tt.want)
		})
	}
}
