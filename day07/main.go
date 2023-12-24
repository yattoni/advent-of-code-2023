package main

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strings"

	util "github.com/yattoni/advent-of-code-2023"
)

type game struct {
	hand          []byte
	bid           int
	handType      HandType
	jokerHandType HandType
}

func (game game) String() string {
	return fmt.Sprintf("game{hand=\"%s\",bid=%d,type=%s,jokerType=%s}", game.hand, game.bid, game.handType, game.jokerHandType)
}

type HandType int

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func (t HandType) String() string {
	if t == -1 {
		return "Unknown"
	}
	return [...]string{
		"FiveOfAKind",
		"FourOfAKind",
		"FullHouse",
		"ThreeOfAKind",
		"TwoPair",
		"OnePair",
		"HighCard",
	}[t]
}

func getHandType(hand []byte) HandType {
	countsByCard := make(map[byte]int)
	for _, card := range hand {
		countsByCard[card]++
	}
	if len(countsByCard) == 1 { // all the same
		return FiveOfAKind
	} else if len(countsByCard) == 5 { // all different
		return HighCard
	} else if len(countsByCard) == 4 { // one card appears twice
		return OnePair
	} else if len(countsByCard) == 2 {
		// 2 different cards and first card appears 4 or 1 time
		if countsByCard[hand[0]] == 4 || countsByCard[hand[0]] == 1 {
			return FourOfAKind
		} else { // 3/2 split instad of 4/1
			return FullHouse
		}
	} else if len(countsByCard) == 3 {
		for _, count := range countsByCard {
			// 3 different cards and one card appears 3 times
			if count == 3 {
				return ThreeOfAKind
			}
		}
		// 3 different cards and no card appears 3 times
		return TwoPair
	} else {
		return -1
	}
}

func getCardStrength(card byte) int {
	cards := [...]byte{
		'A',
		'K',
		'Q',
		'J',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
	}
	for i, c := range cards {
		if c == card {
			return i
		}
	}
	log.Fatalf("Unknown card, %s", string(card))
	return -1
}

func getCardStrength2(card byte) int {
	cards := [...]byte{
		'A',
		'K',
		'Q',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
		'J',
	}
	for i, c := range cards {
		if c == card {
			return i
		}
	}
	log.Fatalf("Unknown card, %s", string(card))
	return -1
}

func compareGame(g1, g2 game) int {
	if g1.handType != g2.handType {
		if g1.handType < g2.handType {
			return -1
		} else {
			return 1
		}
	}
	for i := 0; i < len(g1.hand); i++ {
		if g1.hand[i] != g2.hand[i] {
			if getCardStrength(g1.hand[i]) < getCardStrength(g2.hand[i]) {
				return -1
			} else {
				return 1
			}
		}
	}
	log.Fatalf("Games are equal, g1=%s, g2=%s", g1, g2)
	return 0
}

func compareJokerGame(g1, g2 game) int {
	if g1.jokerHandType != g2.jokerHandType {
		if g1.jokerHandType < g2.jokerHandType {
			return -1
		} else {
			return 1
		}
	}
	for i := 0; i < len(g1.hand); i++ {
		if g1.hand[i] != g2.hand[i] {
			if getCardStrength2(g1.hand[i]) < getCardStrength2(g2.hand[i]) {
				return -1
			} else {
				return 1
			}
		}
	}
	log.Fatalf("Games are equal, g1=%s, g2=%s", g1, g2)
	return 0
}

func getJokerHandType(hand []byte) HandType {
	// convert joker J card into most common card
	countsByCard := make(map[byte]int)
	for _, card := range hand {
		countsByCard[card]++
	}
	if len(countsByCard) == 1 {
		// all cards are the same, nothing todo
		return getHandType(hand)
	}
	if countsByCard['J'] == 0 {
		// no J's to replace
		return getHandType(hand)
	}

	uniqueCards := make([]byte, len(countsByCard))
	for card := range countsByCard {
		uniqueCards = append(uniqueCards, card)
	}
	sort.SliceStable(uniqueCards, func(i, j int) bool {
		return countsByCard[uniqueCards[i]] > countsByCard[uniqueCards[j]]
	})
	for _, mostCommonCard := range uniqueCards {
		if mostCommonCard != 'J' {
			newHand := make([]byte, len(hand))
			copy(newHand, hand)
			for i := 0; i < len(newHand); i++ {
				if newHand[i] == 'J' {
					newHand[i] = mostCommonCard
				}
			}
			return getHandType(newHand)
		}
	}
	log.Fatalf("Expected some J's to replace in hand, %s", string(hand))
	return getHandType(hand)
}

func main() {
	lines := util.ReadFileWithSpacesToLines("input")
	games := make([]game, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		hand := []byte(parts[0])
		bid := util.MustAtoi(parts[1])
		games[i] = game{hand, bid, getHandType(hand), getJokerHandType(hand)}
	}

	fmt.Println("Input:")

	for _, game := range games {
		fmt.Println(game)
	}

	slices.SortFunc[[]game](games, compareGame)

	winnings := 0
	fmt.Println("\nSorted:")
	for i, game := range games {
		fmt.Println(game)
		winnings += game.bid * (len(games) - i)
	}

	fmt.Println("winnings", winnings) // 248836197

	slices.SortFunc[[]game](games, compareJokerGame)

	winnings2 := 0
	fmt.Println("\nSorted Joker:")
	for i, game := range games {
		fmt.Println(game)
		winnings2 += game.bid * (len(games) - i)
	}

	fmt.Println("winnings 2", winnings2) // 251195607
}
