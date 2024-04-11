package apr2024

import "github.com/lovung/ds/slice"

// https://leetcode.com/problems/remove-k-digits/
func removeKdigits1(num string, k int) string {
	numByte := []byte(num)
	for i := 0; i < k; i++ {
		numByte = remove1digits(numByte)
		numByte = clean(numByte)
	}
	return string(numByte)
}

func removeKdigits2(num string, k int) string {
	res := make([]byte, 0, len(num))
	for i := 0; i < len(num); i++ {
		for k > 0 && len(res) > 0 && res[len(res)-1] > num[i] {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, num[i])
	}
	if k <= len(res) {
		res = res[:len(res)-k]
	}
	return string(clean(res))
}

func clean(num []byte) []byte {
	for i := range num {
		if num[i] != '0' {
			return num[i:]
		}
	}
	return []byte{'0'}
}

func remove1digits(num []byte) []byte {
	for i := 0; i < len(num)-1; i++ {
		if num[i] > num[i+1] {
			return slice.RemoveAt(num, i)
		}
	}
	return slice.RemoveAt(num, len(num)-1)
}
