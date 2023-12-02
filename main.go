package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFileAsLines(filename string) []string {
	// Read the file
	dat, err := os.ReadFile(filename)
	check(err)
	// Split the lines
	return strings.Split(strings.ReplaceAll(string(dat), "\r\n", "\n"), "\n")
}

func main() {
	lines := readFileAsLines("data/day01_input.txt")
	result := day01(lines)
	fmt.Printf("Result: %v\n", result)
}
