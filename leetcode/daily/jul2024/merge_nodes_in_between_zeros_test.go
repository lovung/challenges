package jul2024

import (
	"testing"

	"github.com/lovung/ds/lists"
	"github.com/stretchr/testify/assert"
)

func Test_mergeNodes(t *testing.T) {
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
				head: lists.NewLinkedListFromValues([]int{0, 3, 1, 0, 4, 5, 2, 0}),
			},
			want: lists.NewLinkedListFromValues([]int{4, 11}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeNodes(tt.args.head)
			assert.Equal(t, tt.want, got)
		})
	}
}
