package main

import "fmt"

type galaxyTile struct {
	character rune
	xsize     int
	ysize     int
}

func lineHasNoGalaxies(input []galaxyTile) bool {
	for _, t := range input {
		if t.character != '.' {
			return false
		}
	}
	return true
}

func invertGalaxy(input [][]galaxyTile) [][]galaxyTile {
	result := [][]galaxyTile{}

	for x := 0; x < len(input); x++ {
		temp := []galaxyTile{}
		for y := 0; y < len(input[0]); y++ {
			temp = append(temp, input[x][y])
		}
		result = append(result, temp)
	}

	return result
}

func stringsToGalaxy(input []string) [][]galaxyTile {
	result := [][]galaxyTile{}
	for _, line := range input {
		temp := []galaxyTile{}
		for _, c := range line {
			temp = append(temp, galaxyTile{c, 1, 1})
		}
		result = append(result, temp)
	}
	return result
}

func galaxyToStrings(input [][]galaxyTile) []string {
	// calculate galaxy target width
	width := 0
	for _, c := range input[0] {
		width += c.xsize
	}

	// calculate target height
	height := 0
	for _, l := range input {
		height += l[0].ysize
	}

	// result type in proper size
	result := []string{}

	for _, row := range input {
		temp := ""
		for _, col := range row {
			for i := 0; i < col.xsize; i++ {
				temp += string(col.character)
			}
		}
		for i := 0; i < row[0].ysize; i++ {
			result = append(result, temp)
		}
	}

	return result
}

func expandUniverse(input [][]galaxyTile, mult int) [][]galaxyTile {
	// measure on the y-axis
	for y := 0; y < len(input); y++ {
		if lineHasNoGalaxies(input[y]) {
			for x := 0; x < len(input[y]); x++ {
				input[y][x].ysize *= mult
			}
		}
	}

	// measure the x-axis
	for x := 0; x < len(input[0]); x++ {
		temp := []galaxyTile{}
		for y := 0; y < len(input); y++ {
			temp = append(temp, input[y][x])
		}
		if lineHasNoGalaxies(temp) {
			for y := 0; y < len(input); y++ {
				input[y][x].xsize *= mult
			}
		}
	}

	return input
}

func findGalaxies(input [][]galaxyTile) []mapCoord {
	result := []mapCoord{}
	for y, line := range input {
		for x, node := range line {
			if node.character == '#' {
				result = append(result, mapCoord{x, y})
			}
		}
	}
	return result
}

func calculateDistances(input [][]galaxyTile, galaxies []mapCoord) []int {
	distances := []int{}
	for i := 0; i < len(galaxies); i++ {
		src := galaxies[i]
		for _, dst := range galaxies[i+1:] {
			// x-distance
			xdist := 0
			step := 1
			if src.x > dst.x {
				step = -step
			}

			for x := src.x; x != dst.x; x += step {
				xdist += input[src.y][x].xsize
			}

			// y-distance
			ydist := 0
			step = 1
			if src.y > dst.y {
				step = -step
			}

			for y := src.y; y != dst.y; y += step {
				ydist += input[y][dst.x].ysize
			}

			distances = append(distances, xdist+ydist)
		}
	}

	return distances

	/*
		result := []int{}

		nrGalaxies := len(galaxies)
		for a := 0; a < nrGalaxies; a++ {
			for b := a + 1; b < nrGalaxies; b++ {
				distance := 0
				x_start := min(galaxies[a].x, galaxies[b].x)
				x_end := max(galaxies[a].x, galaxies[b].x)
				y_start := min(galaxies[a].y, galaxies[b].y)
				y_end := max(galaxies[a].y, galaxies[b].y)
				for x := x_start; x < x_end; x++ {
					distance += input[galaxies[a].y][x].xsize
				}
				for y := y_start; y < y_end; y++ {
					distance += input[y][galaxies[b].x].ysize
				}
				result = append(result, distance)
			}
		}

		return result
	*/
}

func day11_exec(input [][]galaxyTile, mult int) int {
	expanded := expandUniverse(input, mult)
	galaxies := findGalaxies(expanded)
	distances := calculateDistances(expanded, galaxies)
	sum := 0
	for _, x := range distances {
		sum += x
	}
	return sum
}

func day11_part1(input [][]galaxyTile) int {
	return day11_exec(input, 2)
}

func day11_part2(input [][]galaxyTile) int {
	return day11_exec(input, 1000000)
}

func day11() {
	input := readFileAsLines("data/day11_input.txt")
	galaxy := stringsToGalaxy(input)
	result1 := day11_part1(galaxy)
	result2 := day11_part2(galaxy)
	fmt.Printf("day11 result 1: %d\n", result1)
	fmt.Printf("day11 result 2: %d\n", result2)
}
