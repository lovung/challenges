package mar2024

import (
	"testing"

	"github.com/lovung/ds/lists"
	"github.com/stretchr/testify/assert"
)

func Test_mergeInBetween(t *testing.T) {
	type args struct {
		list1 []int
		a     int
		b     int
		list2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				list1: []int{10, 1, 13, 6, 9, 5},
				a:     3, b: 4,
				list2: []int{1000000, 1000001, 1000002},
			},
			want: []int{10, 1, 13, 1000000, 1000001, 1000002, 5},
		},
		{
			args: args{
				list1: []int{0, 1, 2, 3, 4, 5, 6},
				a:     2, b: 5,
				list2: []int{1000000, 1000001, 1000002, 1000003, 1000004},
			},
			want: []int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeInBetween(
				lists.NewLinkedListFromValues(tt.args.list1),
				tt.args.a, tt.args.b,
				lists.NewLinkedListFromValues(tt.args.list2))
			assert.Equal(t, lists.NewLinkedListFromValues(tt.want), got)
		})
	}
}
