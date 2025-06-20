package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10_part1_testdata1(t *testing.T) {
	input := []string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	}
	assert.Equal(t, 4, day10_part1(input))
}

func TestDay10_part1_testdata2(t *testing.T) {
	input := []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}
	assert.Equal(t, 8, day10_part1(input))
}

func TestDay10_part1(t *testing.T) {
	input := readFileAsLines("data/day10_input.txt")
	assert.Equal(t, 6968, day10_part1(input))
}

func TestDay10_part2_testdata1(t *testing.T) {
	input := []string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	}
	assert.Equal(t, 4, day10_part2(input))
}

func TestDay10_part2_testdata2(t *testing.T) {
	input := []string{
		".F----7F7F7F7F-7....",
		".|F--7||||||||FJ....",
		".||.FJ||||||||L7....",
		"FJL7L7LJLJ||LJ.L-7..",
		"L--J.L7...LJS7F-7L7.",
		"....F-J..F7FJ|L7L7L7",
		"....L7.F7||L7|.L7L7|",
		".....|FJLJ|FJ|F7|.LJ",
		"....FJL-7.||.||||...",
		"....L---J.LJ.LJLJ...",
	}
	assert.Equal(t, 8, day10_part2(input))
}

func TestDay10_part2_testdata3(t *testing.T) {
	input := []string{
		"FF7FSF7F7F7F7F7F---7",
		"L|LJ||||||||||||F--J",
		"FL-7LJLJ||||||LJL-77",
		"F--JF--7||LJLJ7F7FJ-",
		"L---JF-JLJ.||-FJLJJ7",
		"|F|F-JF---7F7-L7L|7|",
		"|FFJF7L7F-JF7|JL---7",
		"7-L-JL7||F7|L7F-7F7|",
		"L.L7LFJ|||||FJL7||LJ",
		"L7JLJL-JLJLJL--JLJ.L",
	}
	assert.Equal(t, 10, day10_part2(input))
}
