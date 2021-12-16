package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInput()
	// fmt.Println(lines)
	part1(lines)
	part2(lines)
}

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func (l *Line) IsHorizotalConst() bool {
	return l.start.y == l.end.y
}

func (l *Line) IsVerticalConst() bool {
	return l.start.x == l.end.x
}

// read input from stdin, line by line
// the separator is \n
func readInput() []Line {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines []Line

	for scanner.Scan() {
		if str := scanner.Text(); str == "" {
			break
		} else {
			points := strings.Split(str, " -> ")
			line := Line{}
			line.start = parsePoint(points[0])
			line.end = parsePoint(points[1])
			lines = append(lines, line)
		}
	}
	return lines
}

func parsePoint(str string) Point {
	var point Point
	strs := strings.Split(str, ",")
	point.x, _ = strconv.Atoi(strs[0])
	point.y, _ = strconv.Atoi(strs[1])
	return point
}

func part1(lines []Line) {
	var matrix [1000][1000]int
	fmt.Println("Part 1")
	for _, line := range lines {
		// fmt.Println(line)
		if line.IsHorizotalConst() {
			// fmt.Println("IsHorizotal")
			if line.start.x < line.end.x {
				for i := line.start.x; i <= line.end.x; i++ {
					matrix[i][line.start.y]++
					// fmt.Println("Mark: ", i, line.start.y)
				}
			} else {
				for i := line.start.x; i >= line.end.x; i-- {
					matrix[i][line.start.y]++
					// fmt.Println("Mark: ", i, line.start.y)
				}
			}
		} else if line.IsVerticalConst() {
			// fmt.Println("IsVertical")
			if line.start.y < line.end.y {
				for i := line.start.y; i <= line.end.y; i++ {
					matrix[line.start.x][i]++
					// fmt.Println("Mark: ", line.start.x, i)
				}
			} else {
				for i := line.start.y; i >= line.end.y; i-- {
					matrix[line.start.x][i]++
					// fmt.Println("Mark: ", line.start.x, i)
				}
			}
		}
	}

	var count int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if matrix[i][j] > 1 {
				// fmt.Println("Overlapping: ", i, j)
				count++
			}
		}
	}
	fmt.Println("Result: ", count)
}

func part2(lines []Line) {
	var matrix [1000][1000]int
	fmt.Println("Part 2")
	for _, line := range lines {
		// fmt.Println(line)
		if line.IsHorizotalConst() {
			// fmt.Println("IsHorizotal")
			if line.start.x < line.end.x {
				for i := line.start.x; i <= line.end.x; i++ {
					matrix[i][line.start.y]++
					// fmt.Println("Mark: ", i, line.start.y)
				}
			} else {
				for i := line.start.x; i >= line.end.x; i-- {
					matrix[i][line.start.y]++
					// fmt.Println("Mark: ", i, line.start.y)
				}
			}
		} else if line.IsVerticalConst() {
			// fmt.Println("IsVertical")
			if line.start.y < line.end.y {
				for i := line.start.y; i <= line.end.y; i++ {
					matrix[line.start.x][i]++
					// fmt.Println("Mark: ", line.start.x, i)
				}
			} else {
				for i := line.start.y; i >= line.end.y; i-- {
					matrix[line.start.x][i]++
					// fmt.Println("Mark: ", line.start.x, i)
				}
			}
		} else {
			// fmt.Println("IsDiagonal")
			if line.start.x < line.end.x && line.start.y < line.end.y {
				for i := 0; i <= line.end.x-line.start.x; i++ {
					matrix[line.start.x+i][line.start.y+i]++
					// fmt.Println("Mark: ", line.start.x+i, line.start.y+i)
				}
			} else if line.start.x > line.end.x && line.start.y > line.end.y {
				for i := 0; i <= line.start.x-line.end.x; i++ {
					matrix[line.start.x-i][line.start.y-i]++
					// fmt.Println("Mark: ", line.start.x-i, line.start.y-i)
				}
			} else if line.start.x < line.end.x && line.start.y > line.end.y {
				for i := 0; i <= line.start.y-line.end.y; i++ {
					matrix[line.start.x+i][line.start.y-i]++
					// fmt.Println("Mark: ", line.start.x+i, line.start.y-i)
				}
			} else if line.start.x > line.end.x && line.start.y < line.end.y {
				for i := 0; i <= line.end.y-line.start.y; i++ {
					matrix[line.start.x-i][line.start.y+i]++
					// fmt.Println("Mark: ", line.start.x-i, line.start.y+i)
				}
			}
		}
	}

	var count int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if matrix[i][j] > 1 {
				// fmt.Println("Overlapping: ", i, j)
				count++
			}
		}
	}
	fmt.Println("Result: ", count)
}
