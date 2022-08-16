package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstBadVersion(t *testing.T) {
	assert.Equal(t, 4, firstBadVersion(5))
	assert.Equal(t, 4, firstBadVersion(6))
	assert.Equal(t, 4, firstBadVersion(7))
	assert.Equal(t, 4, firstBadVersion(8))

	var oldFunc = isBadVersion
	isBadVersion = func(k int) bool { return k >= 1 }
	defer func() {
		isBadVersion = oldFunc
	}()
	assert.Equal(t, 1, firstBadVersion(1))
}
