package mar2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxFrequencyElements(t *testing.T) {
	testCases := []struct {
		input  []int
		expect int
	}{
		{
			input:  []int{},
			expect: 0,
		},
		{
			input:  []int{1, 2, 2, 3, 1, 4},
			expect: 4,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			expect: 5,
		},
		{
			input:  []int{1, 1, 1, 1, 1},
			expect: 5,
		},
	}
	for _, tt := range testCases {
		assert.Equal(t, tt.expect, maxFrequencyElements(tt.input))
		assert.Equal(t, tt.expect, maxFrequencyElements2(tt.input))
	}
}
