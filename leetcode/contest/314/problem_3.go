package contest

import (
	"strings"

	"github.com/lovung/ds/maths"
	"github.com/lovung/ds/stack"
)

// Link: https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string
func robotWithString(s string) string {
	t := stack.NewArrayStack[byte](1e5)
	p := strings.Builder{}
	orgLenS := len(s)

	minChar := byte('z')
	minCharAfter := make([]byte, orgLenS)
	for i := orgLenS - 1; i >= 0; i-- {
		minChar = maths.Min(minChar, s[i])
		minCharAfter[i] = minChar
	}

	i := 0
	for orgLenS > i || t.Len() > 0 {
		if i == orgLenS {
			break
		}
		if t.Len() == 0 || t.Peek() > minCharAfter[i] {
			t.Push(s[i])
			i++
			continue
		}
		p.WriteByte(t.Pop())
	}
	for t.Len() > 0 {
		p.WriteByte(t.Pop())
	}
	return p.String()
}
