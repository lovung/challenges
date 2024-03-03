package topinterview150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gameOfLife(t *testing.T) {
	t.Run("test_1", func(t *testing.T) {
		board := [][]int{
			{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
		}
		gameOfLife(board)
		assert.Equal(t, [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}, board)
	})
	t.Run("test_2", func(t *testing.T) {
		board := [][]int{
			{1, 1}, {0, 1},
		}
		gameOfLife(board)
		assert.Equal(t, [][]int{{1, 1}, {1, 1}}, board)
	})
}
