package contest

import (
	"strings"

	"github.com/lovung/ds/stack"
)

// Link: https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string
// Using string like a stack -> Very slow
func robotWithString(s string) string {
	t := ""
	p := strings.Builder{}
	orgLenS := len(s)

	minChar := byte('z')
	minCharAfter := make([]byte, orgLenS)
	for i := orgLenS - 1; i >= 0; i-- {
		minChar = min(minChar, s[i])
		minCharAfter[i] = minChar
	}

	i := 0
	for orgLenS > i || len(t) > 0 {
		if i == orgLenS {
			p.WriteString(revertString(t))
			break
		}
		if len(t) == 0 || t[len(t)-1] > minCharAfter[i] {
			t += string(s[i])
			i++
			continue
		}
		p.WriteByte(t[len(t)-1])
		t = t[:len(t)-1]
	}
	return p.String()
}

func revertString(s string) string {
	sb := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

// Using stack
func robotWithString2(s string) string {
	t := stack.NewArrayStack[byte](1e5)
	p := strings.Builder{}
	orgLenS := len(s)

	minChar := byte('z')
	minCharAfter := make([]byte, orgLenS)
	for i := orgLenS - 1; i >= 0; i-- {
		minChar = min(minChar, s[i])
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
