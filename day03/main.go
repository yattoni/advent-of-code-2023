package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readFileToRunes(fileName string) [][]rune {
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

type part struct {
	number int
	x      int
	y      int
	z      int // how many characters the number is
}

func (p part) String() string {
	return fmt.Sprintf("{number=%d,x=%d,y=%d,z=%d}", p.number, p.x, p.y, p.z)
}

type coord struct {
	x int
	y int
}

func (c coord) String() string {
	return fmt.Sprintf("{x=%d,y=%d}", c.x, c.y)
}

func buildParts(input [][]rune) []part {
	parts := []part{}

	for y, line := range input {
		numParts := []byte{}
		for x, c := range line {
			if unicode.IsNumber(c) {
				numParts = append(numParts, byte(c))
			} else {
				if len(numParts) > 0 {
					// num, _ := strconv.Atoi(string(byte(c)))
					// int(c - '0')
					num, _ := strconv.Atoi(string(numParts))
					parts = append(parts, part{num, x - len(numParts), y, len(numParts)})
					numParts = []byte{}
				}
			}
		}
		if len(numParts) > 0 {
			num, _ := strconv.Atoi(string(numParts))
			parts = append(parts, part{num, len(line) - len(numParts), y, len(numParts)})
		}
	}
	return parts
}

func findGears(input [][]rune) []coord {
	coords := []coord{}
	for y, line := range input {
		for x, c := range line {
			if byte(c) == '*' {
				coords = append(coords, coord{x, y})
			}
		}
	}
	return coords
}

func isSymbol(c rune) bool {
	return !unicode.IsNumber(c) && byte(c) != '.'
}

func isSurroundedBySymbol(p part, input [][]rune) bool {
	for i := 0; i < p.z; i++ {
		// down
		if p.y+1 != len(input) && isSymbol(input[p.y+1][p.x+i]) {
			return true
		}
		// up
		if p.y != 0 && isSymbol(input[p.y-1][p.x+i]) {
			return true
		}
		// left
		if p.x != 0 && isSymbol(input[p.y][p.x+i-1]) {
			return true
		}
		// right
		if p.x+i+1 != len(input[p.y]) && isSymbol(input[p.y][p.x+i+1]) {
			return true
		}
		// down and left
		if p.x != 0 && p.y+1 != len(input) && isSymbol(input[p.y+1][p.x+i-1]) {
			return true
		}
		// up and left
		if p.x != 0 && p.y != 0 && isSymbol(input[p.y-1][p.x+i-1]) {
			return true
		}
		// down and right
		if p.x+i+1 != len(input[p.y]) && p.y+1 != len(input) && isSymbol(input[p.y+1][p.x+i+1]) {
			return true
		}
		// up and right
		if p.x+i+1 != len(input[p.y]) && p.y != 0 && isSymbol(input[p.y-1][p.x+i+1]) {
			return true
		}
	}
	return false
}

func getGearRatio(gear coord, parts []part) int {
	adjacentParts := []part{}
	for _, p := range parts {
		for i := 0; i < p.z; i++ {
			// left
			if p.x+i-1 == gear.x && p.y == gear.y {
				adjacentParts = append(adjacentParts, p)
				break
			}
			// right
			if p.x+i+1 == gear.x && p.y == gear.y {
				adjacentParts = append(adjacentParts, p)
				break
			}
			// up
			if p.x+i == gear.x && p.y-1 == gear.y {
				adjacentParts = append(adjacentParts, p)
				break
			}
			// down
			if p.x+i == gear.x && p.y+1 == gear.y {
				adjacentParts = append(adjacentParts, p)
				break
			}
			// up left
			if p.x+i-1 == gear.x && p.y-1 == gear.y {
				adjacentParts = append(adjacentParts, p)
				break
			}
			// down left
			if p.x+i-1 == gear.x && p.y+1 == gear.y {
				adjacentParts = append(adjacentParts, p)
				break
			}
			// up right
			if p.x+i+1 == gear.x && p.y-1 == gear.y {
				adjacentParts = append(adjacentParts, p)
				break
			}
			// down right
			if p.x+i+1 == gear.x && p.y+1 == gear.y {
				adjacentParts = append(adjacentParts, p)
				break
			}
		}
	}
	if len(adjacentParts) == 2 {
		fmt.Println("gear:", gear)
		return adjacentParts[0].number * adjacentParts[1].number
	}
	return 0
}

func main() {
	input := readFileToRunes("input")

	parts := buildParts(input)
	fmt.Println(parts)

	sum := 0
	for _, p := range parts {
		if isSurroundedBySymbol(p, input) {
			sum += p.number
		}
	}
	fmt.Println(sum) // 521515

	gears := findGears(input)
	gearRatios := 0
	for _, g := range gears {
		gearRatios += getGearRatio(g, parts)
	}
	fmt.Println(gearRatios) // 69527306
}
