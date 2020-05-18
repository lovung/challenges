package maychallenge

func firstUniqChar(s string) int {
	// Solution 1
	// mc := make(map[rune]int)
	// for i, e := range s {
	//     if mc[e] > 0 {
	//         mc[e] = -1
	//     } else if mc[e] == 0 {
	//         mc[e] = i+1
	//     }
	// }
	// var output int
	// for _, v := range mc {
	//     if v <= 0 {
	//         continue
	//     }
	//     if output == 0 || v < output {
	//         output = v
	//     }
	// }
	// return output-1
	var mc [52]int
	for i, e := range s {
		if mc[e-'a'] > 0 {
			mc[e-'a'] = -1
		} else if mc[e-'a'] == 0 {
			mc[e-'a'] = i + 1
		}
	}
	var output int
	for i := 0; i < 52; i++ {
		if mc[i] <= 0 {
			continue
		}
		if output == 0 || mc[i] < output {
			output = mc[i]
		}
	}
	return output - 1
}
