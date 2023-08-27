package contest

import (
	"testing"

	"github.com/lovung/ds/lists"
	"github.com/stretchr/testify/assert"
)

func Test_insertGreatestCommonDivisors(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		input := lists.NewLinkedListFromValues([]int{18, 6, 10, 3})
		expected := lists.NewLinkedListFromValues([]int{18, 6, 6, 2, 10, 1, 3})
		assert.Equal(t, expected, insertGreatestCommonDivisors(input))
	})
	t.Run("2", func(t *testing.T) {
		input := lists.NewLinkedListFromValues([]int{7})
		expected := lists.NewLinkedListFromValues([]int{7})
		assert.Equal(t, expected, insertGreatestCommonDivisors(input))
	})
}
