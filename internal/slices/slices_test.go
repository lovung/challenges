package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		in  []any
		out []any
	}{
		{
			in:  []any{1, 2, 3},
			out: []any{3, 2, 1},
		},
		{
			in:  []any{"a", "b", "c"},
			out: []any{"c", "b", "a"},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			Reverse(tt.in)
			assert.Equal(t, tt.out, tt.in)
		})
	}
}
