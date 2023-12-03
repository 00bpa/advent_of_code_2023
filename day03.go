package main

import (
	"errors"
	"fmt"
	"strconv"
)

func cellIsSymbol(c byte) bool {
	return !cellIsDigit(c) && c != '.'
}

func cellIsDigit(c byte) bool {
	return c <= '9' && c >= '0'
}

type Coords struct {
	x int
	y int
}

type Size struct {
	w int
	h int
}

func (s Size) Valid(coords Coords) bool {
	return !(coords.x < 0) &&
		!(coords.y < 0) &&
		!(coords.x >= s.w) &&
		!(coords.y >= s.h)
}

func properCoordinates(coords Coords, size Size) []Coords {
	coordList := []Coords{
		{coords.x - 1, coords.y - 1},
		{coords.x - 1, coords.y},
		{coords.x - 1, coords.y + 1},
		{coords.x, coords.y - 1},
		{coords.x, coords.y + 1},
		{coords.x + 1, coords.y - 1},
		{coords.x + 1, coords.y},
		{coords.x + 1, coords.y + 1},
	}

	results := []Coords{}
	for _, c := range coordList {
		if size.Valid(c) {
			results = append(results, c)
		}
	}

	return results
}

func isNextToSymbol(cells [][]byte, x int, y int) bool {
	h := len(cells)
	if h == 0 {
		return false
	}
	w := len(cells[0])
	if w == 0 {
		return false
	}

	coords := properCoordinates(Coords{x, y}, Size{w, h})

	for _, c := range coords {
		if cellIsSymbol(cells[c.y][c.x]) {
			return true
		}
	}

	return false
}

func getNextCell(input []byte, pos int) (byte, error) {
	w := len(input)
	if pos+1 < w {
		return input[pos+1], nil
	}
	return 0, errors.New("out of bounds")
}

func parseDay03_part1(input [][]byte) int {
	var results []int

	for y, line := range input {
		if len(line) == 0 {
			continue
		}

		numBuf := []byte{}
		nextToSymbol := false

		for x, cell := range line {
			if cellIsDigit(cell) {
				numBuf = append(numBuf, cell)
				if !nextToSymbol {
					nextToSymbol = isNextToSymbol(input, x, y)
				}
			}

			c, err := getNextCell(line, x)
			if (!cellIsDigit(c) || err != nil) &&
				len(numBuf) > 0 {
				// Next element is not a digit, we can parse the number here.
				// Otherwise we get a problem at the end of the line
				numStr := string(numBuf)
				numBuf = []byte{}
				if nextToSymbol {
					nextToSymbol = false
					num, err := strconv.Atoi(numStr)
					check(err)
					results = append(results, num)
				}
			}
		}
	}

	// Sum up numbers
	sum := 0
	for _, i := range results {
		sum += i
	}

	return sum
}

func day03() {
	input := readFileAsLines("data/day03_input.txt")

	result := parseDay03_part1(stringSliceToByteArraySlice(input))

	fmt.Printf("Day03 Result: %d\n", result)
}
