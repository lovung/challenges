package main

import "fmt"

func main() {
	var length int
	fmt.Scanf("%d", &length)
	input := readInput()
	// part 1
	part1(input, length)
	part2(input, length)
}

func part1(input []int, length int) {
	counter := make([]int, length)
	for i := range input {
		for j := 0; j < length; j++ {
			if input[i]&(1<<uint(j)) != 0 {
				counter[j]++
			} else {
				counter[j]--
			}
		}
	}
	var gamma, epsilon int
	for i := range counter {
		if counter[i] > 0 {
			gamma += 1 << uint(i)
		} else if counter[i] < 0 {
			epsilon += 1 << uint(i)
		}
	}
	fmt.Println("Part 1")
	fmt.Println(gamma, epsilon)
	fmt.Println("Result: ", gamma*epsilon)
}

func part2(input []int, length int) {
	fmt.Println("Part 2")
	oxy := part2Oxy(input, length)
	co2 := part2CO2(input, length)
	fmt.Println("Oxygen: ", oxy)
	fmt.Println("CO2: ", co2)
	fmt.Println("Result: ", oxy*co2)
}

func part2Oxy(input []int, length int) int {
	for j := length - 1; j >= 0; j-- {
		counter := 0
		if len(input) == 1 {
			return input[0]
		}
		for i := range input {
			if input[i]&(1<<uint(j)) != 0 {
				counter++
			} else {
				counter--
			}
		}
		newList := make([]int, 0)
		for i := range input {
			if input[i]&(1<<uint(j)) != 0 && counter >= 0 {
				newList = append(newList, input[i])
			} else if input[i]&(1<<uint(j)) == 0 && counter < 0 {
				newList = append(newList, input[i])
			}
		}
		input = newList
		if len(input) == 1 {
			return input[0]
		}
	}
	return 0
}

func part2CO2(input []int, length int) int {
	for j := length - 1; j >= 0; j-- {
		counter := 0
		for i := range input {
			if input[i]&(1<<uint(j)) != 0 {
				counter++
			} else {
				counter--
			}
		}
		newList := make([]int, 0)
		for i := range input {
			if input[i]&(1<<uint(j)) != 0 && counter < 0 {
				newList = append(newList, input[i])
			} else if input[i]&(1<<uint(j)) == 0 && counter >= 0 {
				newList = append(newList, input[i])
			}
		}
		input = newList
		if len(input) == 1 {
			return input[0]
		}
	}
	return 0
}

// read input from stdin, line by line
// and convert to int
func readInput() []int {
	var input []int
	for {
		var i int
		_, err := fmt.Scanf("%b", &i)
		if err != nil {
			break
		}
		input = append(input, i)
	}
	return input
}
