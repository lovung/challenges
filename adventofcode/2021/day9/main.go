package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputs := readInput()
	// fmt.Println(inputs)
	part1(inputs)
	part2(inputs)
}

const red = "\033[31m"
const white = "\033[0m"

func part1(inputs [][]int) {
	var sum int
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs[i]); j++ {
			if i > 0 && inputs[i][j] >= inputs[i-1][j] {
				fmt.Print(white, inputs[i][j])
				continue
			}
			if i < len(inputs)-1 && inputs[i][j] >= inputs[i+1][j] {
				fmt.Print(white, inputs[i][j])
				continue
			}
			if j > 0 && inputs[i][j] >= inputs[i][j-1] {
				fmt.Print(white, inputs[i][j])
				continue
			}
			if j < len(inputs[i])-1 && inputs[i][j] >= inputs[i][j+1] {
				fmt.Print(white, inputs[i][j])
				continue
			}
			fmt.Print(string(red), inputs[i][j])
			sum += inputs[i][j] + 1
		}
		fmt.Println()
	}
	fmt.Println("Part 1: ", sum)
}

func part2(inputs [][]int) {
	lowPoints := findLowPoints(inputs)
	sizes := make([]int, 0, len(lowPoints))
	for _, lowPoint := range lowPoints {
		cloneMatrix := make([][]int, len(inputs))
		for i := 0; i < len(inputs); i++ {
			cloneMatrix[i] = make([]int, len(inputs[i]))
			copy(cloneMatrix[i], inputs[i])
		}
		size := findBasinSizeFrom(cloneMatrix, lowPoint)
		fmt.Println(lowPoint, ": ", size)
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)
	fmt.Println(sizes)
	fmt.Println("Part 2: ", sizes[len(sizes)-1]*sizes[len(sizes)-2]*sizes[len(sizes)-3])
}

func findLowPoints(inputs [][]int) [][]int {
	var (
		lowPoints = make([][]int, 0)
	)
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs[i]); j++ {
			if i > 0 && inputs[i][j] >= inputs[i-1][j] {
				continue
			}
			if i < len(inputs)-1 && inputs[i][j] >= inputs[i+1][j] {
				continue
			}
			if j > 0 && inputs[i][j] >= inputs[i][j-1] {
				continue
			}
			if j < len(inputs[i])-1 && inputs[i][j] >= inputs[i][j+1] {
				continue
			}
			lowPoints = append(lowPoints, []int{i, j})
		}
	}
	return lowPoints
}

func findBasinSizeFrom(matrix [][]int, lowPoint []int) int {
	var (
		i = lowPoint[0]
		j = lowPoint[1]
	)
	savedValue := matrix[i][j]
	// fmt.Println("Checking: ", i, j, savedValue)
	size := 1
	matrix[i][j] = -1
	if savedValue == 8 {
		return 1
	}
	if i > 0 && savedValue == matrix[i-1][j]-1 {
		size += findBasinSizeFrom(matrix, []int{i - 1, j})
	}
	if i < len(matrix)-1 && savedValue == matrix[i+1][j]-1 {
		size += findBasinSizeFrom(matrix, []int{i + 1, j})
	}
	if j > 0 && savedValue == matrix[i][j-1]-1 {
		size += findBasinSizeFrom(matrix, []int{i, j - 1})
	}
	if j < len(matrix[i])-1 && savedValue == matrix[i][j+1]-1 {
		size += findBasinSizeFrom(matrix, []int{i, j + 1})
	}
	return size
}

// read input from stdin, line by line
// the separator is \n
func readInput() [][]int {
	file, err := os.Open("./my_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var (
		inputs = make([][]int, 0)
	)

	for scanner.Scan() {
		if str := scanner.Text(); str != "" {
			chars := strings.Split(str, "")
			oneLine := make([]int, 0)
			for _, char := range chars {
				num, err := strconv.Atoi(char)
				if err != nil {
					panic(err)
				}
				oneLine = append(oneLine, num)
			}
			inputs = append(inputs, oneLine)
		} else {
			break
		}
	}
	return inputs
}
