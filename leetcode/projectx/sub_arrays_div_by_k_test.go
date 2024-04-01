package projectx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subarraysDivByK(t *testing.T) {
	assert.Equal(t, 7, subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
	assert.Equal(t, 0, subarraysDivByK([]int{5}, 9))
}
