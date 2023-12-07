package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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
	return strings.Split(strings.ReplaceAll(string(dat), "\r\n", "\n"), "\n")
}

func stringSliceToByteArraySlice(input []string) [][]byte {
	data := [][]byte{}

	for _, inp := range input {
		data = append(data, []byte(inp))
	}

	return data
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
		default:
			fmt.Printf("Unknown parameter: %s\n", arg)
		}
	}
}
