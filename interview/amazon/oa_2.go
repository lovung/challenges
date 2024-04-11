package amazon

import "slices"

// Solution 1: Out of memory
func getMaximumPoints1(days []int32, k int32) int64 {
	n := daysOfTournament(days)
	twoTournamentPoints := make([]int32, int(n)+int(k))
	buildPoints(twoTournamentPoints, days)
	copy(twoTournamentPoints[n:], twoTournamentPoints[:n])
	sum := int64(0)
	for i := 0; i < int(k); i++ {
		sum += int64(twoTournamentPoints[i])
	}
	res := sum
	for l, r := 0, int(k); r < int(n)+int(k); r++ {
		sum += int64(twoTournamentPoints[r]) - int64(twoTournamentPoints[l])
		res = max(res, sum)
		l++
	}
	return res
}

func buildPoints(points, days []int32) []int32 {
	i := 0
	for _, day := range days {
		for j := int32(1); j <= day; j++ {
			points[i] = j
			i++
		}
	}
	return points
}

// Solution 2: Timeout
func getMaximumPoints2(days []int32, k int32) int64 {
	n := daysOfTournament(days)
	sum := int64(0)
	cache := make(map[int]int32)
	for i := 0; i < int(k); i++ {
		val := calculatePointAtDay(cache, days, i)
		sum += int64(val)
	}
	res := sum
	for l, r := 0, int(k); r < int(n)+int(k); r++ {
		valR := calculatePointAtDay(cache, days, r)
		valL := calculatePointAtDay(cache, days, l)
		sum += int64(valR) - int64(valL)
		res = max(res, sum)
		l++
	}
	return res
}

func calculatePointAtDay(cache map[int]int32, days []int32, dayIdx int) int32 {
	res := int32(0)
	orgDayIdx := dayIdx
	if val, ok := (cache)[dayIdx]; ok {
		return val
	}
	for i := 0; dayIdx >= 0; i++ {
		if dayIdx < int(days[i]) {
			res = int32(dayIdx + 1)
		}
		dayIdx -= int(days[i])
		if i == len(days)-1 {
			i = -1
		}
	}
	(cache)[orgDayIdx] = res
	return res
}

func daysOfTournament(days []int32) int64 {
	sum := int64(0)
	for i := range days {
		sum += int64(days[i])
	}
	return sum
}

// Solution 3: Maybe still timeout
func getMaximumPoints3(days []int32, k int32) int64 {
	prefix := prefixSum(days)
	n := prefix[len(prefix)-1]
	sum := int64(0)
	cache := make(map[int]int32)
	for i := 0; i < int(k); i++ {
		val := calculatePointAtDayWithPrefixSum(cache, prefix, i)
		sum += int64(val)
	}
	res := sum
	for l, r := 0, int(k); r < int(n)+int(k); r++ {
		valR := calculatePointAtDayWithPrefixSum(cache, prefix, r)
		valL := calculatePointAtDayWithPrefixSum(cache, prefix, l)
		sum += int64(valR) - int64(valL)
		res = max(res, sum)
		l++
	}
	return res
}

func prefixSum(days []int32) []int64 {
	prefix := make([]int64, len(days))
	sum := int64(0)
	for i := range days {
		sum += int64(days[i])
		prefix[i] = sum
	}
	return prefix
}

func calculatePointAtDayWithPrefixSum(cache map[int]int32, prefixSum []int64, dayIdx int) int32 {
	res := int32(0)
	dayIdx = dayIdx % int(prefixSum[len(prefixSum)-1])
	if val, ok := (cache)[dayIdx]; ok {
		return val
	}
	idx, _ := slices.BinarySearch(prefixSum, int64(dayIdx+1))
	if idx > 0 {
		res = int32(dayIdx) - int32(prefixSum[idx-1]) + 1
	} else {
		res = int32(dayIdx) + 1
	}

	(cache)[dayIdx] = res
	return res
}

func getMaximumPoints4(days []int32, k int32) int64 {
	sum, i, j := int64(0), 0, int32(1)
	for k := k; k > 0; {
		if k > days[i] {
			k -= days[i]
			sum += int64(days[i]) * int64(days[i]+1) / 2
			i++
		} else {
			sum += int64(j)
			k--
			i, j = increase(days, i, j, true)
		}
	}
	res := sum
	extDays := make([]int32, 2*len(days))
	copy(extDays, days)
	copy(extDays[len(days):], days)

	li, lj := 0, int32(1)
	for i < len(extDays) {
		sum += int64(j) - int64(lj)
		res = max(res, sum)
		i, j = increase(extDays, i, j, false)
		li, lj = increase(extDays, li, lj, false)
	}
	return res
}

func increase(days []int32, i int, j int32, reset bool) (int, int32) {
	if j < days[i] {
		j++
	} else {
		i++
		j = 1
	}
	if reset && i == len(days) {
		i = 0
	}
	return i, j
}
