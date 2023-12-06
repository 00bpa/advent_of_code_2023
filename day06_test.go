package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay06_part1_testdata(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	result := day06_part1(input)
	assert.Equal(t, 288, result)
}

func TestDay06_part1(t *testing.T) {
	input := readFileAsLines("data/day06_input.txt")
	result := day06_part1(input)
	assert.Equal(t, 771628, result)
}

func Test06_part2_testdata(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	result := day06_part2(input)
	assert.Equal(t, 71503, result)
}

func TestDay06_part2(t *testing.T) {
	input := readFileAsLines("data/day06_input.txt")
	result := day06_part2(input)
	assert.Equal(t, 27363861, result)
}
