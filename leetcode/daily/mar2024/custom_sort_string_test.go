package mar2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_customSortString(t *testing.T) {
	t.Run("sol_1", func(t *testing.T) {
		assert.Equal(t, "cbad", customSortString("cba", "abcd"))
		assert.Equal(t, "bcad", customSortString("bcafg", "abcd"))
	})

	t.Run("sol_2", func(t *testing.T) {
		assert.Equal(t, "cbad", customSortString2("cba", "abcd"))
		assert.Equal(t, "bcad", customSortString2("bcafg", "abcd"))
	})
}
