package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay08_part1_testdata(t *testing.T) {
	input := []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)		",
	}
	result := day08_part1(input)
	assert.Equal(t, 6, result)
}

func TestDay08_part1(t *testing.T) {
	input := readFileAsLines("data/day08_input.txt")
	result1 := day08_part1(input)
	assert.Equal(t, 23147, result1)
}

func TestDay08_part2_testdata(t *testing.T) {
	input := []string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)	",
	}
	result := day08_part2(input)
	assert.Equal(t, big.NewInt(6), result)
}
