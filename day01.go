package main

import (
	"fmt"
	"regexp"
)

func convertDigit(input []byte) int {
	switch string(input) {
	case "one", "1":
		return 1
	case "two", "2":
		return 2
	case "three", "3":
		return 3
	case "four", "4":
		return 4
	case "five", "5":
		return 5
	case "six", "6":
		return 6
	case "seven", "7":
		return 7
	case "eight", "8":
		return 8
	case "nine", "9":
		return 9
	case "zero", "0":
		return 0
	}

	panic(fmt.Sprintf("Unknown digit: %s", string(input)))
}

func parseFirstAndLastDigit(line string) int {
	input := []byte(line)
	re := regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine|0|1|2|3|4|5|6|7|8|9)`)
	var results []int
	idx := re.FindIndex(input)
	if idx == nil {
		panic("No digits found!")
	}
	for {

		results = append(results, convertDigit(input[idx[0]:idx[1]]))
		input[idx[0]] = '.'
		idx = re.FindIndex(input)
		if idx == nil {
			break
		}
	}

	if len(results) == 0 {
		panic(fmt.Sprintf("No digits found on line: %s", line))
	}

	firstDigit := results[0]
	secondDigit := results[len(results)-1]

	return firstDigit*10 + secondDigit
}

func day01(lines []string) int {
	result := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		num := parseFirstAndLastDigit(line)
		result += num
	}
	return result
}
