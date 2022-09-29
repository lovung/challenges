package grind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianFinder(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		obj := NewMedianFinder()
		obj.AddNum(1)
		assert.Equal(t, 1.0, obj.FindMedian())
		obj.AddNum(2)
		assert.Equal(t, 1.5, obj.FindMedian())
		obj.AddNum(3)
		assert.Equal(t, 2.0, obj.FindMedian())
	})
	t.Run("normal", func(t *testing.T) {
		obj := NewMedianFinder()
		obj.AddNum(6)
		assert.Equal(t, 6.0, obj.FindMedian())
		obj.AddNum(10)
		assert.Equal(t, 8.0, obj.FindMedian())
		obj.AddNum(2)
		assert.Equal(t, 6.0, obj.FindMedian())
		obj.AddNum(6)
		assert.Equal(t, 6.0, obj.FindMedian())
		obj.AddNum(5)
		assert.Equal(t, 6.0, obj.FindMedian())
		obj.AddNum(3)
		assert.Equal(t, 5.5, obj.FindMedian())
	})
	t.Run("normal", func(t *testing.T) {
		obj := NewMedianFinder()
		obj.AddNum(100000)
		obj.AddNum(99999)
		obj.AddNum(99998)
		obj.AddNum(99997)
		obj.AddNum(99996)
		obj.AddNum(99995)
		obj.AddNum(99994)
		assert.Equal(t, 99997.0, obj.FindMedian())
	})
}
