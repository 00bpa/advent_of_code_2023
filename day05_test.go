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
