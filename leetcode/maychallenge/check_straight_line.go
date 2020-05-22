package maychallenge

func checkStraightLine(coordinates [][]int) bool {
	dy := float64(coordinates[1][1] - coordinates[0][1])
	if dy == 0 {
		for i := 1; i < len(coordinates)-1; i++ {
			if coordinates[i+1][1] != coordinates[i][1] {
				return false
			}
		}
		return true
	}
	k := float64(coordinates[1][0]-coordinates[0][0]) / dy
	for i := 1; i < len(coordinates)-1; i++ {
		if k != float64(coordinates[i+1][0]-coordinates[i][0])/float64(coordinates[i+1][1]-coordinates[i][1]) {
			return false
		}
	}
	return true
}
