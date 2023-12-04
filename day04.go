package main

import (
	"fmt"
	"internal/set"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func turnStringNumbersToNumbers(l []string) *set.Set {
	r := &set.Set{}
	for _, x := range l {
		num, err := strconv.Atoi(x)
		check(err)
		r.Insert(num)
	}
	return r
}

func numberOfScratchcardMatches(num1 *set.Set, num2 *set.Set) int {
	return len(*(num1.Intersection(num2)))
}

func parseScratchCard(card string) (*set.Set, *set.Set) {
	winningNumbers, numbers := &set.Set{}, &set.Set{}
	// Parse out the Card Nr part
	parts := strings.Split(card, ":")
	if len(parts) < 2 {
		// Broken card
		return winningNumbers, numbers
	}

	// Get the parts with the numbers
	parts = strings.Split(parts[1], "|")
	if len(parts) < 2 {
		// Broken card
		return winningNumbers, numbers
	}

	// get the numbers
	re := regexp.MustCompile(`\d+`)
	nums1 := re.FindAllString(parts[0], -1)
	nums2 := re.FindAllString(parts[1], -1)
	winningNumbers = turnStringNumbersToNumbers(nums1)
	numbers = turnStringNumbersToNumbers(nums2)
	return winningNumbers, numbers
}

func day04_part1(filename string) int {
	sum := 0
	for _, line := range readFileAsLines(filename) {
		winningNumbers, ourNumbers := parseScratchCard(line)
		matches := numberOfScratchcardMatches(winningNumbers, ourNumbers)
		if matches > 0 {
			sum += int(math.Pow(float64(2), float64(matches-1)))
		}
	}

	return sum
}

func day04_part2(filename string) int {
	input := readFileAsLines(filename)
	multiplierList := make([]int, len(input))

	// Initialize
	for x := range multiplierList {
		multiplierList[x] = 1
	}

	for i, line := range input {
		winningNumbers, ourNumbers := parseScratchCard(line)
		matches := numberOfScratchcardMatches(winningNumbers, ourNumbers)

		currentMultiplier := multiplierList[i]
		for x := i + 1; x <= i+matches; x++ {
			multiplierList[x] += currentMultiplier
		}
	}

	sum := 0
	for _, x := range multiplierList {
		sum += x
	}
	return sum
}

func day04() {
	result1 := day04_part1("data/day04_part1.txt")
	result2 := day04_part2("data/day04_part1.txt")

	fmt.Printf("day04 result1: %d\n", result1)
	fmt.Printf("day04 result2: %d\n", result2)
}
