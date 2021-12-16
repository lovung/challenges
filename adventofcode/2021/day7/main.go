package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Println("Part 1")
	crabPos := readInput()
	sort.Ints(crabPos)
	meMedian := crabPos[len(crabPos)/2]
	var meSum int
	for _, time := range crabPos {
		if time > meMedian {
			meSum += time - meMedian
		} else {
			meSum += meMedian - time
		}
	}
	fmt.Println("Result: ", meSum)
}

func part2() {
	fmt.Println("Part 2")
	crabPos := readInput()
	// sort.Ints(crabPos)
	// meMedian := crabPos[len(crabPos)/2]
	minMeSum := int(math.MaxInt64)
	for i := 0; i < len(crabPos); i++ {
		meSum := 0
		for _, time := range crabPos {
			if time > i {
				meSum += (time - i) * (time - i + 1) / 2
			} else {
				meSum += (i - time) * (i - time + 1) / 2
			}
		}
		// fmt.Println("mesum", meSum)
		if meSum < minMeSum {
			minMeSum = meSum
		}
	}
	fmt.Println("Result: ", minMeSum)
}

// read input from stdin, line by line
// the separator is \n
func readInput() []int {
	file, err := os.Open("./my_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var crabPos []int

	for scanner.Scan() {
		if str := scanner.Text(); str != "" {
			crabStr := strings.Split(str, ",")
			for _, crab := range crabStr {
				if crab != "" {
					time, err := strconv.Atoi(crab)
					if err != nil {
						panic(err)
					}
					crabPos = append(crabPos, time)
				}
			}
		} else {
			break
		}
	}
	return crabPos
}
