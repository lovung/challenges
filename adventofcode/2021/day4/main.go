package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	boards, nums := readInput()
	part1(boards, nums)
	part2(boards, nums)
}

func part1(boards []*board, nums []int) {
	for _, n := range nums {
		for _, b := range boards {
			b.Mark(n)
		}
		// for _, b := range boards {
		// 	fmt.Println(b)
		// }
		// fmt.Println(n)
		for _, b := range boards {
			if b.IsBingo() {
				fmt.Println("bingo")
				fmt.Println("Result: ", b.SumUnmarkNums()*n)
				return
			}
		}
	}
}

func part2(boards []*board, nums []int) {
	for _, n := range nums {
		for _, b := range boards {
			fmt.Println(b)
		}
		fmt.Println(n)
		newBoard := make([]*board, 0)
		for _, b := range boards {
			if !b.IsBingo() {
				newBoard = append(newBoard, b)
				b.Mark(n)
				if b.IsBingo() && len(boards) == 1 {
					fmt.Println("Result: ", b.SumUnmarkNums()*n)
					return
				}
			}
		}
		boards = newBoard
	}
}

// read input from stdin, line by line
// the separator is \n
func readInput() ([]*board, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var boards []*board
	var dawnNums []int
	isReadingDrawNumber := true
	onNewLine := false
	var b *board
	var row int

	for scanner.Scan() {
		if isReadingDrawNumber {
			text := scanner.Text()
			nums := strings.Split(text, ",")
			for _, n := range nums {
				num, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				dawnNums = append(dawnNums, num)
			}
			isReadingDrawNumber = false
			continue
		}
		if str := scanner.Text(); str == "" {
			if onNewLine {
				break
			} else {
				onNewLine = true
			}
			if b != nil {
				boards = append(boards, b)
			}
			fmt.Println("new board")
			b = &board{}
			row = 0
		} else {
			onNewLine = false
			fmt.Println("read: ", str)
			nums := strings.Split(str, " ")
			i := 0
			for _, n := range nums {
				if n == "" {
					continue
				}
				num, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				b.rows[row][i] = num
				i++
			}
			row++
		}
	}
	return boards, dawnNums
}

type board struct {
	rows [5][5]int
	mark [5][5]int
}

func (b *board) Mark(num int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.rows[i][j] == num {
				b.mark[i][j] = 1
			}
		}
	}
}

func (b *board) IsBingo() bool {
	// Sum the rows
	for i := 0; i < 5; i++ {
		sum := 0
		for j := 0; j < 5; j++ {
			sum += b.mark[i][j]
		}
		if sum == 5 {
			return true
		}
	}
	// Sum the columns
	for i := 0; i < 5; i++ {
		sum := 0
		for j := 0; j < 5; j++ {
			sum += b.mark[j][i]
		}
		if sum == 5 {
			return true
		}
	}
	return false
}
func (b *board) SumUnmarkNums() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.mark[i][j] == 0 {
				sum += b.rows[i][j]
			}
		}
	}
	return sum
}
