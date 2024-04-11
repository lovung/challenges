package apr2024

// https://leetcode.com/problems/destination-city/description/
func destCity(paths [][]string) string {
	adjMap := make(map[string]string)
	for i := range paths {
		adjMap[paths[i][0]] = paths[i][1]
	}
	city := paths[0][0]
	for len(city) > 0 {
		if _city, ok := adjMap[city]; ok {
			city = _city
		} else {
			break
		}
	}
	return city
}
