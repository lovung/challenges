package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lanternfishs := readInput()
	lanternfishs2 := make([]int, len(lanternfishs))
	copy(lanternfishs2, lanternfishs)
	part1(lanternfishs)
	part2(lanternfishs2)
}

const reset = 6
const newValue = 8

func part1(lanternfishs []int) {
	for i := 0; i < 80; i++ {
		n := len(lanternfishs)
		for j := 0; j < n; j++ {
			lanternfishs[j]--
			if lanternfishs[j] == -1 {
				lanternfishs[j] = reset
				lanternfishs = append(lanternfishs, newValue)
			}
		}
	}
	fmt.Println("Result: ", len(lanternfishs))
}

func part2(lanternfishs []int) {

}

// read input from stdin, line by line
// the separator is \n
func readInput() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lanternfishs []int

	for scanner.Scan() {
		if str := scanner.Text(); str != "" {
			lanternfishStr := strings.Split(str, ",")
			for _, lanternfish := range lanternfishStr {
				if lanternfish != "" {
					time, err := strconv.Atoi(lanternfish)
					if err != nil {
						panic(err)
					}
					lanternfishs = append(lanternfishs, time)
				}
			}
		} else {
			break
		}
	}
	return lanternfishs
}
