package medium

import (
	"strings"

	"github.com/lovung/ds/slice"
)

// Link: https://leetcode.com/problems/longest-palindromic-substring/
// Solution: Manacher's Algorithm
func longestPalindrome(s string) string {
	length := len(s)*2 + 1
	modifiedStr := modifyString(s)
	PalindromeRadii := make([]int, length)
	var center, radius int
	for center < length {
		for center-(radius+1) >= 0 && center+(radius+1) < length && modifiedStr[center-(radius+1)] == modifiedStr[center+(radius+1)] {
			radius++
		}
		PalindromeRadii[center] = radius

		oldCenter := center
		oldRadius := radius
		center++
		radius = 0
		for center <= oldCenter+oldRadius {
			mirroredCenter := oldCenter - (center - oldCenter)
			maxMirroredRadius := oldCenter + oldRadius - center
			if PalindromeRadii[mirroredCenter] < maxMirroredRadius {
				PalindromeRadii[center] = PalindromeRadii[mirroredCenter]
				center++
			} else if PalindromeRadii[mirroredCenter] > maxMirroredRadius {
				PalindromeRadii[center] = maxMirroredRadius
				center++
			} else {
				radius = maxMirroredRadius
				break
			}
		}
	}

	maxLen, maxCenter := slice.MaxValueAndIndex(PalindromeRadii)
	return removeRune(modifiedStr[maxCenter-maxLen : maxCenter+maxLen])
}

func longestPalindrome2(s string) string {
	length := len(s)*2 + 1
	modifiedStr := modifyString(s)
	PalindromeRadii := make([]int, length)
	var center, radius, maxLen, maxCenter, mirror int
	for i := 0; i < length; i++ {
		mirror = center + center - i
		if i < radius {
			PalindromeRadii[i] = min(radius-i, PalindromeRadii[mirror])
		}

		l := i - (1 + PalindromeRadii[i])
		r := i + (1 + PalindromeRadii[i])
		for r < length && l >= 0 && modifiedStr[r] == modifiedStr[l] {
			PalindromeRadii[i]++
			r++
			l--
		}

		if i+PalindromeRadii[i] > radius {
			center = i
			radius = i + PalindromeRadii[i]

			if PalindromeRadii[i] > maxLen {
				maxLen = PalindromeRadii[i]
				maxCenter = center
			}
		}
	}
	return removeRune(modifiedStr[maxCenter-maxLen : maxCenter+maxLen])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func modifyString(s string) string {
	var out strings.Builder
	for _, e := range s {
		out.WriteRune('#')
		out.WriteRune(e)
	}
	out.WriteRune('#')
	return out.String()
}

func removeRune(s string) string {
	var out strings.Builder
	for _, e := range s {
		if e != '#' {
			out.WriteRune(e)
		}
	}
	return out.String()
}
