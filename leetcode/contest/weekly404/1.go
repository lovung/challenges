package weekly404

func maxHeightOfTriangle(red int, blue int) int {
	return max(
		maxHeightOfTriangle2(red, blue),
		maxHeightOfTriangle2(blue, red),
	)
}

func maxHeightOfTriangle2(red int, blue int) int {
	cnt := 0
	for {
		if cnt%2 == 0 {
			red -= cnt + 1
			if red < 0 {
				break
			}
		} else {
			blue -= cnt + 1
			if blue < 0 {
				break
			}
		}
		cnt++
	}
	return cnt
}
