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
	return gameNum, trimWhitespace(parts[1])
}

func gameDataFits(gameData *map[string]int, red int, green int, blue int) bool {
	if (*gameData)["blue"] <= blue &&
		(*gameData)["red"] <= red &&
		(*gameData)["green"] <= green {
		return true
	}
	return false
}

func day02_exec(inputfile string, f func(*int, int, *map[string]int)) int {
	input := readFileAsLines(inputfile)

	sum := 0

	for _, i := range input {
		if len(trimWhitespace(i)) == 0 {
			continue
		}
		gameNum, gameLine := divideGameAndData(i)
		gameData := parseGameData(gameLine)

		f(&sum, gameNum, gameData)
	}

	return sum
}

func day02_part1(inputfile string) int {
	return day02_exec(inputfile,
		func(sum *int, gameNum int, gameData *map[string]int) {
			// Part 1
			if gameDataFits(gameData, 12, 13, 14) {
				*sum += gameNum
			}
		})
}

func day02_part2(inputfile string) int {
	return day02_exec(inputfile,
		func(sum *int, gameNum int, gameData *map[string]int) {
			// Part 2
			*sum += (*gameData)["blue"] * (*gameData)["red"] * (*gameData)["green"]
		})
}

func day02() {
	input := "data/day02_input.txt"

	sum := day02_part1(input)
	sum2 := day02_part2(input)

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Sum2: %d\n", sum2)
}
