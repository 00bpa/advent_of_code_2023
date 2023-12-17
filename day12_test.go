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

func TestDay12_part2_testdata_line1(t *testing.T) {
	input := []string{
		"???.### 1,1,3",
	}

	assert.Equal(t, 1, day12_part2(input))
}

func TestDay12_part1_testdata_line2(t *testing.T) {
	input := []string{
		".??..??...?##. 1,1,3",
	}

	assert.Equal(t, 4, day12_part1(input))
}

func TestDay12_part2_testdata_line2(t *testing.T) {
	input := []string{
		".??..??...?##. 1,1,3",
	}

	assert.Equal(t, 16384, day12_part2(input))
}

func TestDay12_part1_testdata_line3(t *testing.T) {
	input := []string{
		"?#?#?#?#?#?#?#? 1,3,1,6",
	}

	assert.Equal(t, 1, day12_part1(input))
}

func TestDay12_part2_testdata_line3(t *testing.T) {
	input := []string{
		"?#?#?#?#?#?#?#? 1,3,1,6",
	}

	assert.Equal(t, 1, day12_part2(input))
}

func TestDay12_part1_testdata_line4(t *testing.T) {
	input := []string{
		"????.#...#... 4,1,1",
	}

	assert.Equal(t, 1, day12_part1(input))
}

func TestDay12_part2_testdata_line4(t *testing.T) {
	input := []string{
		"????.#...#... 4,1,1",
	}

	assert.Equal(t, 16, day12_part2(input))
}

func TestDay12_part1_testdata_line5(t *testing.T) {
	input := []string{
		"????.######..#####. 1,6,5",
	}

	assert.Equal(t, 4, day12_part1(input))
}

func TestDay12_part2_testdata_line5(t *testing.T) {
	input := []string{
		"????.######..#####. 1,6,5",
	}

	assert.Equal(t, 2500, day12_part2(input))
}

func TestDay12_part1_testdata_line6(t *testing.T) {
	input := []string{
		"?###???????? 3,2,1",
	}

	assert.Equal(t, 10, day12_part1(input))
}

func TestDay12_part2_testdata_line6(t *testing.T) {
	input := []string{
		"?###???????? 3,2,1",
	}

	assert.Equal(t, 506250, day12_part2(input))
}

func TestDay12_part2_testdata(t *testing.T) {
	input := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}

	assert.Equal(t, 525152, day12_part2(input))
}

func TestDay12_part1(t *testing.T) {
	input := readFileAsLines("data/day12_input.txt")
	assert.Equal(t, 7007, day12_part1(input))
}

func TestDay12_storeCache(t *testing.T) {
	input := "abcdef"
	nums := []int{1, 3, 5, 3}
	result := 1337
	storeCacheValue(input, nums, result)
	otherResult, ok := fetchCacheValue(input, nums)
	assert.Equal(t, true, ok)
	assert.Equal(t, result, otherResult)
}
