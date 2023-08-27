package contest

import (
	"testing"

	"github.com/lovung/ds/lists"
	"github.com/stretchr/testify/assert"
)

func Test_doubleIt(t *testing.T) {
	assert.Equal(t, lists.NewLinkedListFromValues([]int{3, 7, 8}),
		doubleIt(lists.NewLinkedListFromValues([]int{1, 8, 9})))
	assert.Equal(t, lists.NewLinkedListFromValues([]int{1, 9, 9, 8}),
		doubleIt(lists.NewLinkedListFromValues([]int{9, 9, 9})))
}
