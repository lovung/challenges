package mar2024

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lovung/ds/lists"
)

func Test_removeZeroSumSublists(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{
			head: []int{1, 2, -3, 3, 1},
			want: []int{3, 1},
		},
		{
			head: []int{1, 2, 3, -3, 4},
			want: []int{1, 2, 4},
		},
		{
			head: []int{1, 2, 3, -3, -2},
			want: []int{1},
		},
		{
			head: []int{-6, 1, 4, -3, -1, 4},
			want: []int{-6, 1, 4},
		},
		{
			head: []int{-6, 1, 4, -3, -1, 5},
			want: []int{},
		},
		{
			head: []int{1, 3, 2, -3, -2, 5, 5, -5, 1},
			want: []int{1, 5, 1},
		},
		{
			head: []int{2, 2, -2, 1, -1, -1},
			want: []int{2, -1},
		},
		{
			head: []int{10, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1},
			want: []int{10},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("solution_1_case_%d", i), func(t *testing.T) {
			head := lists.NewLinkedListFromValues(tt.head)
			want := lists.NewLinkedListFromValues(tt.want)
			if got := removeZeroSumSublists(head); !reflect.DeepEqual(got, want) {
				t.Errorf("removeZeroSumSublists() = %v, want %v", got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("solution_2_case_%d", i), func(t *testing.T) {
			head := lists.NewLinkedListFromValues(tt.head)
			want := lists.NewLinkedListFromValues(tt.want)
			if got := removeZeroSumSublists2(head); !reflect.DeepEqual(got, want) {
				t.Errorf("removeZeroSumSublists2() = %v, want %v", got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("solution_3_case_%d", i), func(t *testing.T) {
			head := lists.NewLinkedListFromValues(tt.head)
			want := lists.NewLinkedListFromValues(tt.want)
			if got := removeZeroSumSublists3(head); !reflect.DeepEqual(got, want) {
				t.Errorf("removeZeroSumSublists3() = %v, want %v", got, tt.want)
			}
		})
	}
}
