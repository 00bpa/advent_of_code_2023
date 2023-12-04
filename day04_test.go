package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay04_part1_testdata(t *testing.T) {
	result := day04_part1("data/day04_part1_testdata.txt")
	assert.Equal(t, 13, result)
}

func TestDay04_part1(t *testing.T) {
	result := day04_part1("data/day04_part1.txt")
	assert.Equal(t, 21138, result)
}

func TestDay04_part2_testdata(t *testing.T) {
	result := day04_part2("data/day04_part1_testdata.txt")
	assert.Equal(t, 30, result)
}

func TestDay04_part2(t *testing.T) {
	result := day04_part2("data/day04_part1.txt")
	assert.Equal(t, 0, result)
}
