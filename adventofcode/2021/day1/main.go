package main

import "fmt"

func main() {
	input := readInput()
	part1(input)
	part2(input)
}

func part1(input []int) {
	increaseCnt := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increaseCnt++
		}
	}
	fmt.Println("Result: ", increaseCnt)
}

func part2(input []int) {
	sum := input[0] + input[1] + input[2]
	increaseCnt := 0
	for i := 1; i < len(input)-2; i++ {
		if input[i]+input[i+1]+input[i+2] > sum {
			increaseCnt++
		}
		sum = input[i] + input[i+1] + input[i+2]
	}
	fmt.Println("Result: ", increaseCnt)
}

// read input from stdin, line by line
// and convert to int
func readInput() []int {
	var input []int
	for {
		var i int
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			break
		}
		input = append(input, i)
	}
	return input
}
