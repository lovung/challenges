package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	part1(input)
	part2(input)
}

func part1(input []string) {
	depth, position := 0, 0
	for i := range input {
		args := strings.Split(input[i], " ")
		if len(args) != 2 {
			fmt.Println("len is not 2", input[i])
			continue
		}
		switch args[0] {
		case "forward":
			val, _ := strconv.Atoi(args[1])
			position += val
		case "up":
			val, _ := strconv.Atoi(args[1])
			depth -= val
		case "down":
			val, _ := strconv.Atoi(args[1])
			depth += val
		}
	}
	fmt.Println("Result: ", depth*position)
}

func part2(input []string) {
	aim, depth, position := 0, 0, 0
	for i := range input {
		args := strings.Split(input[i], " ")
		if len(args) != 2 {
			fmt.Println("len is not 2", input[i])
			continue
		}
		switch args[0] {
		case "forward":
			val, _ := strconv.Atoi(args[1])
			position += val
			depth += val * aim
		case "up":
			val, _ := strconv.Atoi(args[1])
			aim -= val
		case "down":
			val, _ := strconv.Atoi(args[1])
			aim += val
		}
	}
	fmt.Println("Result: ", depth*position)
}

// read input from stdin, line by line
// the separator is \n
func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		if str := scanner.Text(); str == "" {
			break
		} else {
			text = append(text, scanner.Text())
		}
	}
	return text
}
