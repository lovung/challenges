package june2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasAlternatingBits(t *testing.T) {
	assert.True(t, hasAlternatingBits(5))
	assert.False(t, hasAlternatingBits(7))
	assert.False(t, hasAlternatingBits(11))
}
