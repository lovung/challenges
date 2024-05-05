package biweekly129

func cal1row(row []int) int64 {
	res := int64(0)
	prefixR := prefixFromRight(row)
	prefixL := prefixFromLeft(row)
	for i := 0; i < len(row); i++ {
		if row[i] == 0 {
			continue
		}
		res += int64((row[i] - 1) * (prefixR[i] - 1))
		res += int64((row[i] - 1) * (prefixL[i] - 1))
	}
	return res
}

func numberOfRightTriangles(grid [][]int) int64 {
	res := int64(0)
	prefixBot := prefixOneFromBot(grid)
	for i := range prefixBot {
		res += cal1row(prefixBot[i])
	}
	prefixTop := prefixOneFromTop(grid)
	for i := range prefixTop {
		res += cal1row(prefixTop[i])
	}
	return res
}

func prefixFromLeft(row []int) []int {
	prefix := make([]int, len(row))
	cnt := 0
	for i := range row {
		if row[i] == 0 {
			continue
		}
		cnt++
		prefix[i] = cnt
	}
	return prefix
}

func prefixFromRight(row []int) []int {
	prefix := make([]int, len(row))
	cnt := 0
	for i := len(row) - 1; i >= 0; i-- {
		if row[i] == 0 {
			continue
		}
		cnt++
		prefix[i] = cnt
	}
	return prefix
}

func prefixOneFromBot(grid [][]int) [][]int {
	prefix := make([][]int, len(grid))
	for i := range prefix {
		prefix[i] = make([]int, len(grid[i]))
	}
	for j := 0; j < len(grid[0]); j++ {
		cnt := 0
		for i := len(grid) - 1; i >= 0; i-- {
			if grid[i][j] == 0 {
				continue
			}
			cnt++
			prefix[i][j] = cnt
		}
	}
	return prefix
}

func prefixOneFromTop(grid [][]int) [][]int {
	prefix := make([][]int, len(grid))
	for i := range prefix {
		prefix[i] = make([]int, len(grid[i]))
	}
	for j := 0; j < len(grid[0]); j++ {
		cnt := 0
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == 0 {
				continue
			}
			cnt++
			prefix[i][j] = cnt
		}
	}
	return prefix
}
