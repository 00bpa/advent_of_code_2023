package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay05_part1_testdata(t *testing.T) {
	input := readFileAsLines("data/day05_part1_testdata.txt")
	result := day05_part1(input)
	assert.Equal(t, 35, result)
}

func TestDay05_part1(t *testing.T) {
	input := readFileAsLines("data/day05_input.txt")
	result := day05_part1(input)
	assert.Equal(t, 26273516, result)
}

func TestDay05_part2_testdata(t *testing.T) {
	input := readFileAsLines("data/day05_part1_testdata.txt")
	result := day05_part2(input)
	assert.Equal(t, 46, result)
}

// NOTE: Not testing the part 2 with the input data, it takes too long.
// If I should ever try to improve the algo, the result should be
// 34039469
