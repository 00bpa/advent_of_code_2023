package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay09_part1_testdata(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45	",
	}
	assert.Equal(t, 114, day09_part1(input))
}

func TestDay09_part1(t *testing.T) {
	input := readFileAsLines("data/day09_input.txt")
	assert.Equal(t, 2175229206, day09_part1(input))
}

func TestDay09_part2_testdata(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45	",
	}
	assert.Equal(t, 2, day09_part2(input))
}

func TestDay09_part2(t *testing.T) {
	input := readFileAsLines("data/day09_input.txt")
	assert.Equal(t, 942, day09_part2(input))
}
