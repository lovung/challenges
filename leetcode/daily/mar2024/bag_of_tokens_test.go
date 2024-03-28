package mar2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bagOfTokensScore(t *testing.T) {
	assert.Equal(t, 0, bagOfTokensScore([]int{100}, 50))
	assert.Equal(t, 1, bagOfTokensScore([]int{100, 200}, 150))
	assert.Equal(t, 2, bagOfTokensScore([]int{100, 200, 300, 400}, 200))
	assert.Equal(t, 1, bagOfTokensScore([]int{100, 300, 300, 400}, 200))
	assert.Equal(t, 3, bagOfTokensScore([]int{100, 100, 100, 100, 400}, 100))
	assert.Equal(t, 5, bagOfTokensScore([]int{6, 0, 39, 52, 45, 49, 59, 68, 42, 37}, 99))
}

// tokens  | p | s
// 1 2 3 4 | 2 | 0 -> l = 2, r = 3 is optimal
// x 2 3 4 | 1 | 1 -> l = 2, r = 3 is optimal
// x 2 3 x | 5 | 0 -> l = 2, r = 3 is optimal

// 0 6 37 39 42  45  49  . 52  59  68  | 99 | 0
// // 0 6 43 82 124 169 218 . 270 329 397 | 496 / 2 = 248
// x 6 37 39 42  45  49  . 52  59  68  | 99 | 1
// x x 37 39 42  45  49  . 52  59  68  | 93 | 2
// x x x  39 42  45  49  . 52  59  68  | 56 | 3
// x x x  x  42  45  49  . 52  59  68  | 17 | 4
// x x x  x  42  45  49  . 52  59  x   | 86 | 3
// x x x  x  x   45  49  . 52  59  x   | 44 | 4
// x x x  x  x   45  49    52 .x   x   | 103| 3
// x x x  x  x   x   49    52 .x   x   | 58 | 4
// x x x  x  x   x   x     52 .x   x   | 9  | 5
