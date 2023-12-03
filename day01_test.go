package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01_part2(t *testing.T) {
	result := day01_exec("data/day01_input.txt")
	assert.Equal(t, result, 54770)
}
