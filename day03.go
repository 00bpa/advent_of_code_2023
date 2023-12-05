package main

import (
	"fmt"
	"internal/set"
	"regexp"
	"strconv"
)

type Point struct {
	x int
	y int
}

type Token struct {
	position Point
	length   int
	content  string
}

type Parser struct {
	input   [][]byte
	numbers []*Token
	symbols []*Token
}

func newParser(input [][]byte) *Parser {
	return &Parser{
		input,
		[]*Token{},
		[]*Token{},
	}
}

func (p *Parser) parse() {
	re := regexp.MustCompile(`\d+`)

	// Find numbers
	for y, line := range p.input {
		positions := re.FindAllIndex(line, -1)
		if positions == nil {
			continue
		}
		for _, pos := range positions {
			from := pos[0]
			to := pos[1]

			token := &Token{
				position: Point{from, y},
				length:   to - from,
				content:  string(line[from:to]),
			}
			p.numbers = append(p.numbers, token)
		}
	}

	re = regexp.MustCompile(`[^.\d]`)
	// find symbols
	for y, line := range p.input {
		positions := re.FindAllIndex(line, -1)
		if positions == nil {
			continue
		}
		for _, pos := range positions {
			from := pos[0]
			to := pos[1]

			token := &Token{
				position: Point{from, y},
				length:   to - from,
				content:  string(line[from:to]),
			}
			p.symbols = append(p.symbols, token)
		}
	}
}

func (p *Parser) findNumberAt(point Point) *Token {
	for _, n := range p.numbers {
		if point.y == n.position.y {
			if point.x >= n.position.x && point.x < (n.position.x+n.length) {
				return n
			}
		}
	}
	return nil
}

func (p *Parser) findSymbolAt(point Point) *Token {
	for _, n := range p.symbols {
		if point.y == n.position.y {
			if point.x >= n.position.x && point.x < (n.position.x+n.length) {
				return n
			}
		}
	}
	return nil
}

func surroundingPositions(token *Token) []Point {
	positions := []Point{}

	x := token.position.x
	y := token.position.y
	l := token.length

	// Upper line
	positions = append(positions, Point{x - 1, y - 1})
	positions = append(positions, Point{x + l, y - 1})
	for i := 0; i < l; i++ {
		positions = append(positions, Point{x + i, y - 1})
	}

	// Mid line
	positions = append(positions, Point{x - 1, y})
	positions = append(positions, Point{x + l, y})

	// Lower line
	positions = append(positions, Point{x - 1, y + 1})
	positions = append(positions, Point{x + l, y + 1})
	for i := 0; i < l; i++ {
		positions = append(positions, Point{x + i, y + 1})
	}

	return positions
}

func parseDay03_part1(input [][]byte) int {
	parser := newParser(input)
	parser.parse()

	sum := 0

	for _, token := range parser.numbers {
		for _, pos := range surroundingPositions(token) {
			if parser.findSymbolAt(pos) != nil {
				num, err := strconv.Atoi(token.content)
				check(err)
				sum += num
			}
		}
	}

	return sum
}

func parseDay03_part2(input [][]byte) int {
	parser := newParser(input)
	parser.parse()

	sum := 0

	for _, token := range parser.symbols {
		if token.content != "*" {
			continue
		}

		numbers := set.Set[*Token]{}

		for _, pos := range surroundingPositions(token) {
			t := parser.findNumberAt(pos)
			if t != nil {
				numbers.Insert(t)
			}
		}

		if numbers.Size() > 1 {
			product := 1
			for t := range numbers {
				num, err := strconv.Atoi(t.content)
				check(err)
				product *= num
			}

			sum += product
		}
	}

	return sum
}

func day03() {
	input := readFileAsLines("data/day03_input.txt")
	result := parseDay03_part1(stringSliceToByteArraySlice(input))
	result2 := parseDay03_part2(stringSliceToByteArraySlice(input))

	fmt.Printf("Day03 Result: %d\n", result)
	fmt.Printf("Day03 Result2: %d\n", result2)
}
