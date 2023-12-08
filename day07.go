package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

type Hand struct {
	hand   string
	cards  map[rune]int
	jokers int
	bid    int
	value  int
}

// Hand sorting
type ByHand []Hand

func (a ByHand) Len() int      { return len(a) }
func (a ByHand) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByHand) Less(i, j int) bool {
	if a[i].value == a[j].value {
		for x := range a[i].hand {
			if a[i].hand[x] == a[j].hand[x] {
				continue
			}
			return cardValue(a[i].hand[x]) < cardValue(a[j].hand[x])
		}
		return false
	} else {
		return a[i].value < a[j].value
	}
}

func cardValue(card byte) int {
	for val, c := range []byte("J23456789TQKA") {
		if c == card {
			return val
		}
	}
	panic("Error: Unknown card")
}

func separateCards(hand string) (map[rune]int, int) {
	retval := map[rune]int{}
	jokers := 0
	for _, r := range hand {
		if r == 'J' {
			jokers += 1
			continue
		}
		retval[r] += 1
	}
	return retval, jokers
}

func weightCards(valueMap map[rune]int, jokers int) int {
	cardCounts := []int{}
	for _, v := range valueMap {
		cardCounts = append(cardCounts, v)
	}

	sort.Ints(cardCounts)

	// Reverse cards
	for i, j := 0, len(cardCounts)-1; i < j; i, j = i+1, j-1 {
		cardCounts[i], cardCounts[j] = cardCounts[j], cardCounts[i]
	}

	score := 0

	if len(cardCounts) == 0 {
		cardCounts = append(cardCounts, jokers)
	} else {
		cardCounts[0] += jokers
	}

	for _, x := range cardCounts {
		if x == 5 {
			score = FIVE_OF_A_KIND
		} else if x == 4 {
			score = FOUR_OF_A_KIND
		} else if x == 3 {
			score = THREE_OF_A_KIND
		} else if x == 2 && score == 0 {
			score = ONE_PAIR
		} else if x == 2 && score == ONE_PAIR {
			score = TWO_PAIR
		} else if x == 2 && score == THREE_OF_A_KIND {
			score = FULL_HOUSE
		}
	}

	return score
}

func parseHand(input string) Hand {
	splits := strings.Split(input, " ")
	hand := splits[0]
	bid, err := strconv.Atoi(splits[1])
	check(err)
	valueMap, jokers := separateCards(hand)
	value := weightCards(valueMap, jokers)
	return Hand{
		hand,
		valueMap,
		jokers,
		bid,
		value,
	}
}

func day07_part2(input []string) int {
	hands := []Hand{}
	for _, line := range input {
		if line == "" {
			continue
		}
		hands = append(hands, parseHand(line))
	}

	sort.Sort(ByHand(hands))

	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}
	return sum
}

func day07() {
	input := readFileAsLines("data/day07_input.txt")
	result2 := day07_part2(input)
	fmt.Printf("Day07 result2: %d\n", result2)
}
