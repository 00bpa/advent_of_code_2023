package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInt(input string) int {
	num, err := strconv.Atoi(input)
	check(err)
	return num
}

func trimWhitespace(input string) string {
	return strings.Trim(input, " \t\n\r")
}

func readFile(filename string) []byte {
	// Read the file
	dat, err := os.ReadFile(filename)
	check(err)
	return dat
}

func readFileAsLines(filename string) []string {
	dat := readFile(filename)

	// Split the lines
	split := strings.Split(strings.ReplaceAll(string(dat), "\r\n", "\n"), "\n")

	// Last line empty?
	if trimWhitespace(split[len(split)-1]) == "" {
		return split[:len(split)-1]
	} else {
		return split
	}
}

func stringSliceToByteArraySlice(input []string) [][]byte {
	data := [][]byte{}

	for _, inp := range input {
		data = append(data, []byte(inp))
	}

	return data
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		switch arg {
		case "day01":
			day01()
		case "day02":
			day02()
		case "day03":
			day03()
		case "day04":
			day04()
		case "day05":
			day05()
		case "day06":
			day06()
		case "day07":
			day07()
		case "day08":
			day08()
		case "day09":
			day09()
		case "day10":
			day10()
		case "day11":
			day11()
		case "day12":
			day12()
		default:
			fmt.Printf("Unknown parameter: %s\n", arg)
		}
	}
}
