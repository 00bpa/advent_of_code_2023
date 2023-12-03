package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay03_part1_testdata(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..		",
	}
	got := parseDay03_part1(stringSliceToByteArraySlice(input))
	assert.Equal(t, 4361, got)
}

func TestDay03_part1(t *testing.T) {
	input := readFileAsLines("data/day03_input.txt")
	got := parseDay03_part1(stringSliceToByteArraySlice(input))
	assert.Equal(t, 529618, got)
}

func TestDay03_part2_testdata(t *testing.T) {
	input := readFileAsLines("data/day03_part2_testdata.txt")
	got := parseDay03_part2(stringSliceToByteArraySlice(input))
	assert.Equal(t, 467835, got)
}

func TestDay03_part2(t *testing.T) {
	input := readFileAsLines("data/day03_input.txt")
	got := parseDay03_part2(stringSliceToByteArraySlice(input))
	assert.Equal(t, 77509019, got)
}
