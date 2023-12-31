package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

const (
	HAND_PAIR = iota
	HAND_DOUBLE_PAIR
	HAND_THREE_OF_A_KIND
	HAND_FULL_HOUSE
	HAND_FOUR_OF_A_KIND
	HAND_FIVE_OF_A_KIND
)

var cardMap = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var jokerCardMap = map[byte]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type hand struct {
	Cards          string
	Bid            int
	HandScore      int
	JokerHandScore int
}

func newHand(cards string, bid int) hand {
	// Count cards of each type
	cardCounts := map[byte]int{}
	jokerCardCounts := map[byte]int{}
	for _, c := range cards {
		if _, ok := cardCounts[byte(c)]; !ok {
			cardCounts[byte(c)] = 1
			jokerCardCounts[byte(c)] = 1
		} else {
			cardCounts[byte(c)]++
			jokerCardCounts[byte(c)]++
		}
	}

	// Add joker count to the highest current card count
	jokerCount := cardCounts['J']
	highest := 0
	var highestCardType byte = '0'
	for cardType, count := range cardCounts {
		if cardType != 'J' && count > highest {
			highest = count
			highestCardType = cardType
		}
	}
	if highest > 0 {
		jokerCardCounts[highestCardType] += jokerCount
	} else {
		jokerCardCounts['2'] = jokerCount
	}

	// Now find the best combinations
	handScore := -1
	jokerHandScore := -1
	for cardType := range cardMap {
		switch cardCounts[cardType] {
		case 2:
			switch handScore {
			case HAND_PAIR:
				handScore = HAND_DOUBLE_PAIR
			case HAND_THREE_OF_A_KIND:
				handScore = HAND_FULL_HOUSE
			case -1:
				handScore = HAND_PAIR
			}
		case 3:
			switch handScore {
			case HAND_PAIR:
				handScore = HAND_FULL_HOUSE
			case -1:
				handScore = HAND_THREE_OF_A_KIND
			}
		case 4:
			if handScore != HAND_FIVE_OF_A_KIND {
				handScore = HAND_FOUR_OF_A_KIND
			}
		case 5:
			handScore = HAND_FIVE_OF_A_KIND
		}

		if cardType != 'J' {
			switch jokerCardCounts[cardType] {
			case 2:
				switch jokerHandScore {
				case HAND_PAIR:
					jokerHandScore = HAND_DOUBLE_PAIR
				case HAND_THREE_OF_A_KIND:
					jokerHandScore = HAND_FULL_HOUSE
				case -1:
					jokerHandScore = HAND_PAIR
				}
			case 3:
				switch jokerHandScore {
				case HAND_PAIR:
					jokerHandScore = HAND_FULL_HOUSE
				case -1:
					jokerHandScore = HAND_THREE_OF_A_KIND
				}
			case 4:
				if jokerHandScore != HAND_FIVE_OF_A_KIND {
					jokerHandScore = HAND_FOUR_OF_A_KIND
				}
			case 5:
				jokerHandScore = HAND_FIVE_OF_A_KIND
			}
		}
	}

	return hand{
		Cards:          cards,
		Bid:            bid,
		HandScore:      handScore,
		JokerHandScore: jokerHandScore,
	}
}

func main() {
	file := util.FileFromArgs()

	b := make([]byte, 1)
	hands := []hand{}
	line := ""

	for {
		_, err := file.Read(b)
		if err == io.EOF || b[0] == '\n' {
			parts := strings.SplitN(line, " ", 2)
			bid, _ := strconv.Atoi(parts[1])

			hands = append(hands, newHand(parts[0], bid))

			if err == io.EOF {
				break
			}

			line = ""
			continue
		}

		line += string(b[0])
	}

	slices.SortFunc(hands, func(a, b hand) int {
		if a.HandScore != b.HandScore {
			return a.HandScore - b.HandScore
		}

		for i := 0; i < len(a.Cards); i++ {
			if a.Cards[i] != b.Cards[i] {
				return cardMap[a.Cards[i]] - cardMap[b.Cards[i]]
			}
		}

		return 0
	})

	score := 0
	for i, h := range hands {
		score += h.Bid * (i + 1)
	}

	fmt.Println(score)

	// Part 2 calculate score with jokers

	slices.SortFunc(hands, func(a, b hand) int {
		if a.JokerHandScore != b.JokerHandScore {
			return a.JokerHandScore - b.JokerHandScore
		}

		for i := 0; i < len(a.Cards); i++ {
			if a.Cards[i] != b.Cards[i] {
				return jokerCardMap[a.Cards[i]] - jokerCardMap[b.Cards[i]]
			}
		}

		return 0
	})

	fmt.Println(hands)

	score = 0
	for i, h := range hands {
		score += h.Bid * (i + 1)
	}

	fmt.Println(score)
}
