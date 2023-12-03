package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay02_part1(t *testing.T) {
	input := "data/day02_input.txt"
	assert.Equal(t, 2406, day02_part1(input))
}

func TestDay02_part2(t *testing.T) {
	input := "data/day02_input.txt"
	assert.Equal(t, 78375, day02_part2(input))
}
