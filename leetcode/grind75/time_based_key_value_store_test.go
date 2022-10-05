package grind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeMap_Set(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		this := NewTimeMap()
		assert.Equal(t, "", this.Get("foo", 1))
		this.Set("foo", "bar", 1)
		assert.Equal(t, "bar", this.Get("foo", 1))
		assert.Equal(t, "bar", this.Get("foo", 3))
		this.Set("foo", "bar2", 4)
		assert.Equal(t, "bar2", this.Get("foo", 4))
		assert.Equal(t, "bar2", this.Get("foo", 5))
		assert.Equal(t, "bar", this.Get("foo", 3))
		assert.Equal(t, "", this.Get("foo", 0))
		this.Set("foo", "bar3", 6)
		assert.Equal(t, "bar3", this.Get("foo", 7))
	})
}
