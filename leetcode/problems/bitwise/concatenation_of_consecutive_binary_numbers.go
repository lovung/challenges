package bitwise

import (
	"strings"
)

// Link: https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
const mod = 1e9 + 7
const twoMod = (uint64(1) << 63 % mod) << 1

func concatenatedBinary(n int) int {
	sb := strings.Builder{}
	for i := 1; i <= n; i++ {
		toBinaryString(uint64(i), &sb)
	}
	res := uint64(0)
	str := sb.String()
	k := uint64(1)
	for checkedLen := 0; checkedLen < len(str); checkedLen += 64 {
		l, r := max(0, len(str)-checkedLen-64), len(str)-checkedLen
		res += binaryStringToUint64(str[l:r]) * k
		res %= mod
		k *= twoMod
		k %= mod
	}
	return int(res % mod)
}

func toBinaryString(n uint64, sb *strings.Builder) {
	catchFirstBitOne := false
	for mask := uint64(1) << 63; mask > 0; mask >>= 1 {
		if n&mask > 0 {
			sb.WriteRune('1')
			catchFirstBitOne = true
		} else if catchFirstBitOne {
			sb.WriteRune('0')
		}
	}
}

func binaryStringToUint64(s string) uint64 {
	val := uint64(0)
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			val |= 1 << (n - i - 1)
		}
	}
	return val % mod
}

func concatenatedBinary2(n int) int {
	ans := 1
	for i := 1; i < n; i++ {
		ans *= 1 << countBits(i+1)
		ans |= i + 1
		ans %= mod
	}
	return ans
}

func countBits(n int) int {
	i := 0
	for ; n > 0; i++ {
		n >>= 1
	}
	return i
}
