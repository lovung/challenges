package maychallenge

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	h := len(image)
	w := len(image[0])
	curColor := image[sr][sc]
	floodFillPointer(image, sr, sc, &newColor, &h, &w, &curColor)
	return image
}

func floodFillPointer(image [][]int, sr int, sc int, newColor *int, h, w, curColor *int) {
	if sc >= *w || sc < 0 || sr >= *h || sr < 0 {
		return
	}

	if image[sr][sc] == *newColor || image[sr][sc] != *curColor {
		return
	}

	image[sr][sc] = *newColor

	floodFillPointer(image, sr, sc+1, newColor, h, w, curColor)
	floodFillPointer(image, sr, sc-1, newColor, h, w, curColor)
	floodFillPointer(image, sr+1, sc, newColor, h, w, curColor)
	floodFillPointer(image, sr-1, sc, newColor, h, w, curColor)
}
