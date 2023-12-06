package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day06_part1(input []string) int {
	reTime := regexp.MustCompile("Time:")
	reDist := regexp.MustCompile("Distance:")
	reNum := regexp.MustCompile(`\d+`)
	times := []int{}
	distances := []int{}

	for _, line := range input {
		if reTime.Match([]byte(line)) {
			values := reNum.FindAllString(line, -1)
			for _, v := range values {
				num, err := strconv.Atoi(v)
				check(err)
				times = append(times, num)
			}
			// Next line
			continue
		}

		if reDist.Match([]byte(line)) {
			values := reNum.FindAllString(line, -1)
			for _, v := range values {
				num, err := strconv.Atoi(v)
				check(err)
				distances = append(distances, num)
			}
			// Next line
			continue
		}
	}

	validTimes := []int{}
	for i := range times {
		time := times[i]
		distance := distances[i]

		sum := 0
		for j := 0; j <= time; j++ {
			// Calculate the possible distance
			buttonTime := j
			travelTime := time - j

			possibleDistance := buttonTime * travelTime

			if possibleDistance > distance {
				sum += 1
			}
		}

		if sum > 0 {
			validTimes = append(validTimes, sum)
		}
	}

	if len(validTimes) > 0 {
		result := 1
		for _, i := range validTimes {
			result *= i
		}
		return result
	} else {
		return 0
	}
}

func day06_part2(input []string) int {
	reTime := regexp.MustCompile("Time:")
	reDist := regexp.MustCompile("Distance:")
	reNum := regexp.MustCompile(`\d+`)
	times := []string{}
	distances := []string{}

	for _, line := range input {
		if reTime.Match([]byte(line)) {
			times = reNum.FindAllString(line, -1)
			// Next line
			continue
		}

		if reDist.Match([]byte(line)) {
			distances = reNum.FindAllString(line, -1)
			// Next line
			continue
		}
	}

	// concatenate the values
	time, err := strconv.Atoi(strings.Join(times, ""))
	check(err)
	distance, err := strconv.Atoi(strings.Join(distances, ""))
	check(err)

	sum := 0
	for j := 0; j <= time; j++ {
		// Calculate the possible distance
		buttonTime := j
		travelTime := time - j

		possibleDistance := buttonTime * travelTime

		if possibleDistance > distance {
			sum += 1
		}
	}

	return sum
}

func day06() {
	input := readFileAsLines("data/day06_input.txt")
	result1 := day06_part1(input)
	result2 := day06_part2(input)
	fmt.Printf("Day06 Result 1: %d\n", result1)
	fmt.Printf("Day06 Result 2: %d\n", result2)
}
