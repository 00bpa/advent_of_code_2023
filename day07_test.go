package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay07_part2_testdata(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	result := day07_part2(input)
	assert.Equal(t, 5905, result)
}

func TestDay07_part2(t *testing.T) {
	input := readFileAsLines("data/day07_input.txt")
	result := day07_part2(input)
	assert.Equal(t, 247885995, result)
}
