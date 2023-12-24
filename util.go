package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func NumberRuneToInt(r rune) int {
	return int('0' - r)
}

func ReadFileToRunes(fileName string) [][]rune {
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	lines := strings.Fields(string(fileBytes))

	chars := make([][]rune, len(lines))

	for i, line := range lines {
		chars[i] = make([]rune, len(line))
		for j, c := range line {
			chars[i][j] = c
		}
	}

	return chars
}

func ReadFileToLines(fileName string) []string {
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	return strings.Fields(string(fileBytes))
}

func ReadFileWithSpacesToLines(fileName string) []string {
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	return strings.Split(string(fileBytes), "\n")
}

func MustAtoi(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
