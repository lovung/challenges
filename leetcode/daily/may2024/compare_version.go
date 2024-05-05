package may2024

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/compare-version-numbers/
func compareVersion(version1 string, version2 string) int {
	version1s := convertVersion(version1)
	version2s := convertVersion(version2)
	for i, j := 0, 0; i < len(version1s) || j < len(version2s); i, j = i+1, j+1 {
		v1, v2 := 0, 0
		if i < len(version1s) {
			v1 = version1s[i]
		}
		if j < len(version2s) {
			v2 = version2s[j]
		}
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	return 0
}

func convertVersion(v string) []int {
	parts := strings.Split(v, ".")
	versions := make([]int, 0, len(parts))
	for i := range parts {
		_v, _ := strconv.Atoi(parts[i])
		versions = append(versions, _v)
	}
	return versions
}
