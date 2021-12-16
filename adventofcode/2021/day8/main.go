package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	signals, outputs := readInput()
	part1(outputs)
	part2(signals, outputs)
}

func part1(outputs [][]string) {
	count := 0
	for _, output := range outputs {
		for _, o := range output {
			switch len(o) {

			case 2, 3, 4, 7:
				count++
			}
		}
	}
	fmt.Println("Part 1: ", count)
}

type sevenSegment struct {
	keyMap map[string]string
}

func part2(signals, outputs [][]string) {
	result := 0
	for i := range signals {
		val := part2SmallProblem(signals[i], outputs[i])
		result += val
	}
	fmt.Println("Part 2: ", result)
}

func part2SmallProblem(signal []string, output []string) int {
	// ss := make(map[string]string)
	one := findLen(signal, 2)
	four := findLen(signal, 4)
	seven := findLen(signal, 3)
	eight := findLen(signal, 7)
	six := findSix(signal, one)
	c := stringSub(one, six)
	three := findThree(signal, one)
	f := stringSub(one, c)
	two := findLenFiveNotConstains(signal, f)
	five := findLenFiveNotConstains(signal, c)
	// zero := findZero(signal, six, four)
	nine := findNine(signal, six, four)

	// fmt.Println(zero, one, two, three, four, five, six, seven, eight, nine)

	result := 0
	for _, o := range output {
		switch {
		case equal(o, one):
			result++
		case equal(o, two):
			result += 2
		case equal(o, three):
			result += 3
		case equal(o, four):
			result += 4
		case equal(o, five):
			result += 5
		case equal(o, six):
			result += 6
		case equal(o, seven):
			result += 7
		case equal(o, eight):
			result += 8
		case equal(o, nine):
			result += 9
		}
		result *= 10
	}
	return result / 10
}

func findNine(signal []string, six, four string) string {
	for i := 0; i < len(signal); i++ {
		if len(signal[i]) == 6 && !equal(signal[i], six) && contains(signal[i], four) {
			return signal[i]
		}
	}
	return ""
}

func findZero(signal []string, six, four string) string {
	for i := 0; i < len(signal); i++ {
		if len(signal[i]) == 6 && !equal(signal[i], six) && !contains(signal[i], four) {
			return signal[i]
		}
	}
	return ""
}

func findLenFiveNotConstains(signal []string, c string) string {
	for i := 0; i < len(signal); i++ {
		if len(signal[i]) == 5 && !strings.Contains(signal[i], c) {
			return signal[i]
		}
	}
	return ""
}

func findSix(signal []string, one string) string {
	for i := 0; i < len(signal); i++ {
		if len(signal[i]) == 6 && !contains(signal[i], one) {
			return signal[i]
		}
	}
	return ""
}

func findThree(signal []string, one string) string {
	for i := 0; i < len(signal); i++ {
		if len(signal[i]) == 5 && contains(signal[i], one) {
			return signal[i]
		}
	}
	return ""
}

// return the string constained in a but not in b
func stringSub(a, b string) string {
	var result string
	for i := 0; i < len(a); i++ {
		if !strings.Contains(b, a[i:i+1]) {
			result += a[i : i+1]
		}
	}
	return result
}

// return true if all characters in b are in a
func contains(a, b string) bool {
	for i := 0; i < len(b); i++ {
		if !strings.Contains(a, b[i:i+1]) {
			return false
		}
	}
	return true
}

func findLen(signal []string, length int) string {
	for i := 0; i < len(signal); i++ {
		if len(signal[i]) == length {
			return signal[i]
		}
	}
	return ""
}

func equal(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !strings.Contains(b, a[i:i+1]) {
			return false
		}
	}
	return true
}

// read input from stdin, line by line
// the separator is \n
func readInput() ([][]string, [][]string) {
	file, err := os.Open("./my_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var (
		signal = make([][]string, 0)
		output = make([][]string, 0)
	)

	for scanner.Scan() {
		if str := scanner.Text(); str != "" {
			inputs := strings.Split(str, " | ")
			s := strings.Split(inputs[0], " ")
			signal = append(signal, s)
			o := strings.Split(inputs[1], " ")
			output = append(output, o)
		} else {
			break
		}
	}
	return signal, output
}
