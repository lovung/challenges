package mar2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumLength(t *testing.T) {
	assert.Equal(t, 2, minimumLength("ca"))
	assert.Equal(t, 0, minimumLength("aaaaaaa"))
	assert.Equal(t, 0, minimumLength("cabaabac"))
	assert.Equal(t, 3, minimumLength("aabccabba"))
	assert.Equal(t, 1, minimumLength("bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb"))
	assert.Equal(t, 1, minimumLength("aba"))
	assert.Equal(t, 1, minimumLength("a"))
}
