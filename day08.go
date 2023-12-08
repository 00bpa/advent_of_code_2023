package main

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

func mapNodes(input []string) map[string][]string {
	nodes := map[string][]string{}
	re := regexp.MustCompile(`[0-9A-Z]{3}`)
	for _, line := range input {
		if line == "" {
			continue
		}
		matches := re.FindAllString(line, -1)
		if len(matches) < 3 {
			panic("Unknown format")
		}

		nodes[matches[0]] = []string{matches[1], matches[2]}
	}
	return nodes
}

func findStartingNodes(nodes map[string][]string) []string {
	startingNodes := []string{}
	for k := range nodes {
		if k[2] == 'A' {
			startingNodes = append(startingNodes, k)
		}
	}
	return startingNodes
}

func day08_part1(input []string) int {
	pattern := input[0]
	patternLen := len(pattern)
	nodes := mapNodes(input[2:])
	current := "AAA"
	steps := 0

	for i := 0; ; i++ {
		if current == "ZZZ" {
			break
		}

		direction := pattern[i%patternLen]
		if direction == 'L' {
			current = nodes[current][0]
		} else {
			current = nodes[current][1]
		}

		steps += 1
	}

	return steps
}

func day08_part2(input []string) *big.Int {
	pattern := input[0]
	patternLen := len(pattern)
	allNodes := mapNodes(input[2:])
	startingNodes := findStartingNodes(allNodes)

	// Find all distances
	distances := []int{}
	for _, startingNode := range startingNodes {
		currentNode := startingNode
		step := 0
		distance := 0
		savedNodes := map[string]struct{}{}
		for {
			currentPattern := pattern[step%patternLen]

			if currentNode[2] == 'Z' {
				nodeSave := strings.Join([]string{
					string(currentPattern),
					currentNode,
				}, "")
				// if node is circular, break
				// NOTE: might be better to do that after saving the distance
				_, ok := savedNodes[nodeSave]
				if ok {
					break
				}

				// save node and distance
				savedNodes[nodeSave] = struct{}{}
				distances = append(distances, distance)
				distance = 0
			}
			// advance node
			nodeInfo := allNodes[currentNode]
			if currentPattern == 'L' {
				currentNode = nodeInfo[0]
			} else {
				currentNode = nodeInfo[1]
			}
			step += 1
			distance += 1
		}
	}

	fmt.Printf("Distances: %v\n", distances)

	firstValue := int64(distances[0])
	retval := big.NewInt(firstValue)
	for i := 1; i < len(distances); i++ {
		otherNum := big.NewInt(int64(distances[i]))
		gcd := big.NewInt(0)
		gcd = gcd.GCD(nil, nil, retval, otherNum)
		retval = retval.Mul(retval, otherNum)
		retval = retval.Div(retval, gcd)
	}

	return retval
}

func day08() {
	input := readFileAsLines("data/day08_input.txt")

	result1 := day08_part1(input)
	result2, err := day08_part2(input).MarshalText()
	check(err)

	fmt.Printf("Day08 result 1: %d\n", result1)
	fmt.Printf("Day08 result 2: %s\n", string(result2))
}
