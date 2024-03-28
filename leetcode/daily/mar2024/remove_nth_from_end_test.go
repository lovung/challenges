package mar2024

import (
	"testing"

	"github.com/lovung/ds/lists"
	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		head := lists.NewLinkedListFromValues([]int{1, 2})
		got := removeNthFromEnd(head, 1)
		assert.Equal(t, lists.NewLinkedListFromValues([]int{1}), got)
	})
	t.Run("2", func(t *testing.T) {
		head := lists.NewLinkedListFromValues([]int{1, 2})
		got := removeNthFromEnd(head, 2)
		assert.Equal(t, lists.NewLinkedListFromValues([]int{2}), got)
	})
	t.Run("3", func(t *testing.T) {
		head := lists.NewLinkedListFromValues([]int{1, 2, 3, 4, 5})
		got := removeNthFromEnd(head, 2)
		assert.Equal(t, lists.NewLinkedListFromValues([]int{1, 2, 3, 5}), got)
	})
	t.Run("4", func(t *testing.T) {
		head := lists.NewLinkedListFromValues([]int{1})
		got := removeNthFromEnd(head, 1)
		assert.Nil(t, got)
	})
}
