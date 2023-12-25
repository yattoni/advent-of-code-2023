package main

import (
	"fmt"
	"strings"

	util "github.com/yattoni/advent-of-code-2023"
)

func findNextNumber(sequence []int) int {
	allDifferences := getAllDifferences(sequence)
	fmt.Println("allDifferences", allDifferences)
	next := sequence[len(sequence)-1]
	for i := 0; i < len(allDifferences); i++ {
		next += allDifferences[i][len(allDifferences[i])-1]
	}
	return next
}

func findPreviousNumber(sequence []int) int {
	allDifferences := getAllDifferences(sequence)
	fmt.Println("allDifferences", allDifferences)
	prev := 0
	for i := len(allDifferences) - 1; i >= 0; i-- {
		prev = allDifferences[i][0] - prev
	}
	return sequence[0] - prev
}

func getAllDifferences(sequence []int) [][]int {
	allDifferences := make([][]int, 1)
	differences, allZeros := getDifferences(sequence)
	fmt.Println("sequence", sequence, "differences", differences, "allZeros", allZeros)
	allDifferences[0] = differences
	if !allZeros {
		return append(allDifferences, getAllDifferences(differences)...)
	} else {
		return allDifferences
	}
}

func getDifferences(sequence []int) ([]int, bool) {
	differences := make([]int, len(sequence)-1)
	allZeros := true
	for i := 0; i < len(sequence)-1; i++ {
		differences[i] = sequence[i+1] - sequence[i]
		allZeros = allZeros && (differences[i] == 0)
	}
	return differences, allZeros
}

func main() {
	lines := util.ReadFileWithSpacesToLines(util.INPUT)

	sequences := make([][]int, len(lines))
	for i, line := range lines {
		numStrs := strings.Fields(line)
		sequences[i] = make([]int, len(numStrs))
		for j, str := range numStrs {
			sequences[i][j] = util.MustAtoi(str)
		}
	}

	fmt.Println(sequences)

	sum := 0
	for _, sequence := range sequences {
		next := findNextNumber(sequence)
		fmt.Println("sequence", sequence, "next", next, "\n")
		sum += next
	}
	fmt.Println("sum", sum) // 1861775706

	sum2 := 0
	for _, sequence := range sequences {
		prev := findPreviousNumber(sequence)
		fmt.Println("sequence", sequence, "prev", prev, "\n")
		sum2 += prev
	}
	fmt.Println("sum2", sum2) // 1082
}
