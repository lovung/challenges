package medium

// Link: https://leetcode.com/problems/utf-8-validation/
const (
	wait = iota
	len1Byte
	len2Bytes
	len3Bytes
	len4Bytes
)

const (
	mask1Byte  = 1 << 7
	tailMask   = mask1Byte | 1<<6
	mask2Bytes = tailMask | 1<<5
	mask3Bytes = mask2Bytes | 1<<4
	mask4Bytes = mask3Bytes | 1<<3

	val1Byte  = 0
	valTail   = 0b10000000
	val2Bytes = 0b11000000
	val3Bytes = 0b11100000
	val4Bytes = 0b11110000
)

func validUtf8(data []int) bool {
	state := wait
	remainLen := 0
	for i := range data {
		switch state {
		case wait:
			switch {
			case data[i]&mask1Byte == val1Byte:
				state = wait
			case data[i]&mask2Bytes == val2Bytes:
				state = len2Bytes
				remainLen = 1
			case data[i]&mask3Bytes == val3Bytes:
				state = len3Bytes
				remainLen = 2
			case data[i]&mask4Bytes == val4Bytes:
				state = len4Bytes
				remainLen = 3
			default:
				return false
			}
		case len2Bytes, len3Bytes, len4Bytes:
			if data[i]&tailMask != valTail {
				return false
			}
			remainLen--
			if remainLen == 0 {
				state = wait
			}
		}
	}
	return state == wait
}
