package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay11_map_inversion(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
		"..........",
	}
	galaxy := stringsToGalaxy(input)
	result := invertGalaxy(invertGalaxy(galaxy))

	for y := 0; y < len(result); y++ {
		for x := 0; x < len(result[y]); x++ {
			assert.Equal(t, galaxy[y][x], result[y][x])
		}
	}
}

func TestDay11_expandUniverse(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	output := []string{
		"....#........",
		".........#...",
		"#............",
		".............",
		".............",
		"........#....",
		".#...........",
		"............#",
		".............",
		".............",
		".........#...",
		"#....#.......",
	}

	galaxy := stringsToGalaxy(input)
	galaxy = expandUniverse(galaxy, 2)
	result := galaxyToStrings(galaxy)

	for x := 0; x < len(output); x++ {
		assert.Equal(t, output[x], result[x])
	}
}

func TestDay11_part1_testdata(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	galaxy := stringsToGalaxy(input)
	assert.Equal(t, 374, day11_part1(galaxy))
}

func TestDay11_part1(t *testing.T) {
	input := readFileAsLines("data/day11_input.txt")
	assert.Equal(t, 9370588, day11_part1(stringsToGalaxy(input)))
}

func TestDay11_part2_testdata1(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	assert.Equal(t, 1030, day11_exec(stringsToGalaxy(input), 10))
}

func TestDay11_part2_testdata2(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	assert.Equal(t, 8410, day11_exec(stringsToGalaxy(input), 100))
}

func TestDay11_part2(t *testing.T) {
	input := readFileAsLines("data/day11_input.txt")
	assert.Equal(t, 746207878188, day11_part2(stringsToGalaxy(input)))
}

func TestDay11_distances(t *testing.T) {
	input := []string{
		"#..",
		"...",
		"..#",
	}
	expanded := expandUniverse(stringsToGalaxy(input), 10000)
	distances := calculateDistances(expanded, findGalaxies(expanded))
	assert.Equal(t, 1, len(distances))
	assert.Equal(t, 20002, distances[0])
}

func TestDay11_distances2(t *testing.T) {
	input := []string{
		"#.#",
	}
	expanded := expandUniverse(stringsToGalaxy(input), 10000)
	distances := calculateDistances(expanded, findGalaxies(expanded))
	assert.Equal(t, 1, len(distances))
	assert.Equal(t, 10001, distances[0])
}
