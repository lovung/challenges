package mar2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumOddBinaryNumber(t *testing.T) {
	assert.Equal(t, "001", maximumOddBinaryNumber("010"))
	assert.Equal(t, "1001", maximumOddBinaryNumber("0101"))
	assert.Equal(t, "1111100000001", maximumOddBinaryNumber("0000000111111"))
}
