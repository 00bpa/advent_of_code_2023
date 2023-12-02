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

func readFileAsLines(filename string) []string {
	// Read the file
	dat, err := os.ReadFile(filename)
	check(err)
	// Split the lines
	return strings.Split(strings.ReplaceAll(string(dat), "\r\n", "\n"), "\n")
}

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		switch arg {
		case "day01":
			day01()
		case "day02":
			day02()
		default:
			fmt.Printf("Unknown parameter: %s\n", arg)
		}
	}
}
