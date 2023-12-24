package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	util "github.com/yattoni/advent-of-code-2023"
)

type numberSymbol struct {
	str string
	num int
}

func main() {
	chars := util.ReadFileToRunes("input")
	allNums := [][]rune{}
	for i, line := range chars {
		allNums = append(allNums, []rune{})
		for _, c := range line {
			if unicode.IsNumber(c) {
				allNums[i] = append(allNums[i], c)
			}
		}
	}
	values := []string{}
	for _, line := range allNums {
		values = append(values, fmt.Sprintf("%c%c", byte(line[0]), byte(line[len(line)-1])))
	}
	sum := 0
	for _, val := range values {
		num, _ := strconv.Atoi(val)
		sum += num
	}
	fmt.Println(sum) // 55029

	lines := util.ReadFileToLines("input")

	numberSymbols := []numberSymbol{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"0", 0},
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
		{"zero", 0},
	}

	numbers := make([][]int, len(lines))
	for i, line := range lines {
		numbers[i] = make([]int, 2)
		firstIdx := len(line) + 1
		lastIdx := -1
		for _, symbol := range numberSymbols {
			firstTest := strings.Index(line, symbol.str)
			// -1 if not found
			if firstTest != -1 && firstTest < firstIdx {
				firstIdx = firstTest
				numbers[i][0] = symbol.num
			}
			lastTest := strings.LastIndex(line, symbol.str)
			if lastTest > lastIdx {
				lastIdx = lastTest
				numbers[i][1] = symbol.num
			}
		}
	}
	sum2 := 0
	for i := 0; i < len(numbers); i++ {
		// fmt.Println(numbers[i])
		sum2 += numbers[i][0]*10 + numbers[i][1]
	}
	fmt.Println(sum2) // 55686
}
