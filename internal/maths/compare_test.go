package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxMin(t *testing.T) {
	// int
	assert.Equal(t, Max(1, 2, 3), 3)
	assert.Equal(t, Max(1, 2, 3, 4, 5), 5)
	assert.Equal(t, Min(1, 2, 3), 1)
	assert.Equal(t, Min(1, 2, 3, 4, 5), 1)

	// float
	assert.Equal(t, Max(1.1, 2.2, 3.3), 3.3)
	assert.Equal(t, Max(1.1, 2.2, 3.3, 4.4, 5.5), 5.5)
	assert.Equal(t, Min(1.1, 2.2, 3.3), 1.1)
	assert.Equal(t, Min(1.1, 2.2, 3.3, 4.4, 5.5), 1.1)

	// string
	assert.Equal(t, Max("a", "b", "c"), "c")
	assert.Equal(t, Max("ab", "bc", "cd", "dec", "ef"), "ef")
	assert.Equal(t, Min("a", "b", "c"), "a")
	assert.Equal(t, Min("b", "c", "d", "e", "abcd"), "abcd")
}
