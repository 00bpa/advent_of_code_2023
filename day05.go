package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

const (
	MAPPING_SIZE = 500
)

func parseSeeds(input []string) []int {
	re := regexp.MustCompile(`seeds:`)
	re_seeds := regexp.MustCompile(`\d+`)
	var seeds []int

	for _, line := range input {
		if re.Match([]byte(line)) {
			res := re_seeds.FindAllString(line, -1)
			for _, x := range res {
				num, err := strconv.Atoi(x)
				check(err)
				seeds = append(seeds, num)
			}
			break
		}
	}

	return seeds
}

func parseAlmanacParagraph(input []string, paragraph string) [][]int {
	re := regexp.MustCompile(paragraph + " map:")
	reNums := regexp.MustCompile(`\d+`)
	parsing := false
	numbers := [][]int{}

	for _, line := range input {
		if re.Match([]byte(line)) {
			parsing = true
			continue
		}

		if line == "" {
			parsing = false
			continue
		}

		if parsing {
			numStrings := reNums.FindAllString(line, -1)
			entry := []int{}
			for _, x := range numStrings {
				num, err := strconv.Atoi(x)
				check(err)
				entry = append(entry, num)
			}

			if len(entry) == 3 {
				numbers = append(numbers, entry)
			}
		}
	}

	return numbers
}

func findAlmanacMapping(input [][]int, num int) int {
	for _, i := range input {
		from := i[1]
		to := i[0]
		length := i[2]

		if num >= from && num < from+length {
			return to + (num - from)
		}
	}

	return num
}

func day05_part1(input []string) int {
	seeds := parseSeeds(input)
	seedToSoil := parseAlmanacParagraph(input, "seed-to-soil")
	soilToFertilizer := parseAlmanacParagraph(input, "soil-to-fertilizer")
	fertilizerToWater := parseAlmanacParagraph(input, "fertilizer-to-water")
	waterToLight := parseAlmanacParagraph(input, "water-to-light")
	lightToTemp := parseAlmanacParagraph(input, "light-to-temperature")
	tempToHumidity := parseAlmanacParagraph(input, "temperature-to-humidity")
	humidityToLoc := parseAlmanacParagraph(input, "humidity-to-location")

	results := []int{}

	for _, i := range seeds {
		soil := findAlmanacMapping(seedToSoil, i)
		fertilizer := findAlmanacMapping(soilToFertilizer, soil)
		water := findAlmanacMapping(fertilizerToWater, fertilizer)
		light := findAlmanacMapping(waterToLight, water)
		temp := findAlmanacMapping(lightToTemp, light)
		humidity := findAlmanacMapping(tempToHumidity, temp)
		loc := findAlmanacMapping(humidityToLoc, humidity)
		results = append(results, loc)
	}

	sort.Ints(results)

	if len(results) == 0 {
		return 0
	} else {
		return results[0]
	}
}

func day05_part2(input []string) int {

	return 0
}

func day05() {
	input := readFileAsLines("data/day05_input.txt")
	result1 := day05_part1(input)

	fmt.Printf("Day 05 Result 1: %d\n", result1)
}
