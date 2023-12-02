package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var numberRegexp *regexp.Regexp = regexp.MustCompile(`\d+`)

func parseGameData(line string) *map[string]int {
	parts := strings.Split(line, ";")
	retval := map[string]int{
		"blue":  0,
		"red":   0,
		"green": 0,
	}

	for _, p := range parts {
		values := strings.Split(p, ",")
		for _, v := range values {
			if idx := strings.Index(v, "blue"); idx != -1 {
				num, err := strconv.Atoi(trimWhitespace(v[:idx]))
				check(err)
				if num > retval["blue"] {
					retval["blue"] = num
				}
			} else if idx := strings.Index(v, "red"); idx != -1 {
				num, err := strconv.Atoi(trimWhitespace(v[:idx]))
				check(err)
				if num > retval["red"] {
					retval["red"] = num
				}
			} else if idx := strings.Index(v, "green"); idx != -1 {
				num, err := strconv.Atoi(trimWhitespace(v[:idx]))
				check(err)
				if num > retval["green"] {
					retval["green"] = num
				}
			}
		}
	}

	return &retval
}

func divideGameAndData(line string) (int, string) {
	parts := strings.Split(line, ":")
	gameNum, err := strconv.Atoi(numberRegexp.FindString(parts[0]))
	check(err)
	return gameNum, strings.Trim(parts[1], " \t\n\r")
}

func gameDataFits(gameData *map[string]int, red int, green int, blue int) bool {
	if (*gameData)["blue"] <= blue &&
		(*gameData)["red"] <= red &&
		(*gameData)["green"] <= green {
		return true
	}
	return false
}

func day02() {
	input := readFileAsLines("data/day02_input.txt")

	sum := 0
	sum2 := 0
	for _, i := range input {
		if len(trimWhitespace(i)) == 0 {
			continue
		}
		gameNum, gameLine := divideGameAndData(i)
		gameData := parseGameData(gameLine)

		// Part 1
		if gameDataFits(gameData, 12, 13, 14) {
			sum += gameNum
		}

		// Part 2
		sum2 += (*gameData)["blue"] * (*gameData)["red"] * (*gameData)["green"]
	}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Sum2: %d\n", sum2)
}
