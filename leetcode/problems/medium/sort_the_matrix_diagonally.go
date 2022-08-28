package medium

import "sort"

// Link: https://leetcode.com/problems/sort-the-matrix-diagonally/
func diagonalSort(mat [][]int) [][]int {
	if len(mat) <= 1 || len(mat[0]) <= 1 {
		return mat
	}
	maxI, maxJ := len(mat), len(mat[0])
	for j := maxJ; j >= 1; j-- {
		slice := make([]int, 0)
		tempJ := j
		for i := 0; i < maxJ-j && tempJ < maxJ && i < maxI; i++ {
			slice = append(slice, mat[i][tempJ])
			tempJ++
		}
		sort.Ints(slice)
		tempJ = j
		for i := 0; i < maxJ-j && tempJ < maxJ && i < maxI; i++ {
			mat[i][tempJ] = slice[i]
			tempJ++
		}
	}
	for i := 0; i < maxI; i++ {
		slice := make([]int, 0)
		tempI := i
		for j := 0; j < maxI-i && tempI < maxI && j < maxJ; j++ {
			slice = append(slice, mat[tempI][j])
			tempI++
		}
		sort.Ints(slice)
		tempI = i
		for j := 0; j < maxI-i && tempI < maxI && j < maxJ; j++ {
			mat[tempI][j] = slice[j]
			tempI++
		}
	}
	return mat
}
