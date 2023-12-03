package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay02_part1(t *testing.T) {
	input := "data/day02_input.txt"
	assert.Equal(t, day02_part1(input), 2406)
}

func TestDay02_part2(t *testing.T) {
	input := "data/day02_input.txt"
	assert.Equal(t, day02_part2(input), 78375)
}
