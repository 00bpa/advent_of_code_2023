package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func day12_part1(input []string) int {
	reNum := regexp.MustCompile(`\d+`)
	sum := 0

	for _, line := range input {
		// Parse the numbers
		numTexts := reNum.FindAllString(line, -1)
		numbers := make([]int, len(numTexts))
		for i, t := range numTexts {
			numbers[i] = parseInt(t)
		}

		// split the line
		str := strings.Split(line, " ")[0]
		placeholders := []int{}

		// find placeholders
		for i, c := range str {
			if c == '?' {
				placeholders = append(placeholders, i)
			}
		}

		limit := int(math.Pow(2, float64(len(placeholders))))
		rePattern := regexp.MustCompile(`#+`)
		for n := 0; n < limit; n++ {
			strbuf := []byte(str)
			for i := 0; i < len(placeholders); i++ {
				// calculate the bit
				bit := (n >> i) & 1
				if bit == 1 {
					strbuf[placeholders[i]] = '#'
				} else {
					strbuf[placeholders[i]] = '.'
				}
			}
			patterns := rePattern.FindAllString(string(strbuf), -1)

			if len(patterns) != len(numbers) {
				// no need to check a patternlist that has a different
				// length from the numbers we look for
				continue
			}

			equal := true

			for i := 0; i < len(patterns); i++ {
				if numbers[i] != len(patterns[i]) {
					equal = false
					break
				}
			}

			if equal {
				sum += 1
			}
		}
	}

	return sum
}

func day12() {
	input := readFileAsLines("data/day12_input.txt")
	result1 := day12_part1(input)
	fmt.Printf("Day12 part1 result: %d\n", result1)
}
