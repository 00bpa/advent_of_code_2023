package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay03_part1(t *testing.T) {
	input := readFileAsLines("data/day03_input.txt")

	got := parseDay03_part1(stringSliceToByteArraySlice(input))

	assert.Equal(t, got, 529618)
}
