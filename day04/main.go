package main

import (
	"fmt"
	"slices"
	"strings"

	util "github.com/yattoni/advent-of-code-2023"
)

type card struct {
	id                         int
	winningNumbers             []int
	numbersYouHave             []int
	countWinningNumbersYouHave int
	countCopies                int
}

func (card card) String() string {
	return fmt.Sprintf("{id=%d,winningNumbers=%v,numbersYouHave=%v,countWinningNumbersYouHave=%d,countCopies=%d}",
		card.id, card.winningNumbers, card.numbersYouHave, card.countWinningNumbersYouHave, card.countCopies)
}

func main() {
	lines := util.ReadFileWithSpacesToLines("input")
	cards := make([]card, len(lines))
	for i, line := range lines {
		numbers := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), "|")
		winningNumbersStr := strings.Fields(strings.TrimSpace(numbers[0]))
		numbersYouHaveStr := strings.Fields(strings.TrimSpace(numbers[1]))
		cards[i] = card{i + 1, make([]int, len(winningNumbersStr)), make([]int, len(numbersYouHaveStr)), 0, 0}
		for j, numStr := range winningNumbersStr {
			cards[i].winningNumbers[j] = util.MustAtoi(numStr)
		}
		for j, numStr := range numbersYouHaveStr {
			cards[i].numbersYouHave[j] = util.MustAtoi(numStr)
		}
	}
	fmt.Println(cards)
	sum := 0
	for i, card := range cards {
		numWinningNumbersYouHave := 0
		for _, numberYouHave := range card.numbersYouHave {
			if slices.Contains[[]int](card.winningNumbers, numberYouHave) {
				numWinningNumbersYouHave++
			}
		}
		cards[i].countWinningNumbersYouHave = numWinningNumbersYouHave
		if numWinningNumbersYouHave == 0 {
			continue
		} else if numWinningNumbersYouHave == 1 {
			sum += 1
		} else {
			score := 1
			for j := 1; j < numWinningNumbersYouHave; j++ {
				score *= 2
			}
			sum += score
		}
	}
	fmt.Println(sum) // 27454

	fmt.Println(cards)
	for i := 0; i < len(cards); i++ {
		if cards[i].countWinningNumbersYouHave == 0 {
			continue
		}
		for j := 0; j < cards[i].countCopies+1; j++ { // +1 to include the original
			for k := 0; k < cards[i].countWinningNumbersYouHave; k++ {
				cards[i+k+1].countCopies++
			}
		}
	}
	totalCards := 0
	for _, card := range cards {
		totalCards += 1
		totalCards += card.countCopies
	}
	fmt.Println(totalCards) // 6857330
}
