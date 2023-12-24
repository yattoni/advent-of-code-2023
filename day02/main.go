package main

import (
	"fmt"
	"strings"

	util "github.com/yattoni/advent-of-code-2023"
)

type roll struct {
	red   int
	blue  int
	green int
}

func (roll roll) String() string {
	return fmt.Sprintf("{red=%d,blue=%d,green=%d}", roll.red, roll.blue, roll.green)
}

type game struct {
	id    int
	rolls []roll
}

func (game game) String() string {
	return fmt.Sprintf("{id=%d,rolls=%v}", game.id, game.rolls)
}

func main() {
	lines := util.ReadFileWithSpacesToLines("input")
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.ReplaceAll(lines[i], ":", ";")
		lines[i] = strings.ReplaceAll(lines[i], ",", "")
	}
	games := make([]game, len(lines))
	for i, line := range lines {
		games[i] = game{i + 1, []roll{}}
		rolls := strings.Split(line, ";")[1:] // get rid of Game id;
		for j, r := range rolls {
			games[i].rolls = append(games[i].rolls, roll{})
			rollParts := strings.Split(r, " ")[1:] // get rid leading space
			for k := 0; k < len(rollParts)-1; k += 2 {
				count := util.MustAtoi(rollParts[k])
				if rollParts[k+1] == "blue" {
					games[i].rolls[j].blue = count
				}
				if rollParts[k+1] == "red" {
					games[i].rolls[j].red = count
				}
				if rollParts[k+1] == "green" {
					games[i].rolls[j].green = count
				}
			}
		}
	}
	fmt.Println(games)
	// 12 red cubes, 13 green cubes, and 14 blue cubes.
	sum := 0
	for _, game := range games {
		counts := true
		for _, roll := range game.rolls {
			if roll.red > 12 || roll.green > 13 || roll.blue > 14 {
				counts = false
				break
			}
		}
		if counts {
			sum += game.id
		}
	}
	fmt.Println(sum) // 2406

	power := 0
	for _, game := range games {
		minRed := 0
		minBlue := 0
		minGreen := 0
		for _, roll := range game.rolls {
			if roll.red > minRed {
				minRed = roll.red
			}
			if roll.blue > minBlue {
				minBlue = roll.blue
			}
			if roll.green > minGreen {
				minGreen = roll.green
			}
		}
		power += minRed * minBlue * minGreen
	}
	fmt.Println(power) // 78375
}
