package contest

import "strings"

// Link: https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/
func garbageCollection(garbage []string, travel []int) int {
	const M, P, G = "M", "P", "G"
	sumTravels := make([]int, len(travel)+1)
	for i := range travel {
		sumTravels[i+1] = sumTravels[i] + travel[i]
	}
	timeG, timeP, timeM := 0, 0, 0
	idxG, idxP, idxM := 0, 0, 0
	for i := range garbage {
		if c := strings.Count(garbage[i], M); c > 0 {
			timeM += sumTravels[i] - sumTravels[idxM] + c
			idxM = i
		}
		if c := strings.Count(garbage[i], G); c > 0 {
			timeG += sumTravels[i] - sumTravels[idxG] + c
			idxG = i
		}
		if c := strings.Count(garbage[i], P); c > 0 {
			timeP += sumTravels[i] - sumTravels[idxP] + c
			idxP = i
		}
	}
	return timeG + timeM + timeP
}
