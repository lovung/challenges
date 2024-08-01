package jul2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numWaterBottles(t *testing.T) {
	assert.Equal(t, 13, numWaterBottles(9, 3))
	assert.Equal(t, 19, numWaterBottles(15, 4))
}
