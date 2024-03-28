package topinterview150

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_topKFrequent(t *testing.T) {
	t.Run("test_1", func(t *testing.T) {
		ret := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
		sort.Ints(ret)
		assert.Equal(t, []int{1, 2}, ret)
	})
	t.Run("test_2", func(t *testing.T) {
		ret := topKFrequent([]int{1}, 1)
		sort.Ints(ret)
		assert.Equal(t, []int{1}, ret)
	})
	t.Run("test_3", func(t *testing.T) {
		ret := topKFrequent([]int{1, 2, 2, 3, 3, 3, 4}, 2)
		sort.Ints(ret)
		assert.Equal(t, []int{2, 3}, ret)
	})
}
