package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12_part1_testdata(t *testing.T) {
	input := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}

	assert.Equal(t, 21, day12_part1(input))
}

func TestDay12_part1_testdata_line1(t *testing.T) {
	input := []string{
		"???.### 1,1,3",
	}

	assert.Equal(t, 1, day12_part1(input))
}

func TestDay12_part1_testdata_line2(t *testing.T) {
	input := []string{
		".??..??...?##. 1,1,3",
	}

	assert.Equal(t, 4, day12_part1(input))
}

func TestDay12_part1_testdata_line3(t *testing.T) {
	input := []string{
		"?#?#?#?#?#?#?#? 1,3,1,6",
	}

	assert.Equal(t, 1, day12_part1(input))
}

func TestDay12_part1_testdata_line4(t *testing.T) {
	input := []string{
		"????.#...#... 4,1,1",
	}

	assert.Equal(t, 1, day12_part1(input))
}

func TestDay12_part1_testdata_line5(t *testing.T) {
	input := []string{
		"????.######..#####. 1,6,5",
	}

	assert.Equal(t, 4, day12_part1(input))
}

func TestDay12_part1_testdata_line6(t *testing.T) {
	input := []string{
		"?###???????? 3,2,1",
	}

	assert.Equal(t, 10, day12_part1(input))
}
