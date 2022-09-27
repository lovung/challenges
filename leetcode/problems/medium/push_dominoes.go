package medium

import "strings"

// Link: https://leetcode.com/problems/push-dominoes/
func pushDominoes(dominoes string) string {
	// if a domino is L, R 	-> stay L, R
	// if a domino is "."
	//		->	if on the left of a L only -> L
	//		->	else if on the right of a R only -> R
	//		-> 	else if both, the same with nearest one
	//			-> if the distances are equal, stay "."
	//		-> 	else stay "."
	// S1: Bruce force: O(N^2)
	// S2: Expand from L and R, count the distance and save to array: O(N)
	disFromL := buildDisFromL(dominoes)
	disFromR := buildDisFromR(dominoes)
	sb := strings.Builder{}
	for i := range dominoes {
		if dominoes[i] != '.' {
			sb.WriteByte(dominoes[i])
			continue
		}
		switch {
		case disFromL[i] > 0 && disFromR[i] > 0:
			switch {
			case disFromL[i] < disFromR[i]:
				sb.WriteByte('L')
			case disFromL[i] > disFromR[i]:
				sb.WriteByte('R')
			default:
				sb.WriteByte('.')
			}
		case disFromL[i] > 0:
			sb.WriteByte('L')
		case disFromR[i] > 0:
			sb.WriteByte('R')
		default:
			sb.WriteByte('.')
		}
	}
	return sb.String()
}

func buildDisFromL(dominoes string) []int {
	disFromL := make([]int, len(dominoes))
	cnt := 0
	fall := false
	for i := len(dominoes) - 1; i >= 0; i-- {
		switch dominoes[i] {
		case 'L':
			fall = true
			cnt = 0
		case 'R':
			fall = false
			cnt = 0
		default:
			disFromL[i] = cnt
		}
		if fall {
			cnt++
		}
	}
	return disFromL
}

func buildDisFromR(dominoes string) []int {
	disFromR := make([]int, len(dominoes))
	cnt := 0
	fall := false
	for i := range dominoes {
		switch dominoes[i] {
		case 'R':
			fall = true
			cnt = 0
		case 'L':
			fall = false
			cnt = 0
		default:
			disFromR[i] = cnt
		}
		if fall {
			cnt++
		}
	}
	return disFromR
}
